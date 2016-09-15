package TeleGogo

import (
	"encoding/json"
	"net/http"
)

func responseToError(response *http.Response) error {
	return responseToTgError(response)
}

func responseToTgError(response *http.Response) TelegramError {
	defer response.Body.Close()
	tgErr := telegramError{}
	json.NewDecoder(response.Body).Decode(&tgErr)
	return tgErr
}

func errToTelegramErr(err error) TelegramError {
	return telegramError{OK: false, ErrorCode: 0, Description: err.Error()}
}

type telegramError struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (t telegramError) IsOK() bool {
	return t.OK
}

func (t telegramError) ErrCode() int {
	return t.ErrorCode
}

func (t telegramError) Error() string {
	return t.Description
}

// TelegramError represents a non-OK response from Telegram. Can also be used as type error.
type TelegramError interface {
	IsOK() bool
	// ErrCode is 0 if error is not a Telegram API Error response
	ErrCode() int
	Error() string
}
