package TeleGogo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
)

// createInputFileBody reads a file and creates a multipart writer for POSTing with the file and
// the form name
func createInputFileBody(path string, formName string) (*multipart.Writer, *bytes.Buffer, *os.File, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}

	fileUpload, err := writer.CreateFormFile(formName, path)
	if err != nil {
		return nil, nil, nil, err
	}
	if _, err := io.Copy(fileUpload, file); err != nil {
		return nil, nil, nil, err
	}
	return writer, &buffer, file, nil
}

// ResendPhotoArgs represents the optional and required arguments used when resending an already
// uploaded photo.
type ResendPhotoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// PhotoID Required. Photo to send. Pass a file ID to resend a photo that is already on the Telegram servers
	PhotoID string `json:"photo"`
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

func (a *ResendPhotoArgs) toJSON() ([]byte, error) {
	return json.Marshal(*a)
}

// SendPhotoArgs represents the optional and required arguments for the SendPhoto method
type SendPhotoArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// PhotoID Required. Path to photo file on device.
	PhotoPath string
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

func (a *SendPhotoArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, *os.File, error) {
	writer, buffer, file, err := createInputFileBody(a.PhotoPath, "photo")
	if err != nil {
		return nil, nil, nil, err
	}
	return writer, buffer, file, nil
}

// SendAudioArgs represents the optional and required arguments for the SendAudio method
type SendAudioArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string
	// AudioID Required. Path to audio file on device.
	AudioPath string
	// Duration Optional. Duration of the audio in seconds
	Duration int
	// Performer Optional.
	Performer string
	// Title Optional. Track name
	Title string
	// DisableNotification Optional. Sends the message silently.
	DisableNotification bool
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string
}

func (a *SendAudioArgs) toMultiPart() (*multipart.Writer, *bytes.Buffer, *os.File, error) {
	writer, buffer, file, err := createInputFileBody(a.AudioPath, "audio")
	if err != nil {
		return nil, nil, nil, err
	}
	return writer, buffer, file, nil
}
