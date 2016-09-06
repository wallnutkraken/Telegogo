package TeleGogo

type client struct {
	token string
}

func (c *client) getToken() string {
	return c.token
}

// NewClient Creates a new Client
func NewClient(token string) (Client, error) {
	c := new(client)
	c.token = token
	return c, nil
}

// Client represents a bot in Telegram.
type Client interface {
	getToken() string
}
