package TeleGogo

const (
	Typing         = "typing"
	UploadPhoto    = "upload_photo"
	RecordVideo    = "record_video"
	RecordAudio    = "record_audio"
	UploadVideo    = "upload_video"
	UploadAudio    = "upload_audio"
	UploadDocument = "upload_document"
)

// SendChatActionArgs represents the optional and required arguments to the SendChatAction method
type SendChatActionArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Action. Required. Type of action to broadcast.
	// Choose one, depending on what the user is about to receive:
	// typing for text messages, upload_photo for photos, record_video or upload_video for videos,
	// record_audio or upload_audio for audio files, upload_document for general files,
	// find_location for location data.
	Action string `json:"action"`
}

func (a SendChatActionArgs) methodName() string {
	return "sendChatAction"
}
