package commands

import (
	"fmt"
	"time"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
	"gopkg.in/mgo.v2/bson"
)

func (app *App) StartCommand(msg tgapi.Message) {
	user, err := app.GetUserByTgID(msg.From.ID)
	if err != nil && err.Error() == "not found" {
		nickname := Nickname(2)
		app.createNewUser(msg, nickname)
		msg.SendMessage(app.Bot, fmt.Sprintf(i18n.RuLocal["start_new"], nickname), nil)
		app.BroadcastPlainMessage(fmt.Sprintf(i18n.RuLocal["new_user"], nickname), msg.From.ID, "", 0, 0)
	} else if err != nil {
		fmt.Println("Error (StartCommand/GetUserByTgID): ", err)
	} else if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if !user.IsChatUser {
		app.UsersC.Update(bson.M{"tg_id": user.TgID}, bson.M{
			"$set": bson.M{
				"is_chat_user": true,
				"last_message": int(time.Now().Unix()),
			},
		})
		message, err := msg.SendMessage(app.Bot, fmt.Sprintf(i18n.RuLocal["start"], user.Name), nil)
		if err != nil {
			fmt.Println("Error (StartCommand/SendMessage): ", err)
		}
		app.BroadcastPlainMessage(fmt.Sprintf(i18n.RuLocal["entry_user"], user.Name), msg.From.ID, "", msg.From.ID, message.MessageID)
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["already_in_chat"], nil)
	}
}

func (app *App) createNewUser(msg tgapi.Message, nickname string) error {
	user := User{}
	err := app.UsersC.Find(nil).Sort("-$natural").One(&user)
	var chat_id string
	if err != nil {
		chat_id = "A"
	} else {
		chat_id = NumberToLetters(LettersToNumber(user.ChatID) + 1)
	}

	unix := int(time.Now().Unix())

	return app.UsersC.Insert(User{
		Banned:      false,
		Hidden:      false,
		ChatID:      chat_id,
		IsChatUser:  true,
		LastMessage: unix,
		MuteEndTime: 0,
		Muted:       false,
		Name:        nickname,
		StartDate:   unix,
		TgID:        msg.From.ID,
	})
}
