package TeleGogo

// User represents a Telegram user or bot.
type User struct {
	// ID Unique identifier for this user or bot
	ID int `json:"id"`
	// FirstName User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// LastName Optional. User‘s or bot’s last name
	LastName string `json:"last_name"`
	// Username Optional. User‘s or bot’s username
	Username string `json:"username"`
}