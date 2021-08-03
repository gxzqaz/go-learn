package model

import "time"

type BaseModel struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreateTime time.Time `gorm:"autoCreateTime;index;" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime;" json:"updateTime"`
	Deleted    uint8     `gorm:"default:0;not null" json:"deleted"`
}
