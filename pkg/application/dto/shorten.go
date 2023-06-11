package dto

type ShortenDTO struct {
	URL string `json:"url"`
}

type ShortenedDTO struct {
	ShortenedURL string `json:"shortenedURL"`
}
