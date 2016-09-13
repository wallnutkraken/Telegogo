package TeleGogo

// SendVenueArgs represents the optional and required arguments to SendVenue.
type SendVenueArgs struct {
	// ChatID Required. Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID string `json:"chat_id"`
	// Longitude Required. Longitude of the venue
	Longitude float32 `json:"longitude"`
	// Latitude Required. Latitude of the venue
	Latitude float32 `json:"latitude"`
	// Title Required. Name of the venue.
	Title string `json:"title"`
	// Address Required. Address of the venue
	Address string `json:"address"`
	// FoursquareID Optional. Foursquare identifier of the venue
	ForsquareID string `json:"foursquare_id"`
}

func (a SendVenueArgs) methodName() string {
	return "sendVenue"
}
