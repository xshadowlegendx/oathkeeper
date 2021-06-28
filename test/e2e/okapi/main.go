// Copyright 2021 Ory GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"gopkg.in/square/go-jose.v2"

	"github.com/ory/oathkeeper/x"
	"github.com/ory/x/urlx"
)

var jwtm = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		u := x.ParseURLOrPanic(os.Getenv("OATHKEEPER_API"))
		res, err := http.Get(urlx.AppendPaths(u, "/.well-known/jwks.json").String())
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			panic("not 200")
		}

		var jwks jose.JSONWebKeySet
		if err := json.NewDecoder(res.Body).Decode(&jwks); err != nil {
			panic(err)
		}

		return jwks.Key(token.Header["kid"].(string))[0].Key.(*rsa.PublicKey), nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

func main() {
	router := httprouter.New()

	router.GET("/jwt", jwtHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6677"
	}

	n := negroni.Classic()
	n.UseHandler(router)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: n,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func jwtHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := jwtm.CheckJWT(w, r); err != nil {
		return
	}
	_, _ = w.Write([]byte("ok"))
}
