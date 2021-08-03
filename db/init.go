package db

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learn/constant"
	"go-learn/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Database instance
var globalDB *gorm.DB

// get db
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if gin.Mode() == constant.GinRelease {
	}

	// 迁移 schema
	_ = db.AutoMigrate(&model.User{})
	// 设置连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	globalDB = db
}
