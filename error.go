package TeleGogo

import (
	"fmt"
	"net/http"
)

func responseToError(response *http.Response) error {
	var responseBuffer = make([]byte, 1024)
	len, _ := response.Body.Read(responseBuffer)
	response.Body.Close()
	return fmt.Errorf("Bad response: %s; (%s)", response.Status, string(responseBuffer[:len]))
}
