package services

import (
	"aka-somix/micro-url-shortener/configs"
	"aka-somix/micro-url-shortener/internal/models"
	"aka-somix/micro-url-shortener/internal/repositories"
)

type UrlShortenService struct {
	urlRepository repositories.URLRepository
}

func NewUrlShortenService(urlRepo repositories.URLRepository) *UrlShortenService {
	return &UrlShortenService{
		urlRepository: urlRepo,
	}
}

func (s *UrlShortenService) ShortenURL(originalURL string) (string, error) {
	return configs.BaseURL + "/url/" + "ex1a2b", nil
}

func (s *UrlShortenService) GetUrlFromShort(shortURL string) (models.URL, error) {
	return models.URL{ShortURL: shortURL, OriginalURL: "https://example.com"}, nil
}

func (s *UrlShortenService) GetAllURLs() ([]models.URL, error) {
	return []models.URL{{ShortURL: "ex1a2b", OriginalURL: "https://example.com"}}, nil
}
