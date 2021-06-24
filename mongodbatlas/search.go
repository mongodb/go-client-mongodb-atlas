// Copyright 2021 MongoDB Inc
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

package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const (
	searchBasePath = "api/atlas/v1.0/groups/%s/clusters/%s/fts"
)

// SearchService provides access to the search related functions in the Atlas API.
//
// See more: https://docs.atlas.mongodb.com/reference/api/atlas-search/
type SearchService interface {
	ListIndexes(ctx context.Context, groupID string, clusterName string, databaseName string, collectionName string, opts *ListOptions) ([]*SearchIndex, *Response, error)
	GetIndex(ctx context.Context, groupID, clusterName, indexID string) (*SearchIndex, *Response, error)
	CreateIndex(ctx context.Context, projectID, clusterName string, r *SearchIndex) (*SearchIndex, *Response, error)
	UpdateIndex(ctx context.Context, projectID, clusterName, indexID string, r *SearchIndex) (*SearchIndex, *Response, error)
	DeleteIndex(ctx context.Context, projectID, clusterName, indexID string) (*Response, error)
	ListAnalyzers(ctx context.Context, groupID, clusterName string, listOptions *ListOptions) ([]*SearchAnalyzer, *Response, error)
	UpdateAllAnalyzers(ctx context.Context, groupID, clusterName string, analyzers []*SearchAnalyzer) ([]*SearchAnalyzer, *Response, error)
}

// SearchServiceOp provides an implementation of the SearchService interface.
type SearchServiceOp service

var _ SearchService = &SearchServiceOp{}

