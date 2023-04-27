package model

import "time"

type ListHeader struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Note       string    `json:"note"`
	Updated    time.Time `json:"updated"`
	UpdatedBy  User      `json:"updated-by"`
	Owner      User      `json:"owner"`
	Users      []User
	Priorities []Priority
}
