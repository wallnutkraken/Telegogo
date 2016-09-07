package TeleGogo

// Chat represents a chat.
type Chat struct {
	// ID Unique identifier for this chat.
	ID int `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Title Optional. Title, for channels and group chats
	Title string `json:"title"`
	// Username Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username"`
	// FirstName Optional. First name of the other party in a private chat
	FirstName string `json:"first_name"`
	// LastName Optional. Last name of the other party in a private chat
	LastName string `json:"last_name"`
}
