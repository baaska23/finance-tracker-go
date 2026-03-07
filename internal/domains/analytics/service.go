package analytics

import (
	"net/http"
	"time"
)

type Service struct {
	repo Repository
	httpClient *http.Client
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Service) SendSomething(payload string) error {
	return nil
}