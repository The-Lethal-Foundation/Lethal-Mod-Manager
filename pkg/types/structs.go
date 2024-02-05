package types

type ProgressCallback func(current, total int, title string)

type ModName struct {
	Author  string
	Name    string
	Version string
}

// API response for getting the latest mod version
type ModInfoResponse struct {
	Downloads     int    `json:"downloads"`
	RatingScore   int    `json:"rating_score"`
	LatestVersion string `json:"latest_version"`
}
