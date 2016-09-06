package TeleGogo

import (
	"net/http"
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
}
