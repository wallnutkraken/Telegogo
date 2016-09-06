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

// Message represents a message.
type Message struct {
	// ID Unique message identifier
	ID int `json:"message_id"`
	// From Optional. Sender, can be empty for messages sent to channels
	From User `json:"from"`
	// Date Date the message was sent in Unix time
	Date int `json:"date"`
	// Chat Conversation the message belongs to
	Chat Chat `json:"chat"`
	// ForwardFrom Optional. For forwarded messages, sender of the original message
	ForwardFrom User `json:"forward_from"`
	// ForwardFromChat Optional. For messages forwarded from a channel, information about the original channel
	ForwardFromChat Chat `json:"forward_from_chat"`
	// ForwardDate Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date"`
	// ReplyToMessage Optional. For replies, the original message.
	ReplyToMessage *Message `json:"reply_to_message"`
	// EditDate Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date"`
	// Text Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text"`
	// Entities Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`
	// Audio Optional. Message is an audio file, information about the file
	Audio Audio `json:"audio"`
	// Document Optional. Message is a general file, information about the file
	Document Document `json:"document"`
	// Photo Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo"`
	// Sticker Optional. Message is a sticker, information about the sticker
	Sticker Sticker `json:"sticker"`
	// Video Optional. Message is a video, information about the video
	Video Video `json:"video"`
	// Voice Optional. Message is a voice message, information about the file
	Voice Voice `json:"voice"`
	// Caption Optional. Caption for the document, photo or video, 0-200 characters
	Caption string `json:"caption"`
	// Contact Optional. Message is a shared contact, information about the contact
	Contact Contact `json:"contact"`
	// Location Optional. Message is a shared location, information about the location
	Location Location `json:"location"`
	// Venue Optional. Message is a venue, information about the venue
	Venue Venue `json:"venue"`
	// NewChatMember Optional. A new member was added to the group, information about them
	NewChatMember User `json:"new_chat_member"`
	// LeftChatMember Optional. A member was removed from the group, information about them
	LeftChatMember User `json:"left_chat_member"`
	// NewChatTitle Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`
	// NewChatPhoto Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`
	// DeleteChatPhoto Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// GroupChatCreated Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`
	// SupergroupChatCreated Optional. Service message: the supergroup has been created.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	// ChannelChatCreated Optional. Service message: the channel has been created.
	ChannelChatCreated bool `json:"channel_chat_created"`
	// MigrateToChatID Optional. The group has been migrated to a supergroup with the specified identifier.
	MigrateToChatID int `json:"migrate_to_chat_id"`
	// MigrateFromChatID Optional. The supergroup has been migrated from a group with the specified identifier.
	MigrateFromChatID int `json:"migrate_from_chat_id"`
	// PinnedMessage Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message"`
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
