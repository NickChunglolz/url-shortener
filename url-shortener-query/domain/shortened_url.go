package domain

import (
	"errors"
	"time"
)

type ShortenedUrlRepository interface {
	GetShortenUrlByCode(code string) (*ShortenedUrl, error)
	GetShortenUrlByLongUrl(longUrl string) (*ShortenedUrl, error)
	QueryShortenUrls() ([]*ShortenedUrl, error)
}

type ShortenedUrlId struct {
	shortCode string
}

func ReconstituteShortenedUrlId(shortCode string) (*ShortenedUrlId, error) {
	if shortCode == "" {
		return nil, errors.New("shortCode cannot be empty")
	}

	return &ShortenedUrlId{
		shortCode: shortCode,
	}, nil
}

func (id *ShortenedUrlId) GetShortCode() string {
	return id.shortCode
}

type ShortenedUrl struct {
	id          *ShortenedUrlId
	longUrl     string
	createdTime time.Time
}

func ReconstituteShortenedUrl(shortCode string, longUrl string, createdTime time.Time) (*ShortenedUrl, error) {
	if shortCode == "" {
		return nil, errors.New("shortCode cannot be empty")
	}

	id, _ := ReconstituteShortenedUrlId(shortCode)

	return &ShortenedUrl{
		id: id,
		longUrl: longUrl,
		createdTime: createdTime,
	}, nil
}

func (s *ShortenedUrl) GetId() *ShortenedUrlId {
	return s.id
}

func (s *ShortenedUrl) GetLongUrl() string {
	return s.longUrl
}

func (s *ShortenedUrl) GetCreatedTime() time.Time {
	return s.createdTime
}