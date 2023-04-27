package model

type List struct {
	ListHeader
	Elements []Element `json:"elements"`
}
