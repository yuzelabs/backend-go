package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Avatar    string    `json:"avatar"`
	Address   string    `json:"address,omitempty" gorm:"index"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
