package commands

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) InfoCommand(msg tgapi.Message, userid string) {
	user, _ := app.GetUserByTgID(msg.From.ID)
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if user.IsChatUser {
		userid = strings.TrimSpace(strings.ToUpper(userid))
		if len(userid) > 0 {
			chatuser, _ := app.GetUserByChatID(userid)
			if !reflect.DeepEqual(chatuser, User{}) {
				app.makeInfo(msg, chatuser)
			} else {
				msg.SendMessage(app.Bot, i18n.RuLocal["user_not_found"], nil)
			}
		} else {
			app.makeInfo(msg, user)
		}
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}

func (app *App) makeInfo(msg tgapi.Message, user User) {
	lastMsgDiff := int(time.Now().Unix()) - user.LastMessage
	since := TimeDiff(int(time.Now().Unix()) - user.StartDate)
	lastOnline := TimeDiff(lastMsgDiff)

	var template string
	var useOnline bool
	if lastMsgDiff < 3*60 {
		template = i18n.RuLocal["info_online"]
	} else {
		template = i18n.RuLocal["info"]
		useOnline = true
	}

	if useOnline {
		msg.SendMessage(app.Bot, fmt.Sprintf(template, user.Name, since, lastOnline), nil)
	} else {
		msg.SendMessage(app.Bot, fmt.Sprintf(template, user.Name, since), nil)
	}

	app.UpdateUserLastMessage(msg.From.ID)
}
