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

func GetAllUser(c *gin.Context) {
	rs := QueryAllUser()
	if rs == nil {
		c.JSON(http.StatusOK, "Error: there is no user in DB")
	}
	c.JSON(http.StatusOK, rs)
}

func PostUser(c *gin.Context) {
	var json SendMessage
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, "Error: " + err.Error())
		return
	}
	if _, err = Linebot.PushMessage(json.UserId, linebot.NewTextMessage(json.Text)).Do(); err != nil {
		c.JSON(http.StatusInternalServerError, "Error: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, "Line text successfully send.")
}

