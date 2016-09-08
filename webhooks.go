package TeleGogo

import (
	"bytes"
	"mime/multipart"
)

// SetWebhookArgs represents the optional and required arguments for the SetWebhook method
type SetWebhookArgs struct {
	// URL Optional. HTTPS url to send updates to. Use an empty string to remove webhook integration
	URL string
	// Certificate Optional. The path for your public key certificate so that the root
	// certificate in use can be checked.
	CertificatePath string
}

func (a *SetWebhookArgs) toPOSTBody() (*multipart.Writer, *bytes.Buffer, error) {
	writer, buffer, err := createInputFileBody(a.CertificatePath, "certificate")

	url, err := writer.CreateFormField("url")
	if err != nil {
		return nil, nil, err
	}
	url.Write([]byte(a.URL))
	writer.Close()

	return writer, buffer, nil
}
