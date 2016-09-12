package TeleGogo

import (
	"encoding/json"
	"net/http"
)

func (c *client) sendFile(uploader apiCaller) (*http.Response, error) {
	writer, buffer, err := toMultiPart(uploader)
	if err != nil {
		return nil, err
	}
	post, err := createBotPOST(c.token, uploader.methodName(), buffer)
	post.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(post)
}

func (c *client) sendJSON(uploader apiCaller) (*http.Response, error) {
	jsonBytes, err := json.Marshal(uploader)
	if err != nil {
		return nil, err
	}
	post, err := createBotPostJSON(c.token, uploader.methodName(), jsonBytes)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(post)
}

func (c *client) sendFileMessage(uploader apiCaller) (Message, error) {
	response, err := c.sendFile(uploader)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}

	return responseToMessage(response)
}

func (c *client) sendJSONMessage(uploader apiCaller) (Message, error) {
	response, err := c.sendJSON(uploader)
	if err != nil {
		return Message{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToError(response)
	}

	return responseToMessage(response)
}
