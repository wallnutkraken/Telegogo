package Requests

import "net/http"

const (
	baseURL = "https://api.telegram.org/bot"
)

// Arg represents a GET argument
type Arg struct {
	Name  string
	Value string
}

// CreateGet returns http GET request with the provided token and method name
func CreateGet(token string, methodName string) (*http.Request, error) {
	return http.NewRequest("GET", baseURL+token+"/"+methodName, nil)
}

// CreateGetWithArgs returns http GET request with the provided token, method name and given arguments
func CreateGetWithArgs(token string, methodname string, args ...Arg) (*http.Request, error) {
	request, err := http.NewRequest("GET", baseURL+token+"/"+methodname, nil)
	if err != nil {
		return request, err
	}

	query := request.URL.Query()
	for _, argument := range args {
		query.Add(argument.Name, argument.Value)
	}

	return request, nil
}
