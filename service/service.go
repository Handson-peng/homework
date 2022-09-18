package service

import (
	. "github.com/Handson-peng/homework/model"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
)

var Linebot *linebot.Client
var err error

func InitLinebot(secret, accessToken string) {
	Linebot, err = linebot.New(secret, accessToken)
	if err != nil {
		log.Fatal(err)
	}
}

func LineCallback(c *gin.Context) {
	events, err := Linebot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, "Error: "+ err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, "Error: "+ err.Error())
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				r, _ := Linebot.GetProfile(event.Source.UserID).Do()
				SaveUserMessage(UserMessage{UserId: r.UserID, UserName: r.DisplayName, Message: message.Text})
			}
		}
	}
}