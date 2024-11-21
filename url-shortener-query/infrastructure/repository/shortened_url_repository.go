package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NickChunglolz/url-shortener-query/domain"
	"github.com/NickChunglolz/url-shortener-query/main/utils"
	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
)

type ShortenedUrlDao struct {
    tableName   struct{}  `pg:"shortened_url,alias:shortened_url"`
    Code        string    `pg:"code,pk"`
    LongUrl     string    `pg:"long_url,notnull"`
    CreatedTime time.Time `pg:"created_time,notnull,default:current_timestamp"`
}

type ShortenedUrlRepositoryImpl struct {
	config *utils.Config
	db *pg.DB
	cacheDb *redis.Client
	ctx     context.Context
}

func NewShortenedUrlRepositoryImpl(config *utils.Config, db *pg.DB, cacheDb *redis.Client) *ShortenedUrlRepositoryImpl {
	return &ShortenedUrlRepositoryImpl{
		config: config,
		db: db,
		cacheDb: cacheDb,
		ctx:   context.Background(),
	}
}

func (impl *ShortenedUrlRepositoryImpl) GetShortenUrlByCode(code string) (*domain.ShortenedUrl, error) {
	// Try to get from cache first
	var dao ShortenedUrlDao
    cacheKey := impl.config.Cache.KeyPrefix + code
    cachedData, err := impl.cacheDb.Get(impl.ctx, cacheKey).Result()
    if err == nil {
        if err := json.Unmarshal([]byte(cachedData), &dao); err == nil {
            return domain.ReconstituteShortenedUrl(dao.Code, dao.LongUrl, dao.CreatedTime)
        }
    }

	// If not in cache, get from database
	err = impl.db.Model(&dao).
		Where("code = ?", code).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Store in cache
    if jsonData, err := json.Marshal(dao); err == nil {
        impl.cacheDb.Set(impl.ctx, cacheKey, jsonData, time.Duration(impl.config.Cache.Ttl) * time.Hour)
    }

	return domain.ReconstituteShortenedUrl(dao.Code, dao.LongUrl, dao.CreatedTime)
}

func (impl *ShortenedUrlRepositoryImpl) GetShortenUrlByLongUrl(longUrl string) (*domain.ShortenedUrl, error) {
	var dao ShortenedUrlDao
	err := impl.db.Model(&dao).
		Where("long_url = ?", longUrl).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return domain.ReconstituteShortenedUrl(dao.Code, dao.LongUrl, dao.CreatedTime)
}

func (impl *ShortenedUrlRepositoryImpl) QueryShortenUrls() ([]*domain.ShortenedUrl, error) {
	var daos []ShortenedUrlDao
	err := impl.db.Model(&daos).
		Select()

	if err != nil {
		return nil, err
	}
		
	var results []*domain.ShortenedUrl
	for _, dao := range daos {
		shortenedUrl, err := domain.ReconstituteShortenedUrl(dao.Code, dao.LongUrl, dao.CreatedTime)
		if err != nil {
			return nil, err
		}
		results = append(results, shortenedUrl)
	}
		
	return results, nil
}