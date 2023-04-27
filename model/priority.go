package model

type Priority struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	List   uint64 `json:"list"`
	Weight int    `json:"weight"`
}
