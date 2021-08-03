package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/db"
	_ "go-learn/db"
	"go-learn/model"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		user := model.User{Username: "hello" + time.Now().String(), Password: "test"}
		_ = db.GetDB().Create(&user) // 通过数据的指针来创建
		c.JSON(200, user)
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		return
	}
}
