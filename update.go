package TeleGogo

// Update represents a state update in Telegram that has relevance to a bot client.
type Update struct {
	// ID The update‘s unique identifier
	ID int `json:"update_id"`
	// Message Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message Message `json:"message"`
	// EditedMessage Optional. New version of a message that is known to the bot and was edited
	EditedMessage Message `json:"edited_message"`
	// Inline Optional. New incoming inline query
	Inline InlineQuery `json:"inline_query"`
	// InlineResult Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	InlineResult ChosenInlineResult `json:"chosen_inline_result"`
	// Callback Optional. New incoming callback query
	Callback CallbackQuery `json:"callback_query"`
}

type Message struct {
	Id                    int             `json:"message_id"`
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
	Photo                 []PhotoSize     `json:"photo"`
	Sticker               Sticker         `json:"sticker"`
	Video                 Video           `json:"video"`
	Voice                 Voice           `json:"voice"`
	Caption               string          `json:"caption"`
	Contact               Contact         `json:"contact"`
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
	PinnedMessage         *Message        `json:"pinned_message"`
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type CallbackQuery struct {
}

type User struct {
}

type Chat struct {
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   User   `json:"user"`
}

type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

type Document struct {
	FileID    string    `json:"file_id"`
	Thumbnail PhotoSize `json:"thumb"`
	Filename  string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type Sticker struct {
	FileID    string    `json:"file_id"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail PhotoSize `json:"thumb"`
	Emoji     string    `json:"emoji"`
	FileSize  int       `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
}

type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FourSquareID string   `json:"foursquare_id"`
}

type Video struct {
	FileID    string    `json:"file_id"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Duration  int       `json:"duration"`
	Thumbnail PhotoSize `json:"thumb"`
	MimeType  string    `json:"mime_type"`
	FileSize  int       `json:"file_size"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
