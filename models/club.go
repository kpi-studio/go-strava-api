package models

// Club represents a Strava club
type Club struct {
	ID              int64         `json:"id"`
	ResourceState   ResourceState `json:"resource_state"`
	Name            string        `json:"name"`
	ProfileMedium   string        `json:"profile_medium"`
	Profile         string        `json:"profile"`
	CoverPhoto      string        `json:"cover_photo"`
	CoverPhotoSmall string        `json:"cover_photo_small"`
	SportType       string        `json:"sport_type"`
	City            string        `json:"city"`
	State           string        `json:"state"`
	Country         string        `json:"country"`
	Private         bool          `json:"private"`
	MemberCount     int           `json:"member_count"`
	Featured        bool          `json:"featured"`
	Verified        bool          `json:"verified"`
	URL             string        `json:"url"`
	Owner           bool          `json:"owner"`
	Admin           bool          `json:"admin"`
}

// ClubMembership represents a club membership status
type ClubMembership struct {
	Success    bool   `json:"success"`
	Active     bool   `json:"active"`
	Membership string `json:"membership"`
}