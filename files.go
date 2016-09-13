package TeleGogo

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

// File represents a file ready to be downloaded.
type File struct {
	//ID Unique identifier for this file
	ID string `json:"file_id"`
	// FileSize Optional. File size, if known
	FileSize int `json:"file_size"`
	// FilePath Optional. File path.
	FilePath string `json:"file_path"`
}

type getFileArgs struct {
	FileID string `json:"file_id"`
}

type fileResponse struct {
	OK     bool `json:"ok"`
	Result File `json:"result"`
}

func (a getFileArgs) methodName() string {
	return "getFile"
}
