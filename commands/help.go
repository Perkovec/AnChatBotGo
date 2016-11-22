package commands

import (
	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) HelpCommand(msg tgapi.Message) {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if user.IsChatUser {
		msg.SendMessage(app.Bot, i18n.RuLocal["help"], nil)
		app.UpdateUserLastMessage(msg.From.ID)
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}
