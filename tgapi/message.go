package tgapi

func (msg *Message) SendMessage(bot *Bot, text string, options *SendMessageOpts) (Message, error) {
	return bot.SendMessage(msg.Chat.ID, text, options)
}
