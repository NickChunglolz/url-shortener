package domain

import "time"

type Url struct {
	shortenUrl  string
	longUrl     string
	createdTime time.Time
}

type UrlRepository interface {
	GetShortenUrl(longUrl string) (string, error)
	GetLongUrl(shortCode string) (string, error)
}
