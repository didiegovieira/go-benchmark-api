package entity

import "time"

type Type string

const (
	SortingAlgorithm Type = "sorting_algorithm"
	Serialization    Type = "serialization"
)

type Benchmark struct {
	Type     Type       `json:"type"`
	Data     []int      `json:"data"`
	Date     *time.Time `json:"date"`
	Duration string     `json:"duration"`
}
