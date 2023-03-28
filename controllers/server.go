package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/config"
)

// CreateRouter creates and configure a server
func CreateRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(DB(db))
	setupRoutes(router)
	return router
}

// StartServer start given server
func StartServer(router *gin.Engine) {
	// Read configurations from environmental variables.
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Fatalf("failed to read env variables: #{err}\n")
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + env.HTTPPort,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
