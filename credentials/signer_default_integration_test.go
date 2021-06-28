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

package credentials_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"

	"github.com/ory/oathkeeper/internal"
)

func BenchmarkDefaultSigner(b *testing.B) {
	conf := internal.NewConfigurationWithDefaults()
	reg := internal.NewRegistry(conf)
	ctx := context.Background()

	for alg, keys := range map[string]string{
		"RS256": "file://../test/stub/jwks-rsa-multiple.json",
		"ES256": "file://../test/stub/jwks-ecdsa.json",
		"HS256": "file://../test/stub/jwks-hs.json",
	} {
		b.Run("alg="+alg, func(b *testing.B) {
			jwks, _ := url.Parse(keys)
			for i := 0; i < b.N; i++ {
				if _, err := reg.CredentialsSigner().Sign(ctx, jwks, jwt.MapClaims{
					"custom-claim2": 3.14159,
					"custom-claim3": true,
					"exp":           time.Now().Add(time.Minute).Unix(),
					"iat":           time.Now().Unix(),
					"iss":           "issuer",
					"nbf":           time.Now().Unix(),
					"sub":           "some subject",
				}); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
