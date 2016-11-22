package commands

import (
	"fmt"
	"html"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

type MSGData struct {
	msg      tgapi.Message
	user     User
	template string
	cb       func(User, string, int) (int, int)
}

func (app *App) BroadcastMessage(msg tgapi.Message) {
	user, err := app.GetUserByTgID(msg.From.ID)
	if err != nil {
		fmt.Println("Error (BroadcastMessage/GetUserByTgID): ", err)
	}
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if !reflect.DeepEqual(user, User{}) && user.IsChatUser {
		if len(msg.Text) > 0 {
			app.sendText(msg, user)
		} else if !reflect.DeepEqual(msg.Audio, tgapi.Audio{}) {
			app.sendAudio(msg, user)
		} else if len(msg.Photo) > 0 {
			app.sendPhoto(msg, user)
		} else if !reflect.DeepEqual(msg.Document, tgapi.Document{}) {
			app.sendDocument(msg, user)
		} else if !reflect.DeepEqual(msg.Sticker, tgapi.Sticker{}) {
			app.sendSticker(msg, user)
		} else if !reflect.DeepEqual(msg.Video, tgapi.Video{}) {
			app.sendVideo(msg, user)
		} else if !reflect.DeepEqual(msg.Voice, tgapi.Voice{}) {
			app.sendVoice(msg, user)
		}
	} else {
		msg.SendMessage(app.Bot, i18n.RuLocal["not_in_chat"], nil)
	}
}

func findMsgID(replies Replies, userid int) int {
	for _, v := range replies.Replies {
		if v.TgID == userid {
			return v.MsgID
		}
	}
	return 0
}

func (app *App) sendEach(msgdata MSGData) {
	documentReplies := []bson.M{}
	nickname := msgdata.user.Name
	text := getTextToSend(msgdata.msg, nickname, msgdata.template)
	users, err := app.GetChatUsers()
	if err != nil {
		fmt.Println("Error (BroadcastMessage/sendEach): ", err)
	}
	if msgdata.msg.ReplyToMessage != nil {
		replies, err := app.GetRepliesByID(msgdata.msg.ReplyToMessage.MessageID)
		if err != nil {
			fmt.Println("Error (BroadcastMessage/GetRepliesByID): ", err)
		}
		for _, user := range users {
			if user.TgID != msgdata.msg.From.ID {
				var reply int
				if !reflect.DeepEqual(replies, Replies{}) {
					reply = findMsgID(replies, user.TgID)
				}
				receiverid, msgid := msgdata.cb(user, text, reply)
				documentReplies = append(documentReplies, bson.M{
					"tg_id":  receiverid,
					"msg_id": msgid,
				})
			}
		}
		documentReplies = append(documentReplies, bson.M{
			"tg_id":  msgdata.msg.From.ID,
			"msg_id": msgdata.msg.MessageID,
		})
	} else {
		for _, user := range users {
			if user.TgID != msgdata.msg.From.ID {
				receiverid, msgid := msgdata.cb(user, text, 0)
				documentReplies = append(documentReplies, bson.M{
					"tg_id":  receiverid,
					"msg_id": msgid,
				})
			}
		}
		documentReplies = append(documentReplies, bson.M{
			"tg_id":  msgdata.msg.From.ID,
			"msg_id": msgdata.msg.MessageID,
		})
	}
	err = app.MessagesC.Insert(bson.M{
		"replies": documentReplies,
	})

	app.UpdateUserLastMessage(msgdata.msg.From.ID)

	if err != nil {
		fmt.Println("Error (BroadcastMessage/Insert): ", err)
	}
}

func (app *App) sendVoice(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: i18n.RuLocal["voice_from_user"],
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendVoice(receiver.TgID, msg.Voice.FileID, &tgapi.SendVoiceOpts{
				Caption:          text,
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendVoice): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendVideo(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: i18n.RuLocal["video_from_user"],
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendVideo(receiver.TgID, msg.Video.FileID, &tgapi.SendVideoOpts{
				Caption:          text,
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendVideo): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendSticker(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: "",
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendSticker(receiver.TgID, msg.Sticker.FileID, &tgapi.SendStickerOpts{
				ReplyToMessageID: replyID,
				ReplyMarkup: tgapi.ReplyMarkup{
					InlineKeyboard: [][]tgapi.InlineKeyboardButton{
						[]tgapi.InlineKeyboardButton{
							tgapi.InlineKeyboardButton{
								Text:         receiver.Name,
								CallbackData: "sticker_button",
							},
						},
					},
				},
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendSticker): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendDocument(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: i18n.RuLocal["document_from_user"],
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendDocument(receiver.TgID, msg.Document.FileID, &tgapi.SendDocumentOpts{
				Caption:          text,
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendDocument): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendPhoto(msg tgapi.Message, user User) {
	photoid := msg.Photo[len(msg.Photo)-1].FileID
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: i18n.RuLocal["photo_from_user"],
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendPhoto(receiver.TgID, photoid, &tgapi.SendPhotoOpts{
				Caption:          text,
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendPhoto): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendAudio(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:      msg,
		user:     user,
		template: i18n.RuLocal["audio_from_user"],
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendAudio(receiver.TgID, msg.Audio.FileID, &tgapi.SendAudioOpts{
				Caption:          text,
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendAudio): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func (app *App) sendText(msg tgapi.Message, user User) {
	app.sendEach(MSGData{
		msg:  msg,
		user: user,
		cb: func(receiver User, text string, replyID int) (int, int) {
			message, err := app.Bot.SendMessage(receiver.TgID, text, &tgapi.SendMessageOpts{
				ParseMode:        "HTML",
				ReplyToMessageID: replyID,
			})
			if err != nil {
				fmt.Println("Error (BroadcastMessage/sendText): ", err)
			}
			return receiver.TgID, message.MessageID
		},
	})
}

func getTextToSend(msg tgapi.Message, nickname string, template string) string {
	var text string
	if len(msg.Caption) > 0 || (len(template) == 0 && len(msg.Text) > 0) {
		if len(msg.Caption) > 0 {
			text = msg.Caption
		} else {
			text = msg.Text
		}
		text = nickname + ": " + text
	} else {
		text = html.EscapeString(fmt.Sprintf(template, nickname))
	}

	return text
}
