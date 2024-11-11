package repository

import (
	"time"

	"github.com/NickChunglolz/url-shortener-command/domain"
	"github.com/go-pg/pg/v10"
)

type ShortenedUrlDao struct {
    tableName   struct{}  `pg:"shortened_url,alias:shortened_url"`
    Code        string    `pg:"code,pk"`
    LongUrl     string    `pg:"long_url,notnull"`
    CreatedTime time.Time `pg:"created_time,notnull,default:current_timestamp"`
}
type ShortenedUrlRepositoryImpl struct {
	db *pg.DB
}

func NewShortenedUrlRepositoryImpl(db *pg.DB) *ShortenedUrlRepositoryImpl {
	return &ShortenedUrlRepositoryImpl{
		db: db,
	}
}

func (impl *ShortenedUrlRepositoryImpl) GetShortenUrlByCode(code string) (*domain.ShortenedUrl, error) {
	dao := &ShortenedUrlDao{}
	err := impl.db.Model(dao).
		Where("code = ?", code).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	id, err := domain.ReconstituteShortenedUrlId(code)
	if err != nil {
		return nil, err
	}

	return domain.NewShortenedUrl(id, dao.LongUrl)
}

func (impl *ShortenedUrlRepositoryImpl) CreateShortenedUrl(url *domain.ShortenedUrl) error {
	dao := &ShortenedUrlDao{
		Code:        url.GetId().GetShortCode(),
		LongUrl:     url.GetLongUrl(),
		CreatedTime: url.GetCreatedTime(),
	}

	_, err := impl.db.Model(dao).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (impl *ShortenedUrlRepositoryImpl) DeleteShortenedUrl(code string) error {
	dao := &ShortenedUrlDao{Code: code}

	_, err := impl.db.Model(dao).
		Where("code = ?", code).
		Delete()

	if err != nil {
		return err
	}

	return nil
}
