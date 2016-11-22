package tgapi

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func passSendMessageOpts(params map[string]string, options *SendMessageOpts) {
	if len(options.ParseMode) > 0 {
		params["parse_mode"] = options.ParseMode
	}

	if options.DisableWebPagePreview {
		params["disable_web_page_preview"] = "true"
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID != 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendAudioOpts(params map[string]string, options *SendAudioOpts) {
	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if options.Duration > 0 {
		params["duration"] = strconv.Itoa(options.Duration)
	}

	if len(options.Performer) > 0 {
		params["performer"] = options.Performer
	}

	if len(options.Title) > 0 {
		params["title"] = options.Title
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendPhotoOpts(params map[string]string, options *SendPhotoOpts) {
	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendDocumentOpts(params map[string]string, options *SendDocumentOpts) {
	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendStickerOpts(params map[string]string, options *SendStickerOpts) {
	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendVideoOpts(params map[string]string, options *SendVideoOpts) {
	if options.Duration > 0 {
		params["duration"] = strconv.Itoa(options.Duration)
	}

	if options.Width > 0 {
		params["width"] = strconv.Itoa(options.Width)
	}

	if options.Height > 0 {
		params["height"] = strconv.Itoa(options.Height)
	}

	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passSendVoiceOpts(params map[string]string, options *SendVoiceOpts) {
	if options.Duration > 0 {
		params["duration"] = strconv.Itoa(options.Duration)
	}

	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if options.DisableNotification {
		params["disable_notification"] = "true"
	}

	if options.ReplyToMessageID > 0 {
		params["reply_to_message_id"] = strconv.Itoa(options.ReplyToMessageID)
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passEditMessageCaptionOpts(params map[string]string, options *EditMessageCaptionOpts) {
	if len(options.ChatID) > 0 {
		params["chat_id"] = options.ChatID
	}

	if options.MessageID > 0 {
		params["message_id"] = strconv.Itoa(options.MessageID)
	}

	if len(options.InlineMessageID) > 0 {
		params["inline_message_id"] = options.InlineMessageID
	}

	if len(options.Caption) > 0 {
		params["caption"] = options.Caption
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}

func passEditMessageTextOpts(params map[string]string, options *EditMessageTextOpts) {
	if len(options.ChatID) > 0 {
		params["chat_id"] = options.ChatID
	}

	if options.MessageID > 0 {
		params["message_id"] = strconv.Itoa(options.MessageID)
	}

	if len(options.InlineMessageID) > 0 {
		params["inline_message_id"] = options.InlineMessageID
	}

	if len(options.ParseMode) > 0 {
		params["parse_mode"] = options.ParseMode
	}

	if options.DisableWebPagePreview {
		params["disable_web_page_preview"] = "true"
	}

	if !reflect.DeepEqual(options.ReplyMarkup, ReplyMarkup{}) {
		replyMarkup, _ := json.Marshal(options.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}
}
