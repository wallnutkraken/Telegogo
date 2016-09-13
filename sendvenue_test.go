package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendVenue(t *testing.T) {
	const latitude float32 = 35.852124
	const longitude float32 = -95.317949
	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendVenue(SendVenueArgs{ChatID: testingID, Title: "testVenue",
		Longitude: longitude, Latitude: latitude})
	ass.NoError(err)

	ass.Equal(msg.Venue.Location.Longitude, longitude, "Longitude is different")
	ass.Equal(msg.Venue.Location.Latitude, latitude, "Latitude is different")
}
