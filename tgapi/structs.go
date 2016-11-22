package tgapi

// User instance https://core.telegram.org/bots/api#user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Chat instance https://core.telegram.org/bots/api#chat
type Chat struct {
	ID                          int    `json:"id"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	Username                    string `json:"username"`
	FirstName                   string `json:"first_name"`
	LastName                    string `json:"last_name"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

// Message instance https://core.telegram.org/bots/api#message
type Message struct {
	MessageID             int             `json:"message_id"`
	From                  User            `json:"from"`
	Date                  int             `json:"date"`
	Chat                  Chat            `json:"chat"`
	ForwardFrom           User            `json:"forward_from"`
	ForwardFromChat       Chat            `json:"forward_from_chat"`
	ForwardDate           int             `json:"forward_date"`
	ReplyToMessage        *Message        `json:"reply_to_message"`
	EditDate              int             `json:"edit_date"`
	Text                  string          `json:"text"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 Audio           `json:"audio"`
	Document              Document        `json:"document"`
	Game                  Game            `json:"game"`
	Photo                 []PhotoSize     `json:"photo"`
	Sticker               Sticker         `json:"sticker"`
	Video                 Video           `json:"video"`
	Voice                 Voice           `json:"voice"`
	Caption               string          `json:"caption"`
	Location              Location        `json:"location"`
	Venue                 Venue           `json:"venue"`
	NewChatMember         User            `json:"new_chat_member"`
	LeftChatMember        User            `json:"left_chat_member"`
	NewChatTitle          string          `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`
	GroupChatCreated      bool            `json:"group_chat_created"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created"`
	ChannelChatCreated    bool            `json:"channel_chat_created"`
	MigrateToChatID       int             `json:"migrate_to_chat_id"`
	MigrateFromChatID     int             `json:"migrate_from_chat_id"`
	PinnedMessage         *Message        `json:"pinned_message"`
}

// MessageEntity instance https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   User   `json:"user"`
}

// PhotoSize instance https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

// Audio instance https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

// Document instance https://core.telegram.org/bots/api#document
type Document struct {
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// Sticker instance https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Emoji    string    `json:"emoji"`
	FileSize int       `json:"file_size"`
}

// Video instance https://core.telegram.org/bots/api#video
type Video struct {
	FileID   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// Voice instance https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

// Location instance https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue instance https://core.telegram.org/bots/api#venue
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"`
}

// Game instance https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"test"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

// Animation instance https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// Update instance https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           int                `json:"update_id"`
	Message            *Message           `json:"message"`
	EditedMessage      *Message           `json:"edited_message"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery     `json:"callback_query"`
}

// InlineQuery instance https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
}

// ChosenInlineResult instance https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageID string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

// CallbackQuery instance https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
}

// SendMessageOpts - optional parameters for https://core.telegram.org/bots/api#sendmessage
type SendMessageOpts struct {
	ParseMode             string
	DisableWebPagePreview bool
	DisableNotification   bool
	ReplyToMessageID      int
	ReplyMarkup           ReplyMarkup
}

// ReplyMarkup instance
type ReplyMarkup struct {
	ForceReply         bool                     `json:"force_reply,omitempty"`
	Keyboard           [][]string               `json:"keyboard,omitempty"`
	ResizeKeyboard     bool                     `json:"resize_keyboard,omitempty"`
	InlineKeyboard     [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	OneTimeKeyboard    bool                     `json:"one_time_keyboard,omitempty"`
	HideCustomKeyboard bool                     `json:"hide_keyboard,omitempty"`
	Selective          bool                     `json:"selective,omitempty"`
}

// InlineKeyboardButton instance https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	URL                          string `json:"url,emitempty"`
	CallbackData                 string `json:"callback_data,emitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,emitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,emitempty"`
}

// EditMessageTextOpts - optional parameters for https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextOpts struct {
	ChatID                string
	MessageID             int
	InlineMessageID       string
	ParseMode             string
	DisableWebPagePreview bool
	ReplyMarkup           ReplyMarkup
}

// AnswerCallbackQueryOpts - optional parameters for https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQueryOpts struct {
	Text      string
	ShowAlert bool
	URL       string
}

// EditMessageReplyMarkupOpts - optional parameters for https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkupOpts struct {
	ChatID          string
	MessageID       int
	InlineMessageID string
	ReplyMarkup     ReplyMarkup
}

type SendAudioOpts struct {
	Caption             string
	Duration            int
	Performer           string
	Title               string
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type SendPhotoOpts struct {
	Caption             string
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type SendDocumentOpts struct {
	Caption             string
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type SendStickerOpts struct {
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type SendVideoOpts struct {
	Duration            int
	Width               int
	Height              int
	Caption             string
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type SendVoiceOpts struct {
	Caption             string
	Duration            int
	DisableNotification bool
	ReplyToMessageID    int
	ReplyMarkup         ReplyMarkup
}

type EditMessageCaptionOpts struct {
	ChatID          string
	MessageID       int
	InlineMessageID string
	Caption         string
	ReplyMarkup     ReplyMarkup
}
