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

func (c *client) RecieveUpdates() {
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
}
