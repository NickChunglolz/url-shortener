package domain

import (
	"errors"
	"time"
)

type ShortenedUrlRepository interface {
	GetShortenUrlById(id *ShortenedUrlId) (*ShortenedUrl, error)
	CreateShortenedUrl(url *ShortenedUrl) error
	DeleteShortenedUrl(code string) error
}

type ShortenedUrlId struct {
	shortCode string
}

func NewShortenedUrlId(counter uint64) *ShortenedUrlId {
	return &ShortenedUrlId{
		shortCode: generateShortCode(counter),
	}
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

func NewShortenedUrl(counter uint64, longUrl string) (*ShortenedUrl, error) {
	if longUrl == "" {
		return nil, errors.New("long url cannot be empty")
	}

	return &ShortenedUrl{
		id:          NewShortenedUrlId(counter),
		longUrl:     longUrl,
		createdTime: time.Now(),
	}, nil
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

func generateShortCode(counter uint64) string {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base := uint64(len(alphabet))

	var shortCode string
	n := counter

	for n > 0 {
		shortCode = string(alphabet[n%base]) + shortCode
		n = n / base
	}

	for len(shortCode) < 6 {
		shortCode = "0" + shortCode
	}

	return shortCode
}
