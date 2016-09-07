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

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Can be mention (@username), hashtag, bot_command, url, email,
	// bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block),
	// text_link (for clickable text URLs), text_mention (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// URL Optional. For “text_link” only, url that will be opened after user taps on the text
	URL string `json:"url"`
	// User Optional. For “text_mention” only, the mentioned user
	User User `json:"user"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Performer Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer"`
	// Title Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title"`
	// MimeType Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	// Thumbnail Optional. Document thumbnail as defined by sender
	Thumbnail PhotoSize `json:"thumb"`
	// Filename Optional. Original filename as defined by sender
	Filename string `json:"file_name"`
	// MimeType Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// Sticker represents a sticker.
type Sticker struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	// Filename Optional. Original filename as defined by sender
	Thumbnail PhotoSize `json:"thumb"`
	// Emoji Optional. Emoji associated with the sticker
	Emoji string `json:"emoji"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// UserID Optional. Contact's user identifier in Telegram
	UserID int `json:"user_id"`
}

// Voice represents a voice note.
type Voice struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// MimeType Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// Venue represents a venue
type Venue struct {
	// Location Venue location
	Location Location `json:"location"`
	// Title Name of the venue
	Title string `json:"title"`
	// Address Address of the venue
	Address string `json:"address"`
	// FourSquareID Optional. Foursquare identifier of the venue
	FourSquareID string `json:"foursquare_id"`
}

// Video represents a video file.
type Video struct {
	// FileID Unique identifier for this file
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumbnail Optional. Video thumbnail
	Thumbnail PhotoSize `json:"thumb"`
	// MimeType Optional. Mime type of a file as defined by sender
	MimeType string `json:"mime_type"`
	// FileSize Optional. File size
	FileSize int `json:"file_size"`
}

// Location represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
}
