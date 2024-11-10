package usecase

import (
	"fmt"
	"time"

	"github.com/NickChunglolz/url-shortener-command/domain"
)

type CreateShortenUrlRequest struct {
	OriginalURL string `json:"originalUrl"`
}

type CreateShortenUrlResponse struct {
	ShortCode   string    `json:"shortCode"`
	LongUrl     string    `json:"longUrl"`
	CreatedTime time.Time `json:"createdTime"`
}

type ShortenedUrlCommand struct {
	urlRepository     domain.UrlRepository
	counterRepository domain.CounterRepository
}

func NewShortenedUrlCommand() *ShortenedUrlCommand {
	return &ShortenedUrlCommand{}
}

func (command *ShortenedUrlCommand) CreateShortenUrl(request *CreateShortenUrlRequest) (*CreateShortenUrlResponse, error) {
	if request.OriginalURL == "" {
		return nil, fmt.Errorf("original URL cannot be empty")
	}

	// Get next counter value
	counter, err := command.counterRepository.GetNextCounter()
	if err != nil {
		return nil, fmt.Errorf("failed to generate counter: %w", err)
	}

	id, err := domain.NewShortenedUrlId(counter)
	if err != nil {
		return nil, err
	}

	shortenedUrl, err := domain.NewShortenedUrl(id, request.OriginalURL)
	if err != nil {
		return nil, err
	}

	err = command.urlRepository.CreateShortenedUrl(shortenedUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create shortened URL: %w", err)
	}

	shortenedUrl, err = command.urlRepository.GetShortenUrlByCode(id.GetShortCode())
	if err != nil {
		return nil, err
	}

	return &CreateShortenUrlResponse{
		ShortCode:   shortenedUrl.GetId().GetShortCode(),
		LongUrl:     shortenedUrl.GetLongUrl(),
		CreatedTime: shortenedUrl.GetCreatedTime(),
	}, nil
}
