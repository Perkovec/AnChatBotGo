package tgapi

import "log"

// Bot instance
type Bot struct {
	Token   string
	Offset  int
	Info    User
	OnError func(Message, string)
}

// Create bot
func Create(token string) (*Bot, error) {
	bot, err := getMe(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Token:  token,
		Offset: 0,
		Info:   bot,
	}, nil
}

// Polling looks for updates
func (bot *Bot) Polling(messages chan Message, edites chan Message) {
	for {
		updates, err := bot.GetUpdates()
		if err != nil {
			log.Println("Polling error (tgapi/Polling):", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				messages <- *update.Message
			} else if update.EditedMessage != nil {
				edites <- *update.EditedMessage
			}

			bot.Offset = update.UpdateID + 1
		}
	}
}
