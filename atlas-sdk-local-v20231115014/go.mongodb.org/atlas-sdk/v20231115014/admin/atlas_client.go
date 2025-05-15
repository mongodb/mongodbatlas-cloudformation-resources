// Copyright 2023 MongoDB Inc
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

package admin // import "go.mongodb.org/atlas-sdk/v20231115014/admin"

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/mongodb-forks/digest"

	"go.mongodb.org/atlas-sdk/v20231115014/internal/core"
)

const (
	// DefaultCloudURL is default base URL for the services.
	DefaultCloudURL = "https://cloud.mongodb.com"
	// Version the version of the current API client inherited from.
	Version = core.Version
	// ClientName of the v2 API client.
	ClientName = "go-atlas-sdk-admin"
)

// NewClient returns a new API Client.
func NewClient(modifiers ...ClientModifier) (*APIClient, error) {
	userAgent := fmt.Sprintf("%s/%s (%s;%s)", ClientName, Version, runtime.GOOS, runtime.GOARCH)
	defaultConfig := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{ServerConfiguration{
			URL: DefaultCloudURL,
		},
		},
		UserAgent: userAgent,
	}
	for _, modifierFunction := range modifiers {
		err := modifierFunction(defaultConfig)
		if err != nil {
			return nil, err
		}
	}

	return NewAPIClient(defaultConfig), nil
}

// ClientModifiers lets you create function that controls configuration before creating client.
type ClientModifier func(*Configuration) error

// UseDigestAuth provides Digest authentication for Go SDK.
// UseDigestAuth is provided as helper to create a default HTTP client that supports HTTP Digest authentication.
// Warning: any previously set httpClient will be overwritten. To fully customize HttpClient use UseHTTPClient method.
func UseDigestAuth(apiKey, apiSecret string) ClientModifier {
	return func(c *Configuration) error {
		transport := digest.NewTransport(apiKey, apiSecret)
		httpClient, err := transport.Client()
		if err != nil {
			return err
		}
		c.HTTPClient = httpClient
		return nil
	}
}

// Advanced modifiers.

// UseHTTPClient set custom http client implementation.
//
// Warning: UseHTTPClient overrides any previously set httpClient including the one set by UseDigestAuth.
// To set a custom http client with HTTP diggest support use:
//
//	transport := digest.NewTransportWithHTTPRoundTripper(apiKey, apiSecret, yourHttpTransport)
//	client := UseHTTPClient(transport.Client())
func UseHTTPClient(client *http.Client) ClientModifier {
	return func(c *Configuration) error {
		c.HTTPClient = client
		return nil
	}
}

// UseDebug enable debug level logging.
func UseDebug(debug bool) ClientModifier {
	return func(c *Configuration) error {
		c.Debug = debug
		return nil
	}
}

// UseBaseURL set custom base url. If empty, default is used.
func UseBaseURL(baseURL string) ClientModifier {
	return func(c *Configuration) error {
		if baseURL == "" {
			baseURL = DefaultCloudURL
		}
		urlWithoutSuffix := strings.TrimSuffix(baseURL, "/")
		c.Servers = ServerConfigurations{ServerConfiguration{
			URL: urlWithoutSuffix,
		}}
		return nil
	}
}

// UseUserAgent set custom UserAgent header.
func UseUserAgent(userAgent string) ClientModifier {
	return func(c *Configuration) error {
		c.UserAgent = userAgent
		return nil
	}
}

// AsError checks if API returned known error type.
func AsError(err error) (*ApiError, bool) {
	var openapiError *GenericOpenAPIError
	if ok := errors.As(err, &openapiError); !ok {
		return nil, false
	}
	errModel := openapiError.Model()
	return &errModel, true
}

// IsErrorCode returns true if the error contains the specific code.
func IsErrorCode(err error, code string) bool {
	mappedErr, _ := AsError(err)
	if mappedErr == nil {
		return false
	}
	return mappedErr.GetErrorCode() == code
}
