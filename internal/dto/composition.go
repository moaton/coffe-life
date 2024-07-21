package dto

type Composition struct {
	Name   string `json:"name"`
	Weight int64  `json:"weight"`
	Unit   string `json:"unit"`
}

type Compositions []Composition
