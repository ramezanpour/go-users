package models

import "github.com/jinzhu/gorm"

// User stores information about a user
type User struct {
	gorm.Model
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	IsActive  bool   `json:"active,omitempty" gorm:"default:true"`
}
