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

package main

import "github.com/getkin/kin-openapi/openapi3"

type CfnSchema struct {
	Definitions          interface{} `json:"definitions,omitempty"`
	Properties           interface{} `json:"properties"`
	Description          string      `json:"description"`
	TypeName             string      `json:"typeName"`
	SourceURL            string      `json:"sourceUrl"`
	FileName             string      `json:"-"`
	Handlers             Handlers    `json:"handlers"`
	PrimaryIdentifier    []string    `json:"primaryIdentifier"`
	ReadOnlyProperties   []string    `json:"readOnlyProperties,omitempty"`
	Required             []string    `json:"required,omitempty"`
	AdditionalProperties bool        `json:"additionalProperties"`
}

type RequiredParams struct {
	FileName     string       `json:"-"`
	CreateFields RequireParam `json:"CreateFields"`
	ReadFields   RequireParam `json:"ReadFields"`
	UpdateFields RequireParam `json:"UpdateFields"`
	DeleteFields RequireParam `json:"DeleteFields"`
	ListFields   RequireParam `json:"ListFields"`
}

type RequireParam struct {
	InputParams    []string `json:",omitempty"`
	RequiredParams []string `json:",omitempty"`
}

type Definitions struct {
	Type                 *openapi3.Types     `json:"type,omitempty"`
	Properties           map[string]Property `json:"properties,omitempty"`
	AdditionalProperties bool                `json:"additionalProperties"`
}

type Handlers struct {
	Create struct {
		Permissions []string `json:"permissions"`
	} `json:"create"`
	Read struct {
		Permissions []string `json:"permissions"`
	} `json:"read"`
	Update struct {
		Permissions []string `json:"permissions"`
	} `json:"update"`
	Delete struct {
		Permissions []string `json:"permissions"`
	} `json:"delete"`
}

type Property struct {
	Type                 *openapi3.Types `json:"type,omitempty"`
	MaxLength            *uint64         `json:"maxLength,omitempty"`
	Items                *Items          `json:"items,omitempty"`
	Description          string          `json:"description,omitempty"`
	Ref                  string          `json:"$ref,omitempty"`
	Pattern              string          `json:"pattern,omitempty"`
	Enum                 []interface{}   `json:"enum,omitempty"`
	Required             []string        `json:"-"`
	MinLength            uint64          `json:"minLength,omitempty"`
	InsertionOrder       bool            `json:"insertionOrder,omitempty"`
	AdditionalProperties bool            `json:"additionalProperties,omitempty"`
	ReadOnly             bool            `json:"-"`
}

type Items struct {
	Ref         string          `json:"$ref,omitempty"`
	Type        *openapi3.Types `json:"type,omitempty"`
	Enum        []interface{}   `json:"enum,omitempty"`
	UniqueItems bool            `json:"uniqueItems,omitempty"`
}

var handler = `{
  "create": {
    "permissions": []
  },
  "read": {
    "permissions": []
  },
  "update": {
    "permissions": []
  },
  "delete": {
    "permissions": []
  }
}`

type OpenAPIMapping struct {
	Resources []struct {
		TypeName     string   `json:"typeName"`
		ContentType  string   `json:"contentType"`
		OpenAPIPaths []string `json:"openApiPath"`
	} `json:"resources"`
}
