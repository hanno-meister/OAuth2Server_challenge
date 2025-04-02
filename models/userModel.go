package models

import "gorm.io/gorm"

// Defines a user table using GORMs predefined model struct and custom fields for basic authentication
type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
