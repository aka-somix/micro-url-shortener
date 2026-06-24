package repositories

import (
	"aka-somix/micro-url-shortener/internal/models"
	"context"
	"encoding/json"
	"os"
	"sort"
	"sync"
)

var _ URLRepository = (*JsonUrlRepository)(nil)

type JsonUrlRepository struct {
	mu       sync.RWMutex
	filePath string
	urls     map[string]models.URL
}

func NewJsonUrlRepository(filePath string) (*JsonUrlRepository, error) {
	repo := &JsonUrlRepository{
		filePath: filePath,
		urls:     make(map[string]models.URL),
	}
	if err := repo.load(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *JsonUrlRepository) load() error {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}
	var products []models.URL
	if err := json.Unmarshal(data, &products); err != nil {
		return err
	}
	for _, p := range products {
		r.urls[p.ID] = p
	}
	return nil
}
func (r *JsonUrlRepository) save() error {

	productList := make([]models.URL, 0, len(r.urls))

	for _, v := range r.urls {
		productList = append(productList, v)
	}

	data, err := json.MarshalIndent(productList, "", "  ")

	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *JsonUrlRepository) Create(_ context.Context, url *models.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.urls[url.ID] = *url
	return r.save()
}

func (r *JsonUrlRepository) GetById(_ context.Context, id string) (*models.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	url, exists := r.urls[id]

	if !exists {
		return nil, models.ErrNotFound
	}
	return &url, nil
}

func (r *JsonUrlRepository) GetAll(_ context.Context) ([]models.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	urls := make([]models.URL, 0, len(r.urls))
	for _, url := range r.urls {
		urls = append(urls, url)
	}
	return urls, nil
}

func (r *JsonUrlRepository) GetLatest(_ context.Context, n int) ([]models.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	all := make([]models.URL, 0, len(r.urls))
	for _, url := range r.urls {
		all = append(all, url)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt > all[j].CreatedAt
	})
	if n > len(all) {
		n = len(all)
	}
	return all[:n], nil
}

func (r *JsonUrlRepository) DeleteById(_ context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.urls, id)
	return r.save()
}
