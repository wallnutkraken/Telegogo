package TeleGogo

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type client struct {
	token      string
	httpClient *http.Client
}

// GetUpdates receives incoming updates using long polling.
func (c *client) GetUpdates(options GetUpdatesOptions) ([]Update, error) {
	response, err := c.sendJSON(options)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, responseToError(response)
	}

	upd := updateResponse{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&upd); err != nil {
		return nil, err
	}
	return upd.Result, nil
}

// SetWebhook NOT TESTED. Use this method to specify a url and receive incoming updates via an outgoing
// webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a
// reasonable amount of attempts.
func (c *client) SetWebhook(args SetWebhookArgs) error {
	response, err := c.sendFile(args)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return responseToError(response)
	}
	response.Body.Close()

	return nil
}

// DownloadFile downloads the specified file
func (c *client) DownloadFile(file File, path string) error {
	get, err := createFileGet(c.token, file.FilePath)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(get)
	if err != nil {
		return err
	}
	physicalFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer physicalFile.Close()
	defer resp.Body.Close()

	_, err = io.Copy(physicalFile, resp.Body)
	return err
}

// WhoAmI A simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (c *client) WhoAmI() (User, error) {
	request, err := createBotGet(c.token, "getMe")

	if err != nil {
		return User{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return User{}, err
	}

	if response.StatusCode != http.StatusOK {
		return User{}, responseToError(response)
	}
	tgResp := userResponse{}
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&tgResp)
	response.Body.Close()

	return tgResp.Result, err
}

// SendMessage sends a message with the specified arguments. On success returns the sent Message.
func (c *client) SendMessage(args SendMessageArgs) (Message, error) {
	return c.sendJSONMessage(args)
}

func (c *client) ForwardMessage(args ForwardMessageArgs) (Message, error) {
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}

	return responseToMessage(response)
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (c *client) SendPhoto(args SendPhotoArgs) (Message, error) {
	/* Decide whether this is a newly uploaded file or an old one. */
	if args.FileID == "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

func (c *client) SendAudio(args SendAudioArgs) (Message, error) {
	/* Decide if it's a new or existing file, based on user intent */
	if args.AudioPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendDocument Use this method to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (c *client) SendDocument(args SendDocumentArgs) (Message, error) {
	/* Decide if the document is a resend or a new file, based on args */
	if args.DocumentPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendSticker Use this method to send .webp stickers. On success, the sent Message is returned.
func (c *client) SendSticker(args SendStickerArgs) (Message, error) {
	/* Decide if the sticker is a existing or a new sticker, based on args */
	if args.StickerPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendVideo Use this method to send video files,
// Telegram clients support mp4 videos (other formats may be sent as Document).
// On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (c *client) SendVideo(args SendVideoArgs) (Message, error) {
	/* Decide if the video is an existing or a new video, based on args */
	if args.VideoPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

func responseToMessage(response *http.Response) (Message, error) {
	msg := messageReply{}
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&msg)
	defer response.Body.Close()
	return msg.Result, err
}

// NewClient Creates a new Client
func NewClient(token string) (Client, error) {
	c := new(client)
	c.token = token
	c.httpClient = &http.Client{}
	return c, nil
}

// Client represents a bot in Telegram.
type Client interface {
	DownloadFile(File, string) error
	WhoAmI() (User, error)
	GetUpdates(GetUpdatesOptions) ([]Update, error)
	SendMessage(SendMessageArgs) (Message, error)
	ForwardMessage(ForwardMessageArgs) (Message, error)
	SetWebhook(SetWebhookArgs) error
	SendPhoto(SendPhotoArgs) (Message, error)
	SendAudio(SendAudioArgs) (Message, error)
	SendDocument(SendDocumentArgs) (Message, error)
	SendSticker(SendStickerArgs) (Message, error)
	SendVideo(SendVideoArgs) (Message, error)
}
