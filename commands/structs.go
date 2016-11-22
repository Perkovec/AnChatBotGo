package commands

import (
	"github.com/perkovec/AnChatBotGo/tgapi"
	"gopkg.in/mgo.v2"
)

type App struct {
	Bot       *tgapi.Bot
	UsersC    *mgo.Collection
	MessagesC *mgo.Collection
}

type User struct {
	Banned      bool   `bson:"banned"`
	Hidden      bool   `bson:"hidden"`
	ChatID      string `bson:"chat_id"`
	IsChatUser  bool   `bson:"is_chat_user"`
	LastMessage int    `bson:"last_message"`
	MuteEndTime int    `bson:"mute_end_time"`
	Muted       bool   `bson:"muted"`
	Name        string `bson:"name"`
	StartDate   int    `bson:"start_date"`
	TgID        int    `bson:"tg_id"`
}

type Replies struct {
	Replies []struct {
		TgID  int `bson:"tg_id"`
		MsgID int `bson:"msg_id"`
	} `bson:"replies"`
}
