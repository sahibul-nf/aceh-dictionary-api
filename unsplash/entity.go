// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    unsplashResponse, err := UnmarshalUnsplashResponse(bytes)
//    bytes, err = unsplashResponse.Marshal()

package unsplash

import "encoding/json"

type UnsplashResponse []UnsplashResponseElement

func UnmarshalUnsplashResponse(data []byte) (UnsplashResponse, error) {
	var r UnsplashResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UnsplashResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UnsplashResponseElement struct {
	ID                     string                `json:"id"`
	CreatedAt              string                `json:"created_at"`
	UpdatedAt              string                `json:"updated_at"`
	PromotedAt             string                `json:"promoted_at"`
	Width                  int64                 `json:"width"`
	Height                 int64                 `json:"height"`
	Color                  string                `json:"color"`
	BlurHash               string                `json:"blur_hash"`
	Description            *string               `json:"description"`
	AltDescription         string                `json:"alt_description"`
	Urls                   Urls                  `json:"urls"`
	Links                  UnsplashResponseLinks `json:"links"`
	Likes                  int64                 `json:"likes"`
	LikedByUser            bool                  `json:"liked_by_user"`
	CurrentUserCollections []interface{}         `json:"current_user_collections"`
	Sponsorship            interface{}           `json:"sponsorship"`
	TopicSubmissions       TopicSubmissions      `json:"topic_submissions"`
	User                   User                  `json:"user"`
	Exif                   Exif                  `json:"exif"`
	Location               Location              `json:"location"`
	Views                  int64                 `json:"views"`
	Downloads              int64                 `json:"downloads"`
}

type Exif struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	ISO          int64  `json:"iso"`
}

type UnsplashResponseLinks struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type Location struct {
	Name     interface{} `json:"name"`
	City     interface{} `json:"city"`
	Country  interface{} `json:"country"`
	Position Position    `json:"position"`
}

type Position struct {
	Latitude  interface{} `json:"latitude"`
	Longitude interface{} `json:"longitude"`
}

type TopicSubmissions struct {
	Nature  *Animals `json:"nature,omitempty"`
	Animals *Animals `json:"animals,omitempty"`
}

type Animals struct {
	Status     string `json:"status"`
	ApprovedOn string `json:"approved_on"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
	SmallS3 string `json:"small_s3"`
}

type User struct {
	ID                string       `json:"id"`
	UpdatedAt         string       `json:"updated_at"`
	Username          string       `json:"username"`
	Name              string       `json:"name"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	TwitterUsername   string       `json:"twitter_username"`
	PortfolioURL      string       `json:"portfolio_url"`
	Bio               interface{}  `json:"bio"`
	Location          string       `json:"location"`
	Links             UserLinks    `json:"links"`
	ProfileImage      ProfileImage `json:"profile_image"`
	InstagramUsername string       `json:"instagram_username"`
	TotalCollections  int64        `json:"total_collections"`
	TotalLikes        int64        `json:"total_likes"`
	TotalPhotos       int64        `json:"total_photos"`
	AcceptedTos       bool         `json:"accepted_tos"`
	ForHire           bool         `json:"for_hire"`
	Social            Social       `json:"social"`
}

type UserLinks struct {
	Self      string `json:"self"`
	HTML      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"likes"`
	Portfolio string `json:"portfolio"`
	Following string `json:"following"`
	Followers string `json:"followers"`
}

type ProfileImage struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Social struct {
	InstagramUsername string      `json:"instagram_username"`
	PortfolioURL      string      `json:"portfolio_url"`
	TwitterUsername   string      `json:"twitter_username"`
	PaypalEmail       interface{} `json:"paypal_email"`
}
