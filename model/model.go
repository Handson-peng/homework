package model

import "github.com/Handson-peng/homework/database"

type UserMessage struct{
	UserId string
	UserName string
	Message string
}

type RespondMessage struct {
	UserId string
	UserName string
	Message []string
}

type SendMessage struct{
	UserId string `json:"userid"`
	Text string   `json:"text"`
}

type UserInfo struct {
	UserId string
	UserName string
}

func SaveUserMessage(user UserMessage){
	database.Insert(user)
}

func QueryByName(username string) RespondMessage{
	var result RespondMessage
	var um UserMessage
	q := database.Query[UserMessage]("username", username)
	for _, um = range q {
		result.Message = append(result.Message, um.Message)
	}
	result.UserId = um.UserId
	result.UserName = um.UserName
	return result
}

func QueryById(userid string) RespondMessage{
	var result RespondMessage
	var um UserMessage
	q := database.Query[UserMessage]("userid", userid)
	for _, um = range q {
		result.Message = append(result.Message, um.Message)
	}
	result.UserId = um.UserId
	result.UserName = um.UserName
	return result
}

func QueryAllUser() []UserInfo{
	var result []UserInfo
	usermap := make(map[UserInfo]int)
	q := database.Query[UserMessage]("userid", "")
	for _, um := range q {
		user := UserInfo{UserId: um.UserId, UserName: um.UserName}
		if _, exist := usermap[user]; !exist{
			usermap[user] = 1
			result = append(result, user)
		}
	}
	return result
}
