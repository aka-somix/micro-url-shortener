package services

import (
	"aka-somix/micro-url-shortener/configs"
	"aka-somix/micro-url-shortener/internal/models"
	"aka-somix/micro-url-shortener/internal/repositories"
	"context"

	"github.com/google/uuid"
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
	newShort := uuid.New().String()
	shortUrl := configs.BaseURL + "/url/" + newShort

	s.urlRepository.Create(
		context.TODO(),
		&models.URL{
			ID:          newShort,
			OriginalURL: originalURL,
			ShortURL:    shortUrl,
		})
	return newShort, nil
}

func (s *UrlShortenService) GetUrlFromShort(shortURL string) (*models.URL, error) {
	foundUrl, err := s.urlRepository.GetById(context.TODO(), shortURL)

	if err != nil {
		return nil, err
	}

	return foundUrl, nil
}

func (s *UrlShortenService) GetAllURLs() ([]models.URL, error) {
	allUrls, err := s.urlRepository.GetAll(context.TODO())

	if err != nil {
		return nil, err
	}

	return allUrls, nil
}

func (s *UrlShortenService) GetLatestURLs(n int) ([]models.URL, error) {
	return s.urlRepository.GetLatest(context.TODO(), n)
}
