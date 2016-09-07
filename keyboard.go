package TeleGogo

// ReplyKeyboardMarkup represents a custom keyboard with reply options
type ReplyKeyboardMarkup struct {
	// Keyboard Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard []KeyboardButton `json:"keyboard"`
	// ResizeKeyboard Optional. Requests clients to resize the keyboard vertically for optimal fit
	ResizeKeyboard bool `json:"resize_keyboard"`
	// OneTimeKeyboard Optional. Requests clients to hide the keyboard as soon as it's been used
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	// Selective Optional. Use this parameter if you want to show the keyboard to specific users only.
	Selective bool `json:"selective"`
}

// KeyboardButton represents one button of the reply keyboard
type KeyboardButton struct {
	// Text of the button.
	Text string `json:"text"`
	// RequestContact Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	RequestContact bool `json:"request_contact"`
	// RequestLocation Optional. If True, the user's current location will be sent when the button is pressed.
	RequestLocation bool `json:"request_location"`
}

// ReplyKeyboardHide Upon receiving a message with this object, Telegram clients will hide the current
// custom keyboard and display the default letter-keyboard.
type ReplyKeyboardHide struct {
	// HideKeyboard Requests clients to hide the custom keyboard
	HideKeyboard bool `json:"hide_keyboard"`
	// Selective Optional. Use this parameter if you want to hide keyboard for specific users only.
	Selective bool `json:"selective"`
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text Label text on the button
	Text string `json:"text"`
	// URL Optional. HTTP url to be opened when button is pressed
	URL string `json:"url"`
	// CallbackData Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data"`
	// SwitchInlineQuery Optional. If set, pressing the button will prompt the user to select one of
	// their chats, open that chat and insert the bot‘s username and the specified inline query in the input
	// field. Can be empty, in which case just the bot’s username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
