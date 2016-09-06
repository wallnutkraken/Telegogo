package TeleGogo

type Update struct {
	Id            int                `json:"update_id"`
	Message       Message            `json:"message"`
	EditedMessage Message            `json:"edited_message"`
	Inline        InlineQuery        `json:"inline_query"`
	InlineResult  ChosenInlineResult `json:"chosen_inline_result"`
	Callback      CallbackQuery      `json:"callback_query"`
}

type Message struct {
	Id              int  `json:"message_id"`
	From            User `json:"from"`
	Date            int  `json:"date"`
	Chat            Chat `json:"chat"`
	ForwardFrom     User `json:"forward_from"`
	ForwardFromChat Chat `json:"forward_from_chat"`
	ForwardDate     int  `json:"forward_date"`
	//unfinished
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type CallbackQuery struct {
}

type User struct {
}

type Chat struct {
}
