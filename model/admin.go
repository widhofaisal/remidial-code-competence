package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       string `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
