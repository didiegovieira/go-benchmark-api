package dto

import "errors"

var ErrEmptyArray = errors.New("array cannot be empty")

type SortingInput struct {
	Arr []int `json:"arr"`
}

func (r *SortingInput) Validate() bool {
	return len(r.Arr) > 0
}
