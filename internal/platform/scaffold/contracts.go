// Package scaffold contains contracts shared by the feature scaffolds.
package scaffold

import "errors"

// ErrNotImplemented marks a business operation awaiting implementation.
var ErrNotImplemented = errors.New("operation is not implemented")

type PageQuery struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type Page[T any] struct {
	Items  []T `json:"items"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Problem struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldIssue `json:"fields,omitempty"`
}

type FieldIssue struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
