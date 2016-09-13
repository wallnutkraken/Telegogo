package TeleGogo

// Location represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float32 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float32 `json:"latitude"`
}
