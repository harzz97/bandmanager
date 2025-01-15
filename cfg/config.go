package cfg

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/bandmanager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < 5; i++ {
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Fatalf("failed to connect to the database after multiple retries: %v", err)
		}
	}
	return db
}

func HandleError(c *gin.Context, err error, message string, statusCode int) {
	if err != nil {
		c.JSON(statusCode, gin.H{"error": message, "details": err.Error()})
	} else {
		c.JSON(statusCode, gin.H{"error": message})
	}
}
