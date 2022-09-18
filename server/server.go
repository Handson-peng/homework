package server

import (
	"github.com/Handson-peng/homework/database"
	"github.com/Handson-peng/homework/service"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	go service.InitLinebot(viper.GetString("lineBot.channelSecret"), viper.GetString("lineBot.channelAccessToken"))
	go database.Connect(viper.GetString("mongo.uri"), viper.GetString("mongo.database"), viper.GetString("mongo.collection"))
	router := gin.Default()
	router.POST("/callback", service.LineCallback)
	router.GET("/user", service.GetAllUser)
	router.POST("/user", service.PostUser)
	router.Run(":80")
}
