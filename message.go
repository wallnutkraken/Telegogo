package TeleGogo

// Message represents a message.
type Message struct {
	// ID Unique message identifier
	ID int `json:"message_id"`
	// From Optional. Sender, can be empty for messages sent to channels
	From User `json:"from"`
	// Date Date the message was sent in Unix time
	Date int `json:"date"`
	// Chat Conversation the message belongs to
	Chat Chat `json:"chat"`
	// ForwardFrom Optional. For forwarded messages, sender of the original message
	ForwardFrom User `json:"forward_from"`
	// ForwardFromChat Optional. For messages forwarded from a channel, information about the original channel
	ForwardFromChat Chat `json:"forward_from_chat"`
	// ForwardDate Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date"`
	// ReplyToMessage Optional. For replies, the original message.
	ReplyToMessage *Message `json:"reply_to_message"`
	// EditDate Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date"`
	// Text Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text"`
	// Entities Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`
	// Audio Optional. Message is an audio file, information about the file
	Audio Audio `json:"audio"`
	// Document Optional. Message is a general file, information about the file
	Document Document `json:"document"`
	// Photo Optional. Message is a photo, available sizes of the photo
	Photo []PhotoSize `json:"photo"`
	// Sticker Optional. Message is a sticker, information about the sticker
	Sticker Sticker `json:"sticker"`
	// Video Optional. Message is a video, information about the video
	Video Video `json:"video"`
	// Voice Optional. Message is a voice message, information about the file
	Voice Voice `json:"voice"`
	// Caption Optional. Caption for the document, photo or video, 0-200 characters
	Caption string `json:"caption"`
	// Contact Optional. Message is a shared contact, information about the contact
	Contact Contact `json:"contact"`
	// Location Optional. Message is a shared location, information about the location
	Location Location `json:"location"`
	// Venue Optional. Message is a venue, information about the venue
	Venue Venue `json:"venue"`
	// NewChatMember Optional. A new member was added to the group, information about them
	NewChatMember User `json:"new_chat_member"`
	// LeftChatMember Optional. A member was removed from the group, information about them
	LeftChatMember User `json:"left_chat_member"`
	// NewChatTitle Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`
	// NewChatPhoto Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`
	// DeleteChatPhoto Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// GroupChatCreated Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`
	// SupergroupChatCreated Optional. Service message: the supergroup has been created.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	// ChannelChatCreated Optional. Service message: the channel has been created.
	ChannelChatCreated bool `json:"channel_chat_created"`
	// MigrateToChatID Optional. The group has been migrated to a supergroup with the specified identifier.
	MigrateToChatID int `json:"migrate_to_chat_id"`
	// MigrateFromChatID Optional. The supergroup has been migrated from a group with the specified identifier.
	MigrateFromChatID int `json:"migrate_from_chat_id"`
	// PinnedMessage Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message"`
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Can be mention (@username), hashtag, bot_command, url, email,
	// bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block),
	// text_link (for clickable text URLs), text_mention (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// URL Optional. For “text_link” only, url that will be opened after user taps on the text
	URL string `json:"url"`
	// User Optional. For “text_mention” only, the mentioned user
	User User `json:"user"`
}

// ForceReply Upon receiving a message with this object, Telegram clients will display a reply interface to
// the user (act as if the user has selected the bot‘s message and tapped ’Reply').
type ForceReply struct {
	// ForceReply Shows reply interface to the user, as if they manually selected the bot‘s message and
	// tapped ’Reply'
	ForceReply bool `json:"force_reply"`
	// Selective Optional. Use this parameter if you want to force reply from specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective"`
}

type messageReply struct {
	OK     bool    `json:"ok"`
	Result Message `json:"result"`
}

// SendMessageArgs represents the optional and required arguments for the SendMessage method.
type SendMessageArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Text Required. Text of the message to be sent
	Text string `json:"text"`
	// ParseMode Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic,
	// fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`
	// DisableWebPagePreview Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// DisableNotification Optional. Sends the message silently. iOS users will not receive a notification,
	// Android users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendMessageArgs) methodName() string {
	return "sendMessage"
}

// ForwardMessageArgs represents the optional and required arguments for the ForwardMessage message
type ForwardMessageArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// FromChatID Required. Unique identifier for the chat where the original message was sent
	// (or channel username in the format @channelusername)
	FromChatID string `json:"from_chat_id"`
	// DisableNotification Optional. If true, sends the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// MessageID Required. Unique message identifier
	MessageID int `json:"message_id"`
}

func (f ForwardMessageArgs) methodName() string {
	return "forwardMessage"
}

// SendLocationArgs represents the optional and required arguments for the SendLocation message
type SendLocationArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Latitude Required. Latitude of location
	Latitude float32 `json:"latitude"`
	// Longitude. Required. Longitude of location
	Longitude float32 `json:"longitude"`
	// DisableNotification Optional. Sends the message silently. iOS users will not receive a notification,
	// Android users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ReplyToMessageID Optional. If the message is a reply, ID of the original message
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// ReplyMarkup Optional. Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to hide reply keyboard or to force a reply from the user.
	ReplyMarkup string `json:"reply_markup,omitempty"`
}

func (a SendLocationArgs) methodName() string {
	return "sendLocation"
}
