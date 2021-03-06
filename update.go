package TeleGogo

import "encoding/json"

// Update represents a state update in Telegram that has relevance to a bot client.
type Update struct {
	// ID The update‘s unique identifier
	ID int64 `json:"update_id"`
	// Message Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message Message `json:"message"`
	// EditedMessage Optional. New version of a message that is known to the bot and was edited
	EditedMessage Message `json:"edited_message"`
	// Inline Optional. New incoming inline query
	Inline InlineQuery `json:"inline_query"`
	// InlineResult Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	InlineResult ChosenInlineResult `json:"chosen_inline_result"`
	// Callback Optional. New incoming callback query
	Callback CallbackQuery `json:"callback_query"`
}

type updateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// GetUpdatesOptions represents the required and optional arguments to the GetUpdates method for a bot.
type GetUpdatesOptions struct {
	// Offset Optional. Identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers of previously received updates.
	Offset int64 `json:"offset"`
	// Limit Optional. Limits the number of updates to be retrieved. Values between 1—100 are accepted.
	Limit int `json:"limit,omitempty"`
	// Timeout Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling.
	Timeout int `json:"timeout,omitempty"`
}

func (u GetUpdatesOptions) toJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u GetUpdatesOptions) methodName() string {
	return "getUpdates"
}
