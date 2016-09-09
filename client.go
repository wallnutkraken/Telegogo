package TeleGogo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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

	defer httpResponse.Body.Close()
	responseObj := updateResponse{}
	decoder := json.NewDecoder(httpResponse.Body)

	err = decoder.Decode(&responseObj)
	if err != nil {
		return nil, err
	}

	return responseObj.Result, err
}

// SetWebhook NOT TESTED. Use this method to specify a url and receive incoming updates via an outgoing
// webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a
// reasonable amount of attempts.
func (c *client) SetWebhook(args SetWebhookArgs) error {
	postBody, buffer, err := args.toMultiform()

	if err != nil {
		return err
	}
	req, err := Requests.CreateBotPOST(c.token, "setWebook", buffer)
	req.Header.Set("Content-Type", postBody.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := resp.Body.Read(responseBuffer)
		return fmt.Errorf("Bad response: %s; (%s)", resp.Status, string(responseBuffer[:len]))
	}
	resp.Body.Close()
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

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := response.Body.Read(responseBuffer)
		return User{}, fmt.Errorf("Bad response: %s; (%s)", response.Status, string(responseBuffer[:len]))
	}
	tgResp := userResponse{}
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&tgResp)

	return tgResp.Result, err
}

// SendMessage sends a message with the specified arguments. On success returns the sent Message.
func (c *client) SendMessage(args SendMessageArgs) (Message, error) {
	jsonBytes, err := args.toJSON()
	if err != nil {
		return Message{}, err
	}
	get, err := Requests.CreateBotPostJSON(c.token, "sendMessage", jsonBytes)
	if err != nil {
		return Message{}, err
	}
	tgResponse, err := c.httpClient.Do(get)
	if err != nil {
		return Message{}, err
	}
	defer tgResponse.Body.Close()
	if tgResponse.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := tgResponse.Body.Read(responseBuffer)
		return Message{}, fmt.Errorf("Bad response: %s; (%s)", tgResponse.Status, string(responseBuffer[:len]))
	}

	decoder := json.NewDecoder(tgResponse.Body)
	sentMsgResponse := messageReply{}

	err = decoder.Decode(&sentMsgResponse)

	return sentMsgResponse.Result, err
}

func (c *client) ForwardMessage(args ForwardMessageArgs) (Message, error) {
	jsonBytes, err := args.toJSON()
	if err != nil {
		return Message{}, err
	}
	get, err := Requests.CreateBotPostJSON(c.token, "forwardMessage", jsonBytes)
	if err != nil {
		return Message{}, err
	}
	tgResponse, err := c.httpClient.Do(get)
	if err != nil {
		return Message{}, err
	}
	defer tgResponse.Body.Close()
	if tgResponse.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := tgResponse.Body.Read(responseBuffer)
		return Message{}, fmt.Errorf("Bad response: %s; (%s)", tgResponse.Status, string(responseBuffer[:len]))
	}

	decoder := json.NewDecoder(tgResponse.Body)
	sentMsgResponse := messageReply{}

	err = decoder.Decode(&sentMsgResponse)

	return sentMsgResponse.Result, err
}

func (c *client) SendPhoto(args SendPhotoArgs) (Message, error) {
	writer, buffer, err := args.toMultiPart()

	m := Message{}
	if err != nil {
		return m, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", args.ChatID)

	/* Optional args */
	if args.Caption != "" {
		writer.WriteField("caption", args.Caption)
	}
	if args.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if args.ReplyMarkup != "" {
		writer.WriteField("reply_markup", args.ReplyMarkup)
	}
	if args.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(args.ReplyToMessageID))
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return m, err
	}

	post, err := Requests.CreateBotPOST(c.token, "sendPhoto", buffer)
	post.Header.Add("Content-Type", writer.FormDataContentType())

	/* Send request */
	response, err := c.httpClient.Do(post)
	if err != nil {
		return m, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := response.Body.Read(responseBuffer)
		return m, fmt.Errorf("Bad response: %s; (%s)", response.Status, string(responseBuffer[:len]))
	}
	/* Read reply */
	msgReply := messageReply{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&msgReply); err != nil {
		return m, err
	}

	return msgReply.Result, err
}

func (c *client) SendAudio(args SendAudioArgs) (Message, error) {
	writer, buffer, err := args.toMultiPart()

	m := Message{}
	if err != nil {
		return m, err
	}
	/* Add ChatID */
	writer.WriteField("chat_id", args.ChatID)

	/* Optional args */
	if args.Duration != 0 {
		writer.WriteField("duration", strconv.Itoa(args.Duration))
	}
	if args.Performer != "" {
		writer.WriteField("performer", args.Performer)
	}
	if args.Title != "" {
		writer.WriteField("title", args.Title)
	}
	if args.DisableNotification == true {
		writer.WriteField("disable_notification", "true")
	}
	if args.ReplyMarkup != "" {
		writer.WriteField("reply_markup", args.ReplyMarkup)
	}
	if args.ReplyToMessageID != 0 {
		writer.WriteField("reply_to_message_id", strconv.Itoa(args.ReplyToMessageID))
	}

	/* Close the writer; we're done with the body of the request. */
	if err = writer.Close(); err != nil {
		return m, err
	}

	post, err := Requests.CreateBotPOST(c.token, "sendAudio", buffer)
	post.Header.Add("Content-Type", writer.FormDataContentType())

	/* Send request */
	response, err := c.httpClient.Do(post)
	if err != nil {
		return m, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		var responseBuffer = make([]byte, 1024)
		len, _ := response.Body.Read(responseBuffer)
		return m, fmt.Errorf("Bad response: %s; (%s)", response.Status, string(responseBuffer[:len]))
	}
	/* Read reply */
	msgReply := messageReply{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&msgReply); err != nil {
		return m, err
	}

	return msgReply.Result, err
}

// ResendPhoto Use this method to send photos already on the Telegram servers.
// On success, the sent Message is returned.
func (c *client) ResendPhoto(args ResendPhotoArgs) (Message, error) {
	jsonBytes, err := args.toJSON()

	m := Message{}
	if err != nil {
		return m, err
	}
	post, err := Requests.CreateBotPostJSON(c.token, "sendPhoto", jsonBytes)
	if err != nil {
		return m, err
	}
	tgResponse, err := c.httpClient.Do(post)
	if err != nil {
		return m, err
	}
	defer tgResponse.Body.Close()
	if tgResponse.StatusCode != http.StatusOK {
		var buffer = make([]byte, 1024)
		len, _ := tgResponse.Body.Read(buffer)
		return m, fmt.Errorf("Bad response: %s; (%s)", tgResponse.Status, string(buffer[:len]))
	}
	decoder := json.NewDecoder(tgResponse.Body)
	msgResponse := messageReply{}
	if err = decoder.Decode(&msgResponse); err != nil {
		return m, err
	}
	return msgResponse.Result, nil
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
	ResendPhoto(ResendPhotoArgs) (Message, error)
	SendPhoto(SendPhotoArgs) (Message, error)
	SendAudio(SendAudioArgs) (Message, error)
}
