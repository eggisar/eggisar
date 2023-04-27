package model

import "time"

type Element struct {
	ID        uint64    `json:"id"`
	List      uint64    `json:"list"`
	Name      string    `json:"name"`
	Priority  Priority  `json:"priority"`
	Checked   bool      `json:"checked"`
	Updated   time.Time `json:"updated"`
	UpdatedBy User      `json:"updated-by"`
}
