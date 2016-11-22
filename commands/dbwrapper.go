package commands

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (app *App) GetUserByTgID(userID int) (User, error) {
	user := User{}
	err := app.UsersC.Find(bson.M{"tg_id": userID}).One(&user)
	return user, err
}

func (app *App) GetChatUsers() ([]User, error) {
	users := []User{}
	err := app.UsersC.Find(bson.M{"is_chat_user": true, "banned": false}).All(&users)
	return users, err
}

func (app *App) GetRecentChatUsers() ([]User, error) {
	users := []User{}
	err := app.UsersC.Find(bson.M{"is_chat_user": true, "banned": false}).Sort("-last_message").All(&users)
	return users, err
}

func (app *App) GetRepliesByID(replyid int) (Replies, error) {
	replies := Replies{}
	err := app.MessagesC.Find(bson.M{
		"replies": bson.M{
			"$elemMatch": bson.M{
				"msg_id": replyid,
			},
		},
	}).One(&replies)
	return replies, err
}

func (app *App) UpdateUserLastMessage(userid int) error {
	return app.UsersC.Update(bson.M{"tg_id": userid}, bson.M{
		"$set": bson.M{
			"last_message": time.Now().Unix(),
		},
	})
}

func (app *App) GetUserByNick(nick string) (User, error) {
	user := User{}
	err := app.UsersC.Find(bson.M{"name": nick}).One(&user)
	return user, err
}

func (app *App) GetUserByChatID(id string) (User, error) {
	user := User{}
	err := app.UsersC.Find(bson.M{"chat_id": id}).One(&user)
	return user, err
}

func (app *App) UpdateUserNickname(userid int, name string) error {
	return app.UsersC.Update(bson.M{"tg_id": userid}, bson.M{
		"$set": bson.M{
			"name": name,
		},
	})
}
