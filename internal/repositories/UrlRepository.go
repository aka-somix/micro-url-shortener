package repositories

import (
	"aka-somix/micro-url-shortener/internal/models"
	"context"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	GetById(ctx context.Context, id string) (*models.URL, error)
	GetAll(ctx context.Context) ([]models.URL, error)
	DeleteById(ctx context.Context, id string) error
}
