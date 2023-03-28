package main

import (
	"log"
	"time"

	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/config"
	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/controllers"
	"github.com/Sidhhant/Video-Upload-Download-API-Gin-Golang/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var err error
	env, err := config.ReadFromEnv()
	if err != nil {
		log.Fatal("failed to read env: ", err)
	}

	connection := env.DBUser + ":" + env.DBPassword + "@tcp(db:3306)/" + env.ContextDB + "?charset=utf8&parseTime=True&loc=Local"
	count := 100
	var database *gorm.DB
	for count > 1 {
		database, err = gorm.Open(
			"mysql",
			connection,
		)
		if err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			log.Println(err)
			continue
		}
		defer database.Close()
		break
	}

	// defer database.Close()
	db.Init(database)

	router := controllers.CreateRouter(database)
	controllers.StartServer(router)
}
