package usecase

import (
	"time"

	"github.com/NickChunglolz/url-shortener-query/domain"
)

type ShortenedUrlQuery struct {
	shortenedUrlRepository     domain.ShortenedUrlRepository
}

type GetShortenUrlResponse struct {
	ShortCode   string    `json:"shortCode"`
	LongUrl     string    `json:"longUrl"`
	CreatedTime time.Time `json:"createdTime"`
}

func NewShortenedUrlQuery(shortenedUrlRepository domain.ShortenedUrlRepository) *ShortenedUrlQuery {
	return &ShortenedUrlQuery{
		shortenedUrlRepository: shortenedUrlRepository,
	}
}

func (query *ShortenedUrlQuery) GetShortenUrlByCode(code string) (*GetShortenUrlResponse, error) {
	shortenedUrl, err := query.shortenedUrlRepository.GetShortenUrlByCode(code)
	if err != nil {
		return nil, err
	}

	return query.getShortenUrlResponseMap(shortenedUrl), nil
}

func (query *ShortenedUrlQuery) GetShortenUrlByLongUrl(longUrl string) (*GetShortenUrlResponse, error) {
	shortenedUrl, err := query.shortenedUrlRepository.GetShortenUrlByLongUrl(longUrl)
	if err != nil {
		return nil, err
	}

	return query.getShortenUrlResponseMap(shortenedUrl), nil
}

func (query *ShortenedUrlQuery) QueryShortenUrls() ([]*GetShortenUrlResponse, error) {
	shortenedUrls, err := query.shortenedUrlRepository.QueryShortenUrls()
	if err != nil {
		return nil, err
	}

	res := make([]*GetShortenUrlResponse, 0)

	for _, shortenedUrl := range shortenedUrls {
		res = append(res, query.getShortenUrlResponseMap(shortenedUrl))
	}

	return res, nil
}

func (query *ShortenedUrlQuery) getShortenUrlResponseMap(shortenedUrl *domain.ShortenedUrl) *GetShortenUrlResponse {
	return &GetShortenUrlResponse{
		ShortCode: shortenedUrl.GetId().GetShortCode(),
		LongUrl: shortenedUrl.GetLongUrl(),
		CreatedTime: shortenedUrl.GetCreatedTime(),
	}
}
