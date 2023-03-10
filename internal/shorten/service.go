package shorten

import (
	"context"
	"github.com/google/uuid"
	"zipUrl/internal/model"
)

type Storage interface {
	Set(ctx context.Context, shortening model.Shortening) (*model.Shortening, error)
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
	IncrementVisits(ctx context.Context, identifier string) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input model.ShortenInput) (*model.Shortening, error) {
	var (
		id         = uuid.New().ID()
		identifier = input.Identifier.OrElse(Shorten(id))
	)

	dbShortening := model.Shortening{
		Identifier:  identifier,
		OriginalURL: input.RawURL,
	}

	shortening, err := s.storage.Set(ctx, dbShortening)
	if err != nil {
		return nil, err
	}

	return shortening, nil
}
