package dto

type TimeCalculateInput struct {
	Func func() `json:"func"`
	Name string `json:"name"`
}
