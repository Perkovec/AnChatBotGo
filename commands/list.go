package commands

import (
	"fmt"
	"time"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) ListCommand(msg tgapi.Message) {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if user.IsChatUser {
		users, err := app.GetRecentChatUsers()
		if err != nil {
			fmt.Println("Error (ListCommand/GetRecentChatUsers): ", err)
		}
		var list string
		for _, usr := range users {
			diff := int(time.Now().Unix()) - usr.LastMessage
			if diff < 60*3 {
				list += fmt.Sprintf(i18n.RuLocal["list_item_online"], usr.ChatID, usr.Name)
			} else {
				list += fmt.Sprintf(i18n.RuLocal["list_item"], usr.ChatID, usr.Name, TimeDiff(diff))
			}
		}

		msg.SendMessage(app.Bot, fmt.Sprintf(i18n.RuLocal["list"], list), nil)

		app.UpdateUserLastMessage(msg.From.ID)
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}
