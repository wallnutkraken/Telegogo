package TeleGogo

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
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
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	certificateFile, err := os.Open(a.CertificatePath)
	if err != nil {
		return nil, nil, err
	}
	defer certificateFile.Close()
	cert, err := writer.CreateFormFile("certificate", a.CertificatePath)
	if err != nil {
		return nil, nil, err
	}
	if _, err = io.Copy(cert, certificateFile); err != nil {
		return nil, nil, err
	}
	if cert, err = writer.CreateFormField("url"); err != nil {
		return nil, nil, err
	}
	writer.Close()

	return writer, &buffer, nil
}
