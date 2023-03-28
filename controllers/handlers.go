package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/config"
	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/models"
)

func getDB(c *gin.Context) *gorm.DB {
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Fatal("failed to read env: ", err)
	}

	return c.MustGet(env.ContextDB).(*gorm.DB)
}

func health(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
}

// STEP 6: Let's make a handler
func getAllVideo(c *gin.Context) {
	db := getDB(c)
	var files []models.UploadedFile
	err := db.Where("isdeleted = false").Find(&files).Error
	if err != nil {
		log.Println("not able to getfile, getAllVideo", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Println("get list of video success!")
	c.JSON(http.StatusOK, files)
}

func uploadVideo(c *gin.Context) {
	db := getDB(c)

	file, err := c.FormFile("data")
	if err != nil {
		log.Println("file not found in form data", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	extension := filepath.Ext(file.Filename)
	if !(extension == ".mp4" || extension == ".mpg") {
		log.Println("wrong type of media uploaded.")
		c.Status(http.StatusUnsupportedMediaType)
		return
	}

	/*
		TODO: Use os.Stat to check already existing file and return 409
		TODO: Use in memory file storage in place of ./saved folder because of better API availability.
	*/

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./saved", os.ModePerm)
	if err != nil {
		log.Println("not able to create /saved folder")
		c.Status(http.StatusInternalServerError)
		return
	}

	uuid := uuid.New().String()
	err = c.SaveUploadedFile(file, "saved/"+uuid)
	if err != nil {
		log.Println("not able to save file in directory", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	uplFile := models.UploadedFile{
		FileID:    uuid,
		Name:      file.Filename,
		Size:      file.Size,
		CreatedAt: time.Now(),
		Type:      extension,
		Isdeleted: false,
	}
	err = db.Create(uplFile).Error
	if err != nil {
		log.Println("not able to create entry in database, upload", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// File saved successfully. Return proper result
	c.Header("Location", "http://localhost:8080/v1/saved/"+uuid)
	log.Println("file saved successfully")
	c.Status(http.StatusCreated)

}

func deleteVideo(c *gin.Context) {
	fileid := c.Param("fileid")
	log.Println(fileid)
	db := getDB(c)
	var file models.UploadedFile
	err := db.First(&file, "file_id = ? AND isdeleted = false", string(fileid)).Error
	if err != nil {
		log.Println("failed to get file, deleted: ", err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	err = db.Model(&file).Where("file_id=?", string(fileid)).Update("isdeleted", true).Error
	if err != nil {
		log.Println("failed to save soft delete file", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	log.Println("file delete success!")
	c.Status(http.StatusNoContent)
}

func getVideo(c *gin.Context) {
	fileid := c.Param("fileid")
	db := getDB(c)
	var file models.UploadedFile

	err := db.Where("file_id = ? AND isdeleted = false", fileid).Find(&file).Error
	if err != nil {
		log.Println("failed to get file, getVideo", err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	if file.Type == ".mpg" {
		c.Header("Content-Type", "video/mpeg")
	} else {
		c.Header("Content-Type", "video/mp4")
	}

	if _, err := os.Stat("./saved/" + fileid); errors.Is(err, os.ErrNotExist) {
		log.Println("file doesn't exist, getVideo", err.Error())
		c.Status(http.StatusNotFound)
		return
	}

	log.Println("file found success!")
	c.Header("Content-Description", fileid)
	c.FileAttachment("./saved/"+fileid, file.Name)
	c.Status(http.StatusOK)
}
