package main

import (
	"html"
	"reflect"
	"strconv"
	"strings"

	"github.com/perkovec/AnChatBotGo/commands"
	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func findMsgID(replies commands.Replies, userid int) int {
	for _, v := range replies.Replies {
		if v.TgID == userid {
			return v.MsgID
		}
	}
	return 0
}

func ProcessUpdate(app commands.App, msg tgapi.Message) {
	text := strings.TrimSpace(msg.Text)

	if len(text) > 0 {
		if CRegex["me"].MatchString(text) {
			matches := CRegex["me"].FindAllStringSubmatch(text, -1)
			newText := app.GetMeText(msg, strings.TrimSpace(matches[0][2]))
			if len(newText) > 0 {
				ChangeMessages(app, msg, newText, msg.From.ID, false)
			}
		} else if CRegex["me2"].MatchString(text) {
			matches := CRegex["me2"].FindAllStringSubmatch(text, -1)
			newText := app.GetMeText(msg, strings.TrimSpace(matches[0][2]))
			if len(newText) > 0 {
				ChangeMessages(app, msg, newText, msg.From.ID, false)
			}
		} else if CRegex["some_command"].MatchString(text) {
			msg.SendMessage(app.Bot, i18n.RuLocal["disable_command_edit"], nil)
		} else {
			user, _ := app.GetUserByTgID(msg.From.ID)
			msgText := user.Name + ": " + text
			ChangeMessages(app, msg, html.EscapeString(msgText), msg.From.ID, false)
		}
	} else {
		user, _ := app.GetUserByTgID(msg.From.ID)
		msgText := user.Name + ": " + text
		ChangeMessages(app, msg, html.EscapeString(msgText), msg.From.ID, true)
	}
}

func ChangeMessages(app commands.App, msg tgapi.Message, newText string, exclude int, caption bool) {
	replies, _ := app.GetRepliesByID(msg.MessageID)
	if !reflect.DeepEqual(replies, commands.Replies{}) {
		users, _ := app.GetRecentChatUsers()
		for _, user := range users {
			if user.TgID != exclude {
				reply := findMsgID(replies, user.TgID)
				if caption {
					app.Bot.EditMessageCaption(&tgapi.EditMessageCaptionOpts{
						ChatID:    strconv.Itoa(user.TgID),
						MessageID: reply,
						Caption:   newText,
					})
				} else {
					app.Bot.EditMessageText(newText, &tgapi.EditMessageTextOpts{
						ChatID:    strconv.Itoa(user.TgID),
						MessageID: reply,
						ParseMode: "HTML",
					})
				}
			}
		}
	}
}
