package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Emoji struct {
	gorm.Model
	Id        string
	Name      string
	UserId    string
	ImageUrl  string
	CreatedAt time.Time
}

func main() {
	setDatabaseConfig()
	setRouterConfig()
}

func setDatabaseConfig() *gorm.DB {
	dsn := "root:lostfinder123@tcp(127.0.0.1:3306)/jjal_talk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connecting Database Failed")
	}

	db.AutoMigrate(&Emoji{})

	return db
}

func setRouterConfig() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	router.GET("/hongseok", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello": "Hongseok",
		})
	})

	router.Run(":8000")
}
