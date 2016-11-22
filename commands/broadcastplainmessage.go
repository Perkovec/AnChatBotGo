package commands

import (
	"fmt"

	"github.com/perkovec/AnChatBotGo/tgapi"
	"gopkg.in/mgo.v2/bson"
)

func (app *App) BroadcastPlainMessage(text string, excludeid int, parsemode string, selfid int, selfmsgid int) {
	users, err := app.GetChatUsers()
	if err != nil {
		fmt.Println("Error (BroadcastPlainMessage/GetChatUsers): ", err)
	}
	documentReplies := []bson.M{}

	for _, user := range users {
		if user.TgID != excludeid {
			message, err := app.Bot.SendMessage(user.TgID, text, &tgapi.SendMessageOpts{
				ParseMode: parsemode,
			})
			if err != nil {
				fmt.Println("Error (BroadcastPlainMessage/SendMessage): ", err)
			}
			documentReplies = append(documentReplies, bson.M{
				"tg_id":  user.TgID,
				"msg_id": message.MessageID,
			})
		}
	}
	if selfid > 0 && selfmsgid > 0 {
		documentReplies = append(documentReplies, bson.M{
			"tg_id":  selfid,
			"msg_id": selfmsgid,
		})
	}
	err = app.MessagesC.Insert(bson.M{
		"replies": documentReplies,
	})
	if err != nil {
		fmt.Println("Error (BroadcastPlainMessage/Insert): ", err)
	}
}
