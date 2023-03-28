package controllers

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	// It is good practice to version your API from the start
	v1 := router.Group("/v1")

	v1.GET("/health", health)

	// STEP 7: ...and bind the handler to the endpoint
	video := v1.Group("/files")
	video.GET("", getAllVideo)
	video.POST("", uploadVideo)
	video.DELETE("/:fileid", deleteVideo)
	video.GET("/:fileid", getVideo)
}
