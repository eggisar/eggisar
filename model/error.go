package model

type Error struct {
	Message
	VoidAuth bool `json:"voidAuth"`
}
