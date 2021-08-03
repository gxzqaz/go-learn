package model

// gorm.Model 的定义
type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;size:128" json:"username"`
	Password string `gorm:"size:128" json:"password"`
}
