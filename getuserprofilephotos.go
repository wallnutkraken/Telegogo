package TeleGogo

// GetUserProfilePhotosArgs represents the optional and required arguments to the GetUserProfilePhotos method
type GetUserProfilePhotosArgs struct {
	// ChatID Required. Unique identifier of the target user
	UserID string `json:"user_id"`
	// Offset Optional. Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int `json:"offset,omitempty"`
	// Limit Optional. Limits the number of photos to be retrieved. Values between 1—100 are accepted.
	// Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

func (a GetUserProfilePhotosArgs) methodName() string {
	return "getUserProfilePhotos"
}

type responseProfilePhotos struct {
	OK     bool              `json:"ok"`
	Result UserProfilePhotos `json:"result"`
}
