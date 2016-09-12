package TeleGogo

import (
	"bytes"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.telegram.org/"
	bot     = "bot"
	file    = "file"
)

// CreateBotGet returns http GET request with the provided token and method name
func createBotGet(token string, methodName string) (*http.Request, error) {
	return http.NewRequest("GET", baseURL+bot+token+"/"+methodName, nil)
}

// CreateFileGet returns hyype GET request with the provided token and method name for files
func createFileGet(token string, filepath string) (*http.Request, error) {
	return http.NewRequest("GET", baseURL+file+token+"/"+filepath, nil)
}

// CreateBotPostJSON Creates a POST to the Telegram API with the given parameters
func createBotPostJSON(token string, methodname string, encodedJSON []byte) (*http.Request, error) {
	request, err := http.NewRequest("POST", baseURL+bot+token+"/"+methodname, bytes.NewBuffer(encodedJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return request, err
}

// CreateBotPOST Creates a POST to the Telegram API with no parameters
func createBotPOST(token string, methodname string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("POST", baseURL+bot+token+"/"+methodname, body)
}
