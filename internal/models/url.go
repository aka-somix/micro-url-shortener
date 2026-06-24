package models

type NewURLRequest struct {
	Url string `json:"url"`
}

type URL struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   int64  `json:"created_at"`
}