// ListIndexes Get all Atlas Search indexes for a specified collection.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/
func (s *SearchServiceOp) ListIndexes(ctx context.Context, groupID, clusterName, databaseName, collectionName string, opts *ListOptions) ([]*SearchIndex, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("GroupID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("ClusterName", "must be set")
	}
	if databaseName == "" {
		return nil, nil, NewArgError("databaseName", "must be set")
	}
	if collectionName == "" {
		return nil, nil, NewArgError("collectionName", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, groupID, clusterName)
	path = fmt.Sprintf("%s/indexes/%s/%s", path, databaseName, collectionName)

	path, err := setListOptions(path, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root []*SearchIndex
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// GetIndex gets one Atlas Search index by its indexId.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-one/
func (s *SearchServiceOp) GetIndex(ctx context.Context, groupID, clusterName, indexID string) (*SearchIndex, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
	}
	if indexID == "" {
		return nil, nil, NewArgError("indexID", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, groupID, clusterName)
	path = fmt.Sprintf("%s/indexes/%s", path, indexID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root *SearchIndex
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// CreateIndex creates an Atlas Search index.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-indexes-create-one/
func (s *SearchServiceOp) CreateIndex(ctx context.Context, projectID, clusterName string, r *SearchIndex) (*SearchIndex, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, projectID, clusterName)
	path = fmt.Sprintf("%s/indexes", path)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, r)
	if err != nil {
		return nil, nil, err
	}

	var root *SearchIndex
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// UpdateIndex updates an Atlas Search index by its indexId.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-indexes-update-one/
func (s *SearchServiceOp) UpdateIndex(ctx context.Context, projectID, clusterName, indexID string, r *SearchIndex) (*SearchIndex, *Response, error) {
	if projectID == "" {
		return nil, nil, NewArgError("projectID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
	}
	if indexID == "" {
		return nil, nil, NewArgError("indexID", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, projectID, clusterName)
	path = fmt.Sprintf("%s/indexes/%s", path, indexID)

	req, err := s.Client.NewRequest(ctx, http.MethodPatch, path, r)
	if err != nil {
		return nil, nil, err
	}

	var root *SearchIndex
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// DeleteIndex deletes one Atlas Search index by its indexId.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-indexes-delete-one/
func (s *SearchServiceOp) DeleteIndex(ctx context.Context, projectID, clusterName, indexID string) (*Response, error) {
	if projectID == "" {
		return nil, NewArgError("projectID", "must be set")
	}
	if clusterName == "" {
		return nil, NewArgError("clusterName", "must be set")
	}
	if indexID == "" {
		return nil, NewArgError("indexID", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, projectID, clusterName)
	path = fmt.Sprintf("%s/indexes/%s", path, indexID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.Client.Do(ctx, req, nil)

	return resp, err
}

// ListAnalyzers gets all Atlas Search user-defined analyzers for a specified cluster.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-analyzers-get-all/
func (s *SearchServiceOp) ListAnalyzers(ctx context.Context, groupID, clusterName string, listOptions *ListOptions) ([]*SearchAnalyzer, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, groupID, clusterName)
	path = fmt.Sprintf("%s/analyzers", path)

	path, err := setListOptions(path, listOptions)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root []*SearchAnalyzer
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// UpdateAllAnalyzers Update All User-Defined Analyzers for a specific Cluster.
//
// See more: https://docs.atlas.mongodb.com/reference/api/fts-analyzers-update-all//
func (s *SearchServiceOp) UpdateAllAnalyzers(ctx context.Context, groupID, clusterName string, analyzers []*SearchAnalyzer) ([]*SearchAnalyzer, *Response, error) {
	if groupID == "" {
		return nil, nil, NewArgError("groupID", "must be set")
	}
	if clusterName == "" {
		return nil, nil, NewArgError("clusterName", "must be set")
	}

	path := fmt.Sprintf(searchBasePath, groupID, clusterName)
	path = fmt.Sprintf("%s/analyzers", path)

	req, err := s.Client.NewRequest(ctx, http.MethodPut, path, analyzers)
	if err != nil {
		return nil, nil, err
	}

	var root []*SearchAnalyzer
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// SearchIndex index definition.
type SearchIndex struct {
	Analyzer       string            `json:"analyzer,omitempty"`
	Analyzers      []*CustomAnalyzer `json:"analyzers,omitempty"` // Custom analyzers
	CollectionName string            `json:"collectionName"`
	Database       string            `json:"database"`
	IndexID        string            `json:"indexID,omitempty"`
	Mappings       *IndexMapping     `json:"mappings,omitempty"`
	Name           string            `json:"name"`
	SearchAnalyzer string            `json:"searchAnalyzer,omitempty"`
	Status         string            `json:"status,omitempty"`
}

// IndexMapping containing index specifications for the collection fields.
type IndexMapping struct {
	Dynamic bool                   `json:"dynamic"`
	Fields  *map[string]IndexField `json:"fields,omitempty"`
}

// IndexField field specifications.
type IndexField struct {
	Analyzer       string                 `json:"analyzer,omitempty"`
	Type           string                 `json:"type"`
	Tokenization   string                 `json:"tokenization,omitempty"` // edgeGram|nGram
	MinGrams       *int                   `json:"minGrams,omitempty"`
	MaxGrams       *int                   `json:"maxGrams,omitempty"`
	FoldDiacritics *bool                  `json:"foldDiacritics,omitempty"`
	Fields         *map[string]IndexField `json:"fields,omitempty"`
	SearchAnalyzer string                 `json:"searchAnalyzer,omitempty"`
	IndexOptions   string                 `json:"indexOptions,omitempty"` // docs|freqs|positions
	Store          *bool                  `json:"store,omitempty"`
	IgnoreAbove    *int                   `json:"ignoreAbove,omitempty"`
	Norms          string                 `json:"norms,omitempty"` // include|omit
	Dynamic        *bool                  `json:"dynamic,omitempty"`
	Representation string                 `json:"representation,omitempty"`
	IndexIntegers  *bool                  `json:"indexIntegers,omitempty"`
	IndexDoubles   *bool                  `json:"indexDoubles,omitempty"`
	IndexShapes    *bool                  `json:"indexShapes,omitempty"`
}

// SearchAnalyzer search analyzer definition.
type SearchAnalyzer struct {
	BaseAnalyzer     string   `json:"baseAnalyzer"`
	MaxTokenLength   *int     `json:"maxTokenLength,omitempty"`
	IgnoreCase       *bool    `json:"ignoreCase,omitempty"`
	Name             string   `json:"name"`
	StemExclusionSet []string `json:"stemExclusionSet,omitempty"`
	Stopwords        []string `json:"stopwords,omitempty"`
}

// CustomAnalyzer custom analyzer for index.
type CustomAnalyzer struct {
	// Name of the custom analyzer. Names must be unique within an index, and may not start with any of the following strings: lucene, builtin, mongodb
	Name string `json:"name"`

	// CharFilters Array containing zero or more character filters.
	CharFilters []*AnalyzerCharFilter `json:"charFilters,omitempty"`

	// Tokenizer to use.
	Tokenizer *AnalyzerTokenizer `json:"tokenizer"`

	// TokenFilters Array containing zero or more token filters
	TokenFilters []*AnalyzerTokenFilters `json:"tokenFilters,omitempty"`
}

// AnalyzerCharFilter Characters filters for custom analyzer. For further information, go to
// https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-char-filters-ref
type AnalyzerCharFilter struct {
	// Type The type of this character filter, supports: htmlStrip, icuNormalize, mapping, persian
	Type string `json:"type"`

	// IgnoreTags A list of HTML tags to exclude from filtering. Apply for type: htmlStrip
	IgnoreTags []string `json:"IgnoreTags,omitempty"`

	// Mappings An object containing a comma-separated list of mappings. A mapping indicates that one character
	// or group of characters should be substituted for another, in the format <original> : <replacement>
	// apply for type: mapping
	Mappings *map[string]string `json:"mappings,omitempty"`
}

// AnalyzerTokenizer tokenizer for custom analyzer, for further information, go to
// https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-tokenizers-ref
type AnalyzerTokenizer struct {
	// Type The type of this tokenizer. Supports standard, keyword, whitespace, nGram, edgeGram, regexCaptureGroup,
	// regexSplit, uaxUrEmail
	Type string `json:"type"`

	// MaxTokenLength Maximum length for a single token. Tokens greater than this length are split at maxTokenLength
	// into multiple tokens. Apply for type: standard, whitespace, uaxUrlEmail
	MaxTokenLength *int `json:"maxTokenLength,omitempty"`

	// MinGram Number of characters to include in the shortest token created. Apply for type: nGram, edgeGram
	MinGram *int `json:"minGram,omitempty"`

	// MaxGram Number of characters to include in the longest token created. Apply for type: nGram, edgeGram
	MaxGram *int `json:"maxGram,omitempty"`

	// Pattern A regular expression to match against. Apply for type: regexCaptureGroup, regexSplit
	Pattern string `json:"pattern,omitempty"`

	// Group Index of the character group within the matching expression to extract into tokens. Use 0 to extract all
	// character groups. Apply for type: regexCaptureGroup
	Group *int `json:"group,omitempty"`
}

// AnalyzerTokenFilters token filter for custom analyzer. To get more information, go to
// https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-token-filters-ref
type AnalyzerTokenFilters struct {
	// Type The type of this token filter. Supports: daitchMokotoffSoundex, lowercase, length, icuFolding, icuNormalize
	// nGram, edgeGram, shingle, regex, snowballStemming, stopword, trim
	Type string `json:"type"`

	// OriginalTokens Specifies whether to include or omit the original tokens in the output of the token filter. Value can
	// be one of the following: include - to include the original tokens with the encoded tokens in the output of the token
	// filter. We recommend this value if you want queries on both the original tokens as well as the encoded forms.
	// omit - to omit the original tokens and include only the encoded tokens in the output of the token filter.
	// Use this value if you want to only query on the encoded forms of the original tokens. Apply for type: daitchMokotoffSoundex
	OriginalTokens string `json:"originalTokens,omitempty"`

	// Min The minimum length of a token. Must be less than or equal to max. Apply for type: length
	Min *int `json:"min,omitempty"`

	// Max The maximum length of a token. Must be greater than or equal to min.
	Max *int `json:"max,omitempty"`

	// NormalizationForm Normalization form to apply. Accepted values are:
	// nfd (Canonical Decomposition)
	// nfc (Canonical Decomposition, followed by Canonical Composition)
	// nfkd (Compatibility Decomposition)
	// nfkc (Compatibility Decomposition, followed by Canonical Composition).
	// Apply for type: icuNormalize
	NormalizationForm string `json:"normalizationForm,omitempty"`

	// MinGram The minimum length of generated n-grams. Must be less than or equal to maxGram. Apply for type: nGram, edgeGram
	MinGram *int `json:"minGram,omitempty"`

	// MaxGram The maximum length of generated n-grams. Must be greater than or equal to minGram. Apply for type: nGram, edgeGram
	MaxGram *int `json:"maxGram,omitempty"`

	// TermsNotInBounds Accepted values are: include, omit
	// If include is specified, tokens shorter than minGram or longer than maxGram are indexed as-is. If omit is specified,
	// those tokens are not indexed. Apply for type: nGram, edgeGram
	TermsNotInBounds string `json:"termsNotInBounds,omitempty"`

	// MinShingleSize Minimum number of tokens per shingle. Must be less than or equal to maxShingleSize. Apply for type: shingle
	MinShingleSize *int `json:"minShingleSize,omitempty"`

	// MaxShingleSize Maximum number of tokens per shingle. Must be greater than or equal to minShingleSize. Apply for type: shingle
	MaxShingleSize *int `json:"maxShingleSize,omitempty"`

	// Pattern Regular expression pattern to apply to each token. Apply for type: regex
	Pattern string `json:"pattern,omitempty"`

	// Replacement Replacement string to substitute wherever a matching pattern occurs. Apply for type: regex
	Replacement string `json:"replacement,omitempty"`

	// Matches Acceptable values are: all, first
	// If matches is set to all, replace all matching patterns. Otherwise, replace only the first matching pattern. Apply for type: regex
	Matches string `json:"matches,omitempty"`

	// StemmerName Apply for type: snowballStemming
	StemmerName string `json:"stemmerName,omitempty"`

	// Tokens The list of stop words that correspond to the tokens to remove. Value must be one or more stop words. Apply for type: stopword
	Tokens []string `json:"tokens,omitempty"`

	// IgnoreCase The flag that indicates whether or not to ignore case of stop words when filtering the tokens to remove. The value can be one of the following:
	// true - to ignore case and remove all tokens that match the specified stop words
	// false - to be case-sensitive and remove only tokens that exactly match the specified case
	// If omitted, defaults to true. Apply for type: stopword
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}
