package repositories

import (
	"aka-somix/micro-url-shortener/configs"
	"aka-somix/micro-url-shortener/internal/models"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	urlHashPrefix = "url:"
	urlsSortedSet = "urls:by_time"
)

var _ URLRepository = (*RedisUrlRepository)(nil)

type RedisUrlRepository struct {
	client *redis.Client
}

func NewRedisUrlRepository() (*RedisUrlRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.RedisURL,
		Password: configs.RedisPassword,
		DB:       0,
		Protocol: 2,
	})

	return &RedisUrlRepository{client}, nil
}

type urlHash struct {
	ID          string `redis:"id"`
	OriginalURL string `redis:"original_url"`
	CreatedAt   int64  `redis:"created_at"`
}

func (r *RedisUrlRepository) Create(ctx context.Context, url *models.URL) error {
	count, err := r.client.ZCard(ctx, urlsSortedSet).Result()
	if err != nil {
		return err
	}
	if count >= int64(configs.MaxURLs) {
		return models.ErrStorageFull
	}

	ttl := time.Duration(configs.URLTTLMinutes) * time.Minute

	pipe := r.client.TxPipeline()
	pipe.HSet(ctx, urlHashPrefix+url.ID, urlHash{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		CreatedAt:   url.CreatedAt,
	})
	pipe.Expire(ctx, urlHashPrefix+url.ID, ttl)
	pipe.ZAdd(ctx, urlsSortedSet, redis.Z{
		Score:  float64(url.CreatedAt),
		Member: url.ID,
	})

	_, err = pipe.Exec(ctx)
	return err
}

func (r *RedisUrlRepository) GetById(ctx context.Context, id string) (*models.URL, error) {
	vals, err := r.client.HGetAll(ctx, urlHashPrefix+id).Result()
	if err != nil {
		return nil, err
	}
	if len(vals) == 0 {
		return nil, nil
	}
	return assembleUrl(vals)
}

func (r *RedisUrlRepository) GetAll(ctx context.Context) ([]models.URL, error) {
	keys, err := r.client.Keys(ctx, urlHashPrefix+"*").Result()
	if err != nil {
		return nil, err
	}

	urls := make([]models.URL, 0, len(keys))
	for _, key := range keys {
		vals, err := r.client.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		if len(vals) == 0 {
			id := key[len(urlHashPrefix):]
			r.client.ZRem(ctx, urlsSortedSet, id)
			continue
		}
		url, err := assembleUrl(vals)
		if err != nil {
			return nil, err
		}
		urls = append(urls, *url)
	}
	return urls, nil
}

func (r *RedisUrlRepository) GetLatest(ctx context.Context, n int) ([]models.URL, error) {
	// Get URL Ids from Sorted Set
	ids, err := r.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   urlsSortedSet,
		Start: 0,
		Stop:  int64(n - 1),
		Rev:   true,
	}).Result()
	if err != nil {
		return nil, err
	}

	// Get each URL
	urls := make([]models.URL, 0, len(ids))
	for _, id := range ids {
		vals, err := r.client.HGetAll(ctx, urlHashPrefix+id).Result()
		if err != nil {
			return nil, err
		}
		if len(vals) == 0 {
			r.client.ZRem(ctx, urlsSortedSet, id)
			continue
		}
		url, err := assembleUrl(vals)
		if err != nil {
			return nil, err
		}
		urls = append(urls, *url)
	}
	return urls, nil
}

func (r *RedisUrlRepository) DeleteById(ctx context.Context, id string) error {
	pipe := r.client.TxPipeline()
	pipe.Del(ctx, urlHashPrefix+id)
	pipe.ZRem(ctx, urlsSortedSet, id)
	_, err := pipe.Exec(ctx)
	return err
}

func assembleUrl(vals map[string]string) (*models.URL, error) {
	createdAt, err := strconv.ParseInt(vals["created_at"], 10, 64)
	if err != nil {
		return nil, err
	}
	return &models.URL{
		ID:          vals["id"],
		OriginalURL: vals["original_url"],
		ShortURL:    configs.BaseURL + "/url/" + vals["id"],
		CreatedAt:   createdAt,
	}, nil
}
