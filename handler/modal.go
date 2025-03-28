package handler

type ShortenUrlRequest struct {
	Url       string `json:"url"`
	QueryName string `json:"name,omitempty"`
}

type ShortnerUrlResponse struct {
	Url string `json:"url"`
}
