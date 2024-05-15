package dto

import "errors"

var ErrEmptyArray = errors.New("array cannot be empty")

type RequestInput struct {
	Arr []int `json:"arr"`
}

// Validate checks if the array is empty
func (r *RequestInput) Validate() bool {
	return len(r.Arr) == 0
}
