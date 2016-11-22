package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/perkovec/AnChatBotGo/commands"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

const token = "294497096:AAEHhOZV2EtQbmjDwjbVQB7KFAaAJeHAcxI"

func main() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("MongoDB connected")

	bot, err := tgapi.Create(token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("@" + bot.Info.Username)

	DB := session.DB("AnChatBot")

	myApp := commands.App{
		Bot:       bot,
		UsersC:    DB.C("users"),
		MessagesC: DB.C("messages"),
	}

	messages := make(chan tgapi.Message)
	updates := make(chan tgapi.Message)

	bot.OnError = func(msg tgapi.Message, err string) {
		ProcessError(myApp, msg, err)
	}

	go bot.Polling(messages, updates)
	go newUpdates(myApp, updates)

	for message := range messages {
		ProcessMsg(myApp, message)
	}
}

func newUpdates(myApp commands.App, updates chan tgapi.Message) {
	for update := range updates {
		ProcessUpdate(myApp, update)
	}
}
