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

package helper_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/oathkeeper/helper"
)

const (
	defaultHeaderName = "Authorization"
)

func TestBearerTokenFromRequest(t *testing.T) {
	t.Run("case=token should be received from default header if custom location is not set and token is present", func(t *testing.T) {
		expectedToken := "token"
		request := &http.Request{Header: http.Header{defaultHeaderName: {"bearer " + expectedToken}}}
		token := helper.BearerTokenFromRequest(request, nil)
		assert.Equal(t, expectedToken, token)
	})
	t.Run("case=should return empty string if custom location is not set and token is not present in default header", func(t *testing.T) {
		request := &http.Request{}
		token := helper.BearerTokenFromRequest(request, nil)
		assert.Empty(t, token)
	})
	t.Run("case=should return empty string if custom location is set to header and token is not present in that header", func(t *testing.T) {
		customHeaderName := "Custom-Auth-Header"
		request := &http.Request{Header: http.Header{defaultHeaderName: {"bearer token"}}}
		tokenLocation := helper.BearerTokenLocation{Header: &customHeaderName}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Empty(t, token)
	})
	t.Run("case=should return empty string if custom location is set to query parameter and token is not present in that query parameter", func(t *testing.T) {
		customQueryParameterName := "Custom-Auth"
		request := &http.Request{Header: http.Header{defaultHeaderName: {"bearer token"}}}
		tokenLocation := helper.BearerTokenLocation{QueryParameter: &customQueryParameterName}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Empty(t, token)
	})
	t.Run("case=token should be received from custom header if custom location is set to header and token is present", func(t *testing.T) {
		expectedToken := "token"
		customHeaderName := "Custom-Auth-Header"
		request := &http.Request{Header: http.Header{customHeaderName: {expectedToken}}}
		tokenLocation := helper.BearerTokenLocation{Header: &customHeaderName}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Equal(t, expectedToken, token)
	})
	t.Run("case=token should be received from custom header if custom location is set to query parameter and token is present", func(t *testing.T) {
		expectedToken := "token"
		customQueryParameterName := "Custom-Auth"
		request := &http.Request{
			Form: map[string][]string{
				customQueryParameterName: {expectedToken},
			},
		}
		tokenLocation := helper.BearerTokenLocation{QueryParameter: &customQueryParameterName}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Equal(t, expectedToken, token)
	})
	t.Run("case=token should be received from default header if custom token location is set, but neither Header nor Query Param is configured", func(t *testing.T) {
		expectedToken := "token"
		request := &http.Request{Header: http.Header{defaultHeaderName: {"bearer " + expectedToken}}}
		tokenLocation := helper.BearerTokenLocation{}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Equal(t, expectedToken, token)
	})
	t.Run("case=token should be received from default header if custom location is set to default location and token is present", func(t *testing.T) {
		expectedToken := "token"
		customHeaderName := "Authorization"
		request := &http.Request{Header: http.Header{customHeaderName: {"bearer " + expectedToken}}}
		tokenLocation := helper.BearerTokenLocation{Header: &customHeaderName}
		token := helper.BearerTokenFromRequest(request, &tokenLocation)
		assert.Equal(t, expectedToken, token)
	})
}
