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

// User represents a Telegram user or bot.
type User struct {
	// ID Unique identifier for this user or bot
	ID int `json:"id"`
	// FirstName User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// LastName Optional. User‘s or bot’s last name
	LastName string `json:"last_name"`
	// Username Optional. User‘s or bot’s username
	Username string `json:"username"`
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
