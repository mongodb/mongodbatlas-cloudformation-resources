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

// TestConvertToAnySlice tests JSON string array to any slice conversion
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

// TestNewTokenizerModel tests tokenizer struct to map conversion (signature changed in commit)
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

// TestConvertStringToStoredSource tests NEW StoredSource conversion helper (added in commit 4c080db8)
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
