package dao

import "time"

type LoginInput struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserViewModel struct {
	ID        uint      `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Token     string    `json:"token,omitempty"`
	IsActive  bool      `json:"active"`
}
