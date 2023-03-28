package db

import (
	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/models"
	"github.com/jinzhu/gorm"
)

// DB is the database connection
var DB *gorm.DB

// Init sets the given database conncetion as the de-facto conncetion for this app
func Init(db *gorm.DB) {
	DB = db

	// STEP 2: We need to let gorm automigrate and create example table if it does exist
	DB.AutoMigrate(&models.UploadedFile{})
}
