package TeleGogo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// createInputFileBody reads a file and creates a multipart writer for POSTing with the file and
// the form name
func createInputFileBody(path string, formName string) (*multipart.Writer, *bytes.Buffer, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	fileUpload, err := writer.CreateFormFile(formName, path)
	if err != nil {
		return nil, nil, err
	}
	if _, err := io.Copy(fileUpload, file); err != nil {
		return nil, nil, err
	}
	return writer, &buffer, nil
}

// SendPhotoArgs represents the optional and required arguments for the SendPhoto method
type SendPhotoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// PhotoID Required. Path to photo file on device. In one request only this, or PhotoID should be set.
	// Use this if you're uploading a new file. Use PhotoID if you want to send a photo that is already on
	// Telegram's servers, and you know it's file ID.
	PhotoPath string `json:"-"`
	// FileID Required. Unique ID for a photo file on Telegram's servers. In one request only this,
	// or PhotoPath should be set. Use this if you want to send a photo that is already on Telegram's
	// servers and you know it's file ID. Otherwise, use PhotoPath.
	FileID string `json:"photo"`
	// Caption Optional. Photo caption. 0-200 chars.
	Caption string `json:"caption,omitempty"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendPhotoArgs) methodName() string {
	return "sendPhoto"
}

func (a SendPhotoArgs) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a SendPhotoArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.PhotoPath, "photo")
	if err != nil {
		return nil, nil, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", a.ChatID)

	/* Optional args */
	if a.Caption != "" {
		writer.WriteField("caption", a.Caption)
	}
	if a.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if a.ReplyMarkup != "" {
		writer.WriteField("reply_markup", a.ReplyMarkup)
	}
	if a.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(a.ReplyToMessageID))
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}
	return writer, buffer, nil
}

// SendAudioArgs represents the optional and required arguments for the SendAudio method
type SendAudioArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// AudioPath Required. Path to audio file on device IF uploading a new file. If you are trying to
	// send an audio file which is already on Telegram's servers, please keep this field empty and
	// use the AudioFileID field.
	AudioPath string `json:"-"`
	// AudioFileID Required. File ID of the audio file you want to send. If you are trying to
	// send a new audio file which is not already on Telegram's servers that you know of, please keep
	// this field empty and use the AudioPath field instead
	AudioFileID string `json:"audio"`
	// Duration Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Performer Optional.
	Performer string `json:"performer,omitempty"`
	// Title Optional. Track name
	Title string `json:"title,omitempty"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendAudioArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.AudioPath, "audio")
	if err != nil {
		return nil, nil, err
	}

	/* Add ChatID */
	writer.WriteField("chat_id", a.ChatID)

	/* Optional args */
	if a.Duration != 0 {
		writer.WriteField("duration", strconv.Itoa(a.Duration))
	}
	if a.Performer != "" {
		writer.WriteField("performer", a.Performer)
	}
	if a.Title != "" {
		writer.WriteField("title", a.Title)
	}
	if a.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if a.ReplyMarkup != "" {
		writer.WriteField("reply_markup", a.ReplyMarkup)
	}
	if a.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(a.ReplyToMessageID))
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}

	return writer, buffer, nil
}

func (a SendAudioArgs) methodName() string {
	return "sendAudio"
}

func (a SendAudioArgs) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

// SendDocumentArgs represents the optional and required arguments for the SendDocument method
type SendDocumentArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// DocumentPath Required. Use if you want to upload a new file to Telegram. If you are looking to send
	// a document that already exists on Telegram's servers and you know it's file ID, use DocumentFileID
	// instead. Only one of these fields may be filled.
	DocumentPath string `json:"-"`
	// DocumentFileID Required. Use if you want to send a document that already exists on Telegram's servers
	// and you know it's file ID. If you want to upload a new file, use DocumentPath instead. Only one of
	// these fields may be filled.
	DocumentFileID string `json:"document"`
	// Caption Optional. Document caption. 0-200 characters
	Caption string `json:"caption,omitempty"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup"`
}

func (a SendDocumentArgs) methodName() string {
	return "sendDocument"
}

func (a SendDocumentArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.DocumentPath, "document")
	if err != nil {
		return nil, nil, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", a.ChatID)

	/* Optional args */
	if a.Caption != "" {
		writer.WriteField("caption", a.Caption)
	}
	if a.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if a.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(a.ReplyToMessageID))
	}
	if a.ReplyMarkup != "" {
		writer.WriteField("reply_markup", a.ReplyMarkup)
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}

	return writer, buffer, nil
}

func (a SendDocumentArgs) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

// SendStickerArgs represents the optional and required arguments for the SendSticker method
type SendStickerArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// StickerPath Required. Use if uploading a sticker file (.webp). Represents the path to a sticker file.
	// Only use if sending a new sticker. If sending an existing sticker, please use StickerPath.
	StickerPath string `json:"-"`
	// StickerID Required. Use if sending a sticker file that already exists. Represents the FileID of
	// the sticker. If you want to send a new sticker from a file, please use StickerPath.
	StickerID string `json:"sticker"`
	//DisableNotification Optional. Sends a message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkupAdditional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendStickerArgs) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a SendStickerArgs) methodName() string {
	return "sendSticker"
}

func (a SendStickerArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.StickerPath, "sticker")
	if err != nil {
		return nil, nil, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", a.ChatID)

	/* Optional args */
	if a.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if a.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(a.ReplyToMessageID))
	}
	if a.ReplyMarkup != "" {
		writer.WriteField("reply_markup", a.ReplyMarkup)
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}

	return writer, buffer, nil
}

// SendVideoArgs represents the optional and required arguments for the SendVideo method
type SendVideoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// VideoPath Required. File path of the video you want to send. If you want to send a video that already
	// exists on Telegram's servers and you know it's file ID, use VideoID instead.
	VideoPath string `json:"-"`
	// VideoID Required. File ID of the video you want to send. TO upload a new video file, please use
	// VideoPath instead.
	VideoID string `json:"video"`
	// Duration Optional. Duration of sent video in seconds
	Duration int `json:"duration, omitempty"`
	// Width Optional. Video width
	Width int `json:"width,omitempty"`
	// Height Optional. Video height
	Height int `json:"height,omitempty"`
	// Caption Optional. Video caption, 0-200 characters
	Caption string `json:"caption"`
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendVideoArgs) toJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a SendVideoArgs) methodName() string {
	return "sendVideo"
}

func (a SendVideoArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.VideoPath, "video")
	if err != nil {
		return nil, nil, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", a.ChatID)

	/* Optional args */
	if a.Duration != 0 {
		writer.WriteField("duration", strconv.Itoa(a.Duration))
	}
	if a.Width != 0 {
		writer.WriteField("width", strconv.Itoa(a.Width))
	}
	if a.Height != 0 {
		writer.WriteField("height", strconv.Itoa(a.Height))
	}
	if a.Caption != "" {
		writer.WriteField("caption", a.Caption)
	}
	if a.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if a.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(a.ReplyToMessageID))
	}
	if a.ReplyMarkup != "" {
		writer.WriteField("reply_markup", a.ReplyMarkup)
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return nil, nil, err
	}

	return writer, buffer, nil
}

type fileUploader interface {
	toMultiPart() (*multipart.Writer, *bytes.Buffer, error)
	methodName() string
}

type jsonUploader interface {
	toJSON() ([]byte, error)
	methodName() string
}
