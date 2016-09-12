package TeleGogo

// SendPhotoArgs represents the optional and required arguments for the SendPhoto method
type SendPhotoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// PhotoID Required. Path to photo file on device. In one request only this, or PhotoID should be set.
	// Use this if you're uploading a new file. Use PhotoID if you want to send a photo that is already on
	// Telegram's servers, and you know it's file ID.
	PhotoPath string `json:"-" tele:"!photo"`
	// FileID Required. Unique ID for a photo file on Telegram's servers. In one request only this,
	// or PhotoPath should be set. Use this if you want to send a photo that is already on Telegram's
	// servers and you know it's file ID. Otherwise, use PhotoPath.
	FileID string `json:"photo" tele:"-"`
	// Caption Optional. Photo caption. 0-200 chars.
	Caption string `json:"caption,omitempty" tele:"caption"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty" tele:"reply_markup"`
}

func (a SendPhotoArgs) methodName() string {
	return "sendPhoto"
}

// SendAudioArgs represents the optional and required arguments for the SendAudio method
type SendAudioArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// AudioPath Required. Path to audio file on device IF uploading a new file. If you are trying to
	// send an audio file which is already on Telegram's servers, please keep this field empty and
	// use the AudioFileID field.
	AudioPath string `json:"-" tele:"!audio"`
	// AudioFileID Required. File ID of the audio file you want to send. If you are trying to
	// send a new audio file which is not already on Telegram's servers that you know of, please keep
	// this field empty and use the AudioPath field instead
	AudioFileID string `json:"audio" tele:"-"`
	// Duration Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty" tele:"duration"`
	// Performer Optional.
	Performer string `json:"performer,omitempty" tele:"performer"`
	// Title Optional. Track name
	Title string `json:"title,omitempty" tele:"title"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty" tele:"reply_markup"`
}

func (a SendAudioArgs) methodName() string {
	return "sendAudio"
}

// SendDocumentArgs represents the optional and required arguments for the SendDocument method
type SendDocumentArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// DocumentPath Required. Use if you want to upload a new file to Telegram. If you are looking to send
	// a document that already exists on Telegram's servers and you know it's file ID, use DocumentFileID
	// instead. Only one of these fields may be filled.
	DocumentPath string `json:"-" tele:"!document"`
	// DocumentFileID Required. Use if you want to send a document that already exists on Telegram's servers
	// and you know it's file ID. If you want to upload a new file, use DocumentPath instead. Only one of
	// these fields may be filled.
	DocumentFileID string `json:"document" tele:"-"`
	// Caption Optional. Document caption. 0-200 characters
	Caption string `json:"caption,omitempty" tele:"caption"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup" tele:"reply_markup"`
}

func (a SendDocumentArgs) methodName() string {
	return "sendDocument"
}

// SendStickerArgs represents the optional and required arguments for the SendSticker method
type SendStickerArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// StickerPath Required. Use if uploading a sticker file (.webp). Represents the path to a sticker file.
	// Only use if sending a new sticker. If sending an existing sticker, please use StickerPath.
	StickerPath string `json:"-" tele:"!sticker"`
	// StickerID Required. Use if sending a sticker file that already exists. Represents the FileID of
	// the sticker. If you want to send a new sticker from a file, please use StickerPath.
	StickerID string `json:"sticker" tele:"-"`
	//DisableNotification Optional. Sends a message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkupAdditional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty" tele:"reply_markup"`
}

func (a SendStickerArgs) methodName() string {
	return "sendSticker"
}

// SendVideoArgs represents the optional and required arguments for the SendVideo method
type SendVideoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// VideoPath Required. File path of the video you want to send. If you want to send a video that already
	// exists on Telegram's servers and you know it's file ID, use VideoID instead.
	VideoPath string `json:"-" tele:"!video"`
	// VideoID Required. File ID of the video you want to send. TO upload a new video file, please use
	// VideoPath instead.
	VideoID string `json:"video" tele:"-"`
	// Duration Optional. Duration of sent video in seconds
	Duration int `json:"duration, omitempty" tele:"duration"`
	// Width Optional. Video width
	Width int `json:"width,omitempty" tele:"width"`
	// Height Optional. Video height
	Height int `json:"height,omitempty" tele:"height"`
	// Caption Optional. Video caption, 0-200 characters
	Caption string `json:"caption" tele:"caption"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkup Optional. Additional interface options.
	ReplyMarkup string `json:"reply_markup,omitempty" tele:"reply_markup"`
}

func (a SendVideoArgs) methodName() string {
	return "sendVideo"
}

// SendVoiceArgs represents the optional and required arguments for the SendVoice method
type SendVoiceArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id" tele:"chat_id"`
	// VoicePath Required. File path of the voice you want to send. If you want to send a voice that already
	// exists on Telegram's servers and you know it's file ID, use VoiceID instead.
	VoicePath string `json:"-" tele:"!voice"`
	// VoiceID Required. File ID of the voice you want to send. To upload a new voice file, please use
	// VoicePath instead.
	VoiceID string `json:"voice" tele:"-"`
	// Duration Optional. Duration of sent audio in seconds
	Duration int `json:"duration" tele:"duration"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty" tele:"disable_notification"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty" tele:"reply_to_message_id"`
	// ReplyMarkup Optional. Additional interface options.
	ReplyMarkup string `json:"reply_markup,omitempty" tele:"reply_markup"`
}

func (a SendVoiceArgs) methodName() string {
	return "sendVoice"
}

type apiCaller interface {
	methodName() string
}
