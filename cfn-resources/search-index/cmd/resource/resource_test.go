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

package resource_test

import (
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/search-index/cmd/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertToAnySlice(t *testing.T) {
	testCases := []struct {
		name        string
		input       []string
		expected    []any
		expectError bool
	}{
		{
			name:        "Empty list",
			input:       []string{},
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Single valid JSON",
			input:       []string{`{"type": "icuNormalize"}`},
			expected:    []any{map[string]any{"type": "icuNormalize"}},
			expectError: false,
		},
		{
			name: "Multiple valid JSONs",
			input: []string{
				`{"type": "icuNormalize"}`,
				`{"type": "lowercase"}`,
			},
			expected: []any{
				map[string]any{"type": "icuNormalize"},
				map[string]any{"type": "lowercase"},
			},
			expectError: false,
		},
		{
			name:        "Invalid JSON",
			input:       []string{`invalid json`},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.ConvertToAnySlice(tc.input)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestNewTokenizerModel(t *testing.T) {
	testCases := []struct {
		tokenizer *resource.ApiAtlasFTSAnalyzersTokenizer
		expected  map[string]any
		name      string
	}{
		{
			name:      "Nil input",
			tokenizer: nil,
			expected:  nil,
		},
		{
			name: "All parameters",
			tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
				MaxGram:        ptr.Int(15),
				MinGram:        ptr.Int(3),
				Type:           ptr.String("edgeGram"),
				Group:          ptr.Int(1),
				Pattern:        ptr.String("\\W+"),
				MaxTokenLength: ptr.Int(255),
			},
			expected: map[string]any{
				"maxGram":        15,
				"minGram":        3,
				"type":           "edgeGram",
				"group":          1,
				"pattern":        "\\W+",
				"maxTokenLength": 255,
			},
		},
		{
			name: "Partial parameters",
			tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
				Type:           ptr.String("standard"),
				MaxTokenLength: ptr.Int(100),
			},
			expected: map[string]any{
				"type":           "standard",
				"maxTokenLength": 100,
			},
		},
		{
			name: "Only type",
			tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
				Type: ptr.String("keyword"),
			},
			expected: map[string]any{
				"type": "keyword",
			},
		},
		{
			name: "With pattern and group",
			tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
				Type:    ptr.String("regexSplit"),
				Pattern: ptr.String("[\\s,]+"),
				Group:   ptr.Int(0),
			},
			expected: map[string]any{
				"type":    "regexSplit",
				"pattern": "[\\s,]+",
				"group":   0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resource.NewTokenizerModel(tc.tokenizer)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestConvertStringToStoredSource(t *testing.T) {
	testCases := []struct {
		expected    any
		input       *string
		name        string
		expectError bool
	}{
		{
			name:        "Nil input",
			input:       nil,
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Empty string",
			input:       ptr.String(""),
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Boolean true",
			input:       ptr.String("true"),
			expected:    true,
			expectError: false,
		},
		{
			name:        "Boolean false",
			input:       ptr.String("false"),
			expected:    false,
			expectError: false,
		},
		{
			name:        "JSON object with include",
			input:       ptr.String(`{"include":["name","price"]}`),
			expected:    map[string]any{"include": []any{"name", "price"}},
			expectError: false,
		},
		{
			name:        "JSON object with exclude",
			input:       ptr.String(`{"exclude":["_id"]}`),
			expected:    map[string]any{"exclude": []any{"_id"}},
			expectError: false,
		},
		{
			name:        "Complex JSON object",
			input:       ptr.String(`{"include":["name","address"],"exclude":["internal"]}`),
			expected:    map[string]any{"include": []any{"name", "address"}, "exclude": []any{"internal"}},
			expectError: false,
		},
		{
			name:        "Invalid JSON",
			input:       ptr.String(`{invalid`),
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.ConvertStringToStoredSource(tc.input)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestBuildFields(t *testing.T) {
	testCases := []struct {
		model       *resource.Model
		expected    *[]any
		name        string
		expectError bool
	}{
		{
			name:        "Nil fields",
			model:       &resource.Model{Fields: nil},
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Empty fields",
			model:       &resource.Model{Fields: ptr.String("")},
			expected:    nil,
			expectError: false,
		},
		{
			name: "Valid fields array",
			model: &resource.Model{
				Fields: ptr.String(`[{"type":"vector","path":"plot_embedding","numDimensions":1536,"similarity":"euclidean"}]`),
			},
			expected: &[]any{
				map[string]any{
					"type":          "vector",
					"path":          "plot_embedding",
					"numDimensions": float64(1536),
					"similarity":    "euclidean",
				},
			},
			expectError: false,
		},
		{
			name: "Invalid JSON fields",
			model: &resource.Model{
				Fields: ptr.String(`[invalid json]`),
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.BuildFields(tc.model)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestBuildAnalyzers(t *testing.T) {
	testCases := []struct {
		model       *resource.Model
		name        string
		expectError bool
		expectNil   bool
	}{
		{
			name:        "Empty analyzers",
			model:       &resource.Model{Analyzers: []resource.ApiAtlasFTSAnalyzersViewManual{}},
			expectError: false,
			expectNil:   true,
		},
		{
			name: "Valid analyzer with char and token filters",
			model: &resource.Model{
				Analyzers: []resource.ApiAtlasFTSAnalyzersViewManual{
					{
						Name:         ptr.String("myAnalyzer"),
						CharFilters:  []string{`{"type":"icuNormalize"}`},
						TokenFilters: []string{`{"type":"lowercase"}`},
						Tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
							Type: ptr.String("standard"),
						},
					},
				},
			},
			expectError: false,
			expectNil:   false,
		},
		{
			name: "Invalid char filter JSON",
			model: &resource.Model{
				Analyzers: []resource.ApiAtlasFTSAnalyzersViewManual{
					{
						Name:        ptr.String("myAnalyzer"),
						CharFilters: []string{`invalid json`},
						Tokenizer: &resource.ApiAtlasFTSAnalyzersTokenizer{
							Type: ptr.String("standard"),
						},
					},
				},
			},
			expectError: true,
			expectNil:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.BuildAnalyzers(tc.model)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.expectNil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
				}
			}
		})
	}
}

func TestBuildSynonyms(t *testing.T) {
	testCases := []struct {
		model     *resource.Model
		name      string
		expectNil bool
	}{
		{
			name:      "Empty synonyms",
			model:     &resource.Model{Synonyms: []resource.ApiAtlasFTSSynonymMappingDefinitionView{}},
			expectNil: true,
		},
		{
			name: "Valid synonyms",
			model: &resource.Model{
				Synonyms: []resource.ApiAtlasFTSSynonymMappingDefinitionView{
					{
						Analyzer: ptr.String("lucene.standard"),
						Name:     ptr.String("mySynonyms"),
						Source: &resource.SynonymSource{
							Collection: ptr.String("synonyms"),
						},
					},
				},
			},
			expectNil: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.BuildSynonyms(tc.model)
			require.NoError(t, err)
			if tc.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Len(t, *result, 1)
			}
		})
	}
}

func TestBuildStoredSource(t *testing.T) {
	testCases := []struct {
		expected    any
		model       *resource.Model
		name        string
		expectError bool
	}{
		{
			name:        "Nil stored source",
			model:       &resource.Model{StoredSource: nil},
			expected:    nil,
			expectError: false,
		},
		{
			name:        "Boolean true",
			model:       &resource.Model{StoredSource: ptr.String("true")},
			expected:    true,
			expectError: false,
		},
		{
			name:        "Boolean false",
			model:       &resource.Model{StoredSource: ptr.String("false")},
			expected:    false,
			expectError: false,
		},
		{
			name:        "JSON object",
			model:       &resource.Model{StoredSource: ptr.String(`{"include":["name"]}`)},
			expected:    map[string]any{"include": []any{"name"}},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.BuildStoredSource(tc.model)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestBuildTypeSets(t *testing.T) {
	testCases := []struct {
		model     *resource.Model
		name      string
		expectNil bool
	}{
		{
			name:      "Empty type sets",
			model:     &resource.Model{TypeSets: []resource.TypeSet{}},
			expectNil: true,
		},
		{
			name: "Valid type sets",
			model: &resource.Model{
				TypeSets: []resource.TypeSet{
					{
						Name:  ptr.String("myTypeSet"),
						Types: ptr.String(`[{"type":"string"}]`),
					},
				},
			},
			expectNil: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resource.BuildTypeSets(tc.model)
			require.NoError(t, err)
			if tc.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Len(t, *result, 1)
			}
		})
	}
}
