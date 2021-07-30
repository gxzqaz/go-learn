package model

import (
	"gorm.io/gorm"
)

// gorm.Model 的定义
type User struct {
	gorm.Model
	Username string ``
	Password string ``
}
