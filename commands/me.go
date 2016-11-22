package commands

import (
	"fmt"
	"html"
	"reflect"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) MeCommand(msg tgapi.Message, text string) {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if user.IsChatUser {
		app.BroadcastPlainMessage(
			fmt.Sprintf(i18n.RuLocal["me"], user.Name, html.EscapeString(text)),
			msg.From.ID,
			"HTML",
			msg.From.ID,
			msg.MessageID)

		app.UpdateUserLastMessage(msg.From.ID)
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}

func (app *App) GetMeText(msg tgapi.Message, text string) string {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if !reflect.DeepEqual(user, User{}) {
		return fmt.Sprintf(i18n.RuLocal["me"], user.Name, html.EscapeString(text))
	}
	return ""
}
