package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) NickCommand(msg tgapi.Message, newNickname string) {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if user.IsChatUser {
		newNickname = strings.TrimSpace(newNickname)
		if len(newNickname) < 1 {
			msg.SendMessage(app.Bot, i18n.RuLocal["short_nickname"], nil)
		} else {
			userByNick, _ := app.GetUserByNick(newNickname)
			if reflect.DeepEqual(userByNick, User{}) {
				oldNickame := user.Name
				app.UpdateUserNickname(user.TgID, newNickname)
				app.UpdateUserLastMessage(user.TgID)
				message, _ := msg.SendMessage(app.Bot, fmt.Sprintf(i18n.RuLocal["new_nick"], newNickname), nil)
				app.BroadcastPlainMessage(
					fmt.Sprintf(i18n.RuLocal["new_user_nick"], oldNickame, newNickname),
					msg.From.ID,
					"",
					msg.From.ID,
					message.MessageID)
			} else {
				msg.SendMessage(app.Bot, i18n.RuLocal["nickname_exists"], nil)
			}
		}
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}
