package main

type CfnSchema struct {
	AdditionalProperties bool        `json:"additionalProperties"`
	Definitions          interface{} `json:"definitions,omitempty"`
	Description          string      `json:"description"`
	Handlers             Handlers    `json:"handlers"`
	PrimaryIdentifier    []string    `json:"primaryIdentifier"`
	Properties           interface{} `json:"properties"`
	ReadOnlyProperties   []string    `json:"readOnlyProperties,omitempty"`
	Required             []string    `json:"required,omitempty"`
	TypeName             string      `json:"typeName"`
	SourceUrl            string      `json:"sourceUrl"`
	FileName             string      `json:"-"`
}

type RequiredParams struct {
	CreateFields RequireParam `json:"CreateFields"`
	ReadFields   RequireParam `json:"ReadFields"`
	UpdateFields RequireParam `json:"UpdateFields"`
	DeleteFields RequireParam `json:"DeleteFields"`
	ListFields   RequireParam `json:"ListFields"`
	FileName     string       `json:"-"`
}

type RequireParam struct {
	InputParams    []string `json:",omitempty"`
	RequiredParams []string `json:",omitempty"`
}

type Definitions struct {
	Type                 string              `json:"type,omitempty"`
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
	Type                 string        `json:"type,omitempty"`
	Description          string        `json:"description,omitempty"`
	MaxLength            *uint64       `json:"maxLength,omitempty"`
	MinLength            uint64        `json:"minLength,omitempty"`
	InsertionOrder       bool          `json:"insertionOrder,omitempty"`
	Ref                  string        `json:"$ref,omitempty"`
	AdditionalProperties bool          `json:"additionalProperties,omitempty"`
	Enum                 []interface{} `json:"enum,omitempty"`
	Pattern              string        `json:"pattern,omitempty"`
	Items                *Items        `json:"items,omitempty"`
	ReadOnly             bool          `json:"-"`
	Required             []string      `json:"-"`
}

type Items struct {
	Ref         string        `json:"$ref,omitempty"`
	Type        string        `json:"type,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`
	UniqueItems bool          `json:"uniqueItems,omitempty"`
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

type OpenApiMapping struct {
	Resources []struct {
		TypeName     string   `json:"typeName"`
		OpenApiPaths []string `json:"openApiPath"`
	} `json:"resources"`
}
