package TeleGogo

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// createInputFileBody reads a file and creates a multipart writer for POSTing with the file and
// the form name
func createInputFileBody(path string, formName string) (*multipart.Writer, *bytes.Buffer, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	fileUpload, err := writer.CreateFormFile(formName, path)
	if err != nil {
		return nil, nil, err
	}
	if _, err := io.Copy(fileUpload, file); err != nil {
		return nil, nil, err
	}
	return writer, &buffer, nil
}
