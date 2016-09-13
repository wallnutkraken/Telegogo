package TeleGogo

// SendContactArgs represents the optional and required arguments to SendContact.
type SendContactArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// PhoneNumber Required. Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// FirstName Required. Contact's first Name
	FirstName string `json:"first_name"`
	// LastName Optional. Contact's last Name
	LastName string `json:"last_name"`
	// DisableNotification Optional. Sends the message silently. iOS users will not receive a notification,
	// Android users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendContactArgs) methodName() string {
	return "sendContact"
}
