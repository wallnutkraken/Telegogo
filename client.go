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

	defer httpResponse.Body.Close()
	responseObj := updateResponse{}
	decoder := json.NewDecoder(httpResponse.Body)

	err = decoder.Decode(&responseObj)
	if err != nil {
		return nil, err
	}

	return responseObj.Result, err
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

	decoder := json.NewDecoder(tgResponse.Body)
	sentMsgResponse := messageReply{}

	err = decoder.Decode(&sentMsgResponse)

	return sentMsgResponse.Result, err
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
}
