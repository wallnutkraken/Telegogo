package Requests

import (
	"bytes"
	"net/http"
)

const (
	baseURL = "https://api.telegram.org/"
	bot     = "bot"
	file    = "file"
)

// Arg represents a GET argument
type Arg struct {
	Name  string
	Value string
}

// CreateBotGet returns http GET request with the provided token and method name
func CreateBotGet(token string, methodName string) (*http.Request, error) {
	return http.NewRequest("GET", baseURL+bot+token+"/"+methodName, nil)
}

// CreateBotGetWithArgs returns http GET request with the provided token, method name and given arguments
func CreateBotGetWithArgs(token string, methodname string, args ...Arg) (*http.Request, error) {
	request, err := http.NewRequest("GET", baseURL+bot+token+"/"+methodname, nil)
	if err != nil {
		return request, err
	}

	query := request.URL.Query()
	for _, argument := range args {
		query.Add(argument.Name, argument.Value)
	}

	return request, nil
}

// CreateFileGet returns hyype GET request with the provided token and method name for files
func CreateFileGet(token string, filepath string) (*http.Request, error) {
	return http.NewRequest("GET", baseURL+file+token+"/"+filepath, nil)
}

// CreateBotPostJSON Creates a POST to the Telegram API with the given parameters
func CreateBotPostJSON(token string, methodname string, encodedJSON []byte) (*http.Request, error) {
	request, err := http.NewRequest("POST", baseURL+bot+token+"/"+methodname, bytes.NewBuffer(encodedJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return request, err
}
