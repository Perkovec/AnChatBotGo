package tgapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func callMethod(methodName string, token string, data interface{}) ([]byte, error) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methodName)

	var byteData bytes.Buffer
	if err := json.NewEncoder(&byteData).Encode(data); err != nil {
		return []byte{}, err
	}

	res, err := http.Post(apiURL, "application/json", &byteData)
	if err != nil {
		return []byte{}, err
	}

	res.Close = true
	defer res.Body.Close()

	json, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func getMe(token string) (User, error) {
	me, err := callMethod("getMe", token, nil)
	if err != nil {
		return User{}, err
	}

	var res struct {
		Ok          bool
		Result      User
		Description string
	}

	err = json.Unmarshal(me, &res)
	if err != nil {
		return User{}, fmt.Errorf("Invalid token (tgapi/getMe)")
	}

	if res.Ok {
		return res.Result, nil
	}

	return User{}, fmt.Errorf("Error (tgapi/getMe): %s", res.Description)
}

// GetUpdates implementation https://core.telegram.org/bots/api#getupdates
func (bot *Bot) GetUpdates() ([]Update, error) {
	reqParams := map[string]string{
		"offset": strconv.Itoa(bot.Offset),
	}

	uJSON, err := callMethod("getUpdates", bot.Token, reqParams)
	if err != nil {
		return []Update{}, err
	}

	var res struct {
		Ok          bool
		Result      []Update
		Description string
	}

	err = json.Unmarshal(uJSON, &res)
	if err != nil {
		return []Update{}, err
	}

	if !res.Ok {
		return []Update{}, fmt.Errorf("Error (tgapi/getUpdates): %s", res.Description)
	}

	return res.Result, nil
}

// SendMessage implementation https://core.telegram.org/bots/api#sendmessage
func (bot *Bot) SendMessage(chatID int, text string, options *SendMessageOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"text":    text,
	}

	if options != nil {
		passSendMessageOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendMessage", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendMessage error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendMessage error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendMessage error (tglib - tg): %s", res.Description)
	}

	return res.Result, nil
}

func (bot *Bot) SendAudio(chatID int, audio string, options *SendAudioOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"audio":   audio,
	}

	if options != nil {
		passSendAudioOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendAudio", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendAudio error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendAudio error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendAudio error (tglib - tg): %s", res.Description)
	}

	return res.Result, nil
}

func (bot *Bot) SendPhoto(chatID int, photo string, options *SendPhotoOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"photo":   photo,
	}

	if options != nil {
		passSendPhotoOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendPhoto", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendPhoto error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendPhoto error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendPhoto error (tglib - tg): %s", err)
	}

	return res.Result, nil
}
func (bot *Bot) SendDocument(chatID int, document string, options *SendDocumentOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id":  strconv.Itoa(chatID),
		"document": document,
	}

	if options != nil {
		passSendDocumentOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendDocument", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendDocument error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendDocument error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendDocument error (tglib - tg): %s", err)
	}

	return res.Result, nil
}

func (bot *Bot) SendSticker(chatID int, sticker string, options *SendStickerOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"sticker": sticker,
	}

	if options != nil {
		passSendStickerOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendSticker", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendSticker error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendSticker error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendSticker error (tglib - tg): %s", err)
	}

	return res.Result, nil
}

func (bot *Bot) SendVideo(chatID int, video string, options *SendVideoOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"video":   video,
	}

	if options != nil {
		passSendVideoOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendVideo", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendVideo error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendVideo error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendVideo error (tglib - tg): %s", err)
	}

	return res.Result, nil
}

func (bot *Bot) SendVoice(chatID int, voice string, options *SendVoiceOpts) (Message, error) {
	reqParams := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"voice":   voice,
	}

	if options != nil {
		passSendVoiceOpts(reqParams, options)
	}

	mJSON, err := callMethod("sendVoice", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("SendVoice error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("SendVoice error (tglib - parse): %s", err)
	}

	if !res.Ok {
		bot.OnError(Message{
			From: User{
				ID: chatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("SendVoice error (tglib - tg): %s", err)
	}

	return res.Result, nil
}

func (bot *Bot) EditMessageCaption(options *EditMessageCaptionOpts) (Message, error) {
	reqParams := map[string]string{}

	if options != nil {
		passEditMessageCaptionOpts(reqParams, options)
	}

	mJSON, err := callMethod("editMessageCaption", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("EditMessageCaption error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("EditMessageCaption error (tglib - parse): %s", err)
	}

	if !res.Ok {
		ichatID, _ := strconv.Atoi(options.ChatID)
		bot.OnError(Message{
			From: User{
				ID: ichatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("EditMessageCaption error (tglib - tg): %s", err)
	}

	return res.Result, nil
}

func (bot *Bot) EditMessageText(text string, options *EditMessageTextOpts) (Message, error) {
	reqParams := map[string]string{
		"text": text,
	}

	if options != nil {
		passEditMessageTextOpts(reqParams, options)
	}

	mJSON, err := callMethod("editMessageText", bot.Token, reqParams)
	if err != nil {
		return Message{}, fmt.Errorf("EditMessageText error (tglib - call): %s", err)
	}

	var res struct {
		Ok          bool
		Result      Message
		Description string
	}

	err = json.Unmarshal(mJSON, &res)
	if err != nil {
		return Message{}, fmt.Errorf("EditMessageText error (tglib - parse): %s", err)
	}

	if !res.Ok {
		ichatID, _ := strconv.Atoi(options.ChatID)
		bot.OnError(Message{
			From: User{
				ID: ichatID,
			},
		}, res.Description)
		return Message{}, fmt.Errorf("EditMessageText error (tglib - tg): %s", err)
	}

	return res.Result, nil
}
