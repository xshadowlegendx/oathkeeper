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

// The standard error format
// swagger:response genericError
type genericError struct {
	// in: body
	Body struct {
		Code int `json:"code,omitempty"`

		Status string `json:"status,omitempty"`

		Request string `json:"request,omitempty"`

		Reason string `json:"reason,omitempty"`

		Details []map[string]interface{} `json:"details,omitempty"`

		Message string `json:"message"`
	}
}

// An empty response
// swagger:response emptyResponse
type emptyResponse struct{}
