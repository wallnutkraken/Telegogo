package TeleGogo

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/wallnutkraken/TeleGogo/Requests"
)

type client struct {
	token      string
	httpClient *http.Client
}

func (c *client) getToken() string {
	return c.token
}

// GetUpdates receives incoming updates using long polling.
func (c *client) GetUpdates(options GetUpdatesOptions) ([]Update, error) {
	get, err := Requests.CreateBotGetWithArgs(c.token, "getUpdates", options.toArgs()...)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.httpClient.Do(get)
	if err != nil {
		return nil, err
	}

	responseObj := updateResponse{}
	decoder := json.NewDecoder(httpResponse.Body)

	err = decoder.Decode(&responseObj)
	if err != nil {
		return nil, err
	}
	httpResponse.Body.Close()

	return responseObj.Result, err
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
	get, err := Requests.CreateFileGet(c.getToken(), file.FilePath)
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
	request, err := Requests.CreateBotGet(c.token, "getMe")

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
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}

	decoder := json.NewDecoder(response.Body)
	sentMsgResponse := messageReply{}

	err = decoder.Decode(&sentMsgResponse)
	response.Body.Close()

	return sentMsgResponse.Result, err
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

	decoder := json.NewDecoder(response.Body)
	sentMsgResponse := messageReply{}

	err = decoder.Decode(&sentMsgResponse)

	return sentMsgResponse.Result, err
}

func (c *client) sendNewPhoto(args SendPhotoArgs) (Message, error) {
	response, err := c.sendFile(args)
	if err != nil {
		return Message{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}
	/* Read reply */
	msgReply := messageReply{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&msgReply); err != nil {
		return Message{}, err
	}

	response.Body.Close()
	return msgReply.Result, err
}

func (c *client) sendExistingPhoto(args SendPhotoArgs) (Message, error) {
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}
	msg := messageReply{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&msg); err != nil {
		return Message{}, err
	}
	response.Body.Close()
	return msg.Result, err
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (c *client) SendPhoto(args SendPhotoArgs) (Message, error) {
	/* Decide whether this is a newly uploaded file or an old one. */
	if args.FileID == "" {
		return c.sendNewPhoto(args)
	}
	return c.sendExistingPhoto(args)
}

func (c *client) sendNewAudio(args SendAudioArgs) (Message, error) {
	response, err := c.sendFile(args)
	if err != nil {
		return Message{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}
	/* Read reply */
	msgReply := messageReply{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&msgReply); err != nil {
		return Message{}, err
	}
	response.Body.Close()

	return msgReply.Result, err
}

func (c *client) sendExistingAudio(args SendAudioArgs) (Message, error) {
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}
	/* Parse reply */
	return responseToMessage(response)
}

func (c *client) SendAudio(args SendAudioArgs) (Message, error) {
	/* Decide if it's a new or existing file, based on user intent */
	if args.AudioFileID != "" {
		return c.sendNewAudio(args)
	}
	return c.sendExistingAudio(args)
}

func (c *client) resendPhoto(args SendPhotoArgs) (Message, error) {
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}
	decoder := json.NewDecoder(response.Body)
	msgResponse := messageReply{}
	if err = decoder.Decode(&msgResponse); err != nil {
		return Message{}, err
	}
	response.Body.Close()

	return msgResponse.Result, nil
}

func responseToMessage(response *http.Response) (Message, error) {
	msg := messageReply{}
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(msg)
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
	getToken() string
	DownloadFile(File, string) error
	WhoAmI() (User, error)
	GetUpdates(GetUpdatesOptions) ([]Update, error)
	SendMessage(SendMessageArgs) (Message, error)
	ForwardMessage(ForwardMessageArgs) (Message, error)
	SetWebhook(SetWebhookArgs) error
	SendPhoto(SendPhotoArgs) (Message, error)
	SendAudio(SendAudioArgs) (Message, error)
}
