package models

type NewURLRequest struct {
	Url string `json:"url"`
}

type URL struct {
	ID          string `json:"id"           redis:"id"`
	OriginalURL string `json:"original_url" redis:"original_url"`
	ShortURL    string `json:"short_url"    redis:"short_url"`
	CreatedAt   int64  `json:"created_at"   redis:"created_at"`
}
