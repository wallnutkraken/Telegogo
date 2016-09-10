package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendNewDocument(t *testing.T) {
	ass := assert.New(t)
	bot, err := NewClient(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendDocument(SendDocumentArgs{ChatID: testingID,
		DocumentPath: "testfiles/testdocument.txt"})
	ass.NoError(err)
	ass.NotZero(msg.Document.FileSize, "No file attached in returned message")
}

func TestResendDocument(t *testing.T) {
	const docID = "BQADBAADPAADjT7dDavEmb2gmPHeAg"
	ass := assert.New(t)
	bot, err := NewClient(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendDocument(SendDocumentArgs{ChatID: testingID, DocumentFileID: docID})
	ass.NoError(err)
	ass.NotZero(msg.Document.FileSize, "No document in returned message")
}
