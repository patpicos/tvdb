package tvdb

type loginAPIResponse struct {
	Token string `json:"token"`
}

type searchAPIResponse struct {
	Data []Series `json:"data"`
}

type seriesAPIResponse struct {
	Data Series `json:"data"`
}

type actorsAPIResponse struct {
	Data []Actor `json:"data"`
}

type episodesAPIResponse struct {
	Data []Episode `json:"data"`
}

type episodeAPIResponse struct {
	Data Episode `json:"data"`
}

type summaryAPIResponse struct {
	Data Summary `json:"data"`
}

type imagesAPIResponse struct {
	Data []Image `json:"data"`
}

type languagesAPIResponse struct {
	Data []Language `json:"data"`
}

//Updates contains an slice of show ids with their last updated date in Epoch Time
type updatesAPIResponse struct {
	Data []Update `json:"data"`
}
