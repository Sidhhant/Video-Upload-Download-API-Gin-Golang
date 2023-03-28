package controllers

import (
	"log"

	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DB middleware attaches a database connection to gin's Context
func DB(db *gorm.DB) gin.HandlerFunc {
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Fatal("failed to read env variables: ", err)
	}

	return func(c *gin.Context) {
		c.Set(env.ContextDB, db)
		c.Next()
	}
}
