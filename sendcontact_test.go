package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendContact(t *testing.T) {
	const phoneNumber = "+12025550143"
	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendContact(SendContactArgs{ChatID: testingID, FirstName: "Dondo",
		PhoneNumber: phoneNumber})
	ass.NoError(err)

	ass.Equal(phoneNumber, msg.Contact.PhoneNumber, "Phone Number is not equal")
}
