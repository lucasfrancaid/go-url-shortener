package domain

type Shortener struct {
	HashedURL string
	URL       string
}

type ShortenerStats struct {
	HashedURL string
	Counter   int64
}
