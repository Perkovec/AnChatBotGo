package commands

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"github.com/perkovec/AnChatBotGo/i18n"
	"github.com/perkovec/AnChatBotGo/tgapi"
)

func (app *App) StopCommand(msg tgapi.Message) {
	user, err := app.GetUserByTgID(msg.From.ID)
	if err != nil {
		fmt.Println("Error (StopCommand/GetUserByTgID): ", err)
	}
	if user.Banned {
		msg.SendMessage(app.Bot, i18n.RuLocal["you_are_banned"], nil)
	} else if !reflect.DeepEqual(user, User{}) {
		app.UsersC.Update(bson.M{"tg_id": user.TgID}, bson.M{
			"$set": bson.M{
				"is_chat_user": false,
			},
		})
		message, err := msg.SendMessage(app.Bot, i18n.RuLocal["stop"], nil)
		if err != nil {
			fmt.Println("Error (StopCommand/SendMessage): ", err)
		}
		app.BroadcastPlainMessage(fmt.Sprintf(i18n.RuLocal["leave_chat"], user.Name), msg.From.ID, "", msg.From.ID, message.MessageID)
	}
}
