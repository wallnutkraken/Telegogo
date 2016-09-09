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
	ass := assert.New(t)
	bot, err := NewClient(testBotToken)
	ass.NoError(err)

	updates, err := bot.GetUpdates(GetUpdatesOptions{})
	var update *Document
	for _, upd := range updates {
		if upd.Message.Document.FileSize != 0 {
			update = &upd.Message.Document
		}
	}
	ass.NotNil(update, "Found no update with a Document in it; inconclusive")

	msg, err := bot.SendDocument(SendDocumentArgs{ChatID: testingID, DocumentFileID: update.FileID})
	ass.NoError(err)
	ass.NotZero(msg.Document.FileSize, "No document in returned message")
}
