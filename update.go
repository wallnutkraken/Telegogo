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
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type CallbackQuery struct {
}
