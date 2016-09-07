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

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	// Count Total number of profile pictures the target user has
	Count int `json:"total_count"`
	// Photos Requested profile pictures (in up to 4 sizes each)
	Photos []PhotoSize `json:"photos"`
}
