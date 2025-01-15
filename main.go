package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harzz97/bandmanager/api"
	"github.com/harzz97/bandmanager/cfg"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	db := cfg.InitDB()
	r := gin.Default()

	// Load routes
	// api.AdminRoutes(r)
	api.MasterListRoutes(r, db)
	api.SetlistRoutes(r, db)
	// api.AudienceRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
