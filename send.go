package TeleGogo

import (
	"net/http"

	"github.com/wallnutkraken/TeleGogo/Requests"
)

func (c *client) sendFile(uploader fileUploader) (*http.Response, error) {
	writer, buffer, err := uploader.toMultiPart()
	if err != nil {
		return nil, err
	}
	post, err := Requests.CreateBotPOST(c.token, uploader.methodName(), buffer)
	post.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(post)
}

func (c *client) sendJSON(uploader jsonUploader) (*http.Response, error) {
	jsonBytes, err := uploader.toJSON()
	if err != nil {
		return nil, err
	}
	post, err := Requests.CreateBotPostJSON(c.token, uploader.methodName(), jsonBytes)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(post)
}