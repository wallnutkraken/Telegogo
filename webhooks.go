package TeleGogo

// SetWebhookArgs represents the optional and required arguments for the SetWebhook method
type SetWebhookArgs struct {
	// URL Optional. HTTPS url to send updates to. Use an empty string to remove webhook integration
	URL string `tele:"url"`
	// Certificate Optional. The path for your public key certificate so that the root
	// certificate in use can be checked.
	CertificatePath string `tele:!certificate`
}

func (a SetWebhookArgs) methodName() string {
	return "setWebhook"
}
