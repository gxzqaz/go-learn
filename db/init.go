package db

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database instance
var globalDB *gorm.DB

func GetDB() *gorm.DB {
	return globalDB
}

//到时候改为从nacos中获取就可以了
//go:embed mysql.json
var configStr []byte

type configuration struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	config := configuration{}
	if err := json.Unmarshal(configStr, &config); err != nil {
		// do nothing
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print()
	}
	globalDB = db
}
