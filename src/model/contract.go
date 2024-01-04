package model

import "time"

type Contract struct {
	ID        uint       `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
}
