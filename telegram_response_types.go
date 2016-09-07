package TeleGogo

type userResponse struct {
	OK     bool `json:"ok"`
	Result User `json:"result"`
}
