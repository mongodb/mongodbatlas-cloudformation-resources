// Copyright 2026 MongoDB Inc
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

package util

import "net/http"

// StatusNotFound checks if the response status code is 404 Not Found.
func StatusNotFound(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusNotFound
}

// StatusConflict checks if the response status code is 409 Conflict.
func StatusConflict(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusConflict
}

// StatusBadRequest checks if the response status code is 400 Bad Request.
func StatusBadRequest(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusBadRequest
}

// StatusUnauthorized checks if the response status code is 401 Unauthorized.
func StatusUnauthorized(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusUnauthorized
}

// StatusInternalServerError checks if the response status code is 500 Internal Server Error.
func StatusInternalServerError(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusInternalServerError
}

// StatusServiceUnavailable checks if the response status code is 503 Service Unavailable.
func StatusServiceUnavailable(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusServiceUnavailable
}
