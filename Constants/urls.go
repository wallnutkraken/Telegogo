package Constants

const (
	baseURL = "https://api.telegram.org/bot"
)

// CreateGet returns a URL for a GET request using the token and method provided.
func CreateGet(token string, methodName string) string {
	return baseURL + token + "/" + methodName
}
