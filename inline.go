package TeleGogo

// InlineQuery represents an incoming inline query. When the user sends an empty query,
// your bot could return some default or trending results.
type InlineQuery struct {
	// ID Unique identifier for this query
	ID string `json:"id"`
	// From Sender
	From User `json:"from"`
	// Location Optional. Sender location, only for bots that request user location
	Location Location `json:"location"`
	// Query Text of the query (up to 512 characters)
	Query string `json:"query"`
	// Offset Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ResultID The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`
	// From The user that chose the result
	From User `json:"from"`
	// Location Optional. Sender location, only for bots that require user location
	Location Location `json:"Optional. Sender location, only for bots that require user location"`
	// InlineMessageID Optional. Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	InlineMessageID string `json:"inline_message_id"`
	// Query The query that was used to obtain the result
	Query string `json:"query"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	// ID Unique identifier for this query
	ID string `json:"id"`
	// From Sender
	From User `json:"from"`
	// Message Optional. Message with the callback button that originated the query.
	Message Message `json:"message"`
	// InlineMessageID Optional. Identifier of the message sent via the bot in inline mode, that originated the query
	InlineMessageID string `json:"inline_message_id"`
	// Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field
	Data string `json:"data"`
}
