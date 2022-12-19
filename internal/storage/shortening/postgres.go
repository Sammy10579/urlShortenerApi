package shortening

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
	"zipUrl/internal/model"
)

type storage struct {
	db *pgx.Conn
}

func (s *storage) Post(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
	return nil, nil
}

func (s *storage) Get(ctx context.Context, shorteningID string) (*model.Shortening, error) {
	return nil, nil
}

type postgresShortening struct {
	Identifier  string    `bson:"_id"`
	CreatedBy   string    `bson:"created_by"`
	OriginalURL string    `bson:"original_url"`
	Visits      int64     `bson:"visits"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func postgresShorteningFromModel(shortening model.Shortening) postgresShortening {
	return postgresShortening{
		Identifier:  shortening.Identifier,
		CreatedBy:   shortening.CreatedBy,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}

func modelShorteningFromPostgres(shortening postgresShortening) *model.Shortening {
	return &model.Shortening{
		Identifier:  shortening.Identifier,
		CreatedBy:   shortening.CreatedBy,
		OriginalURL: shortening.OriginalURL,
		Visits:      shortening.Visits,
		CreatedAt:   shortening.CreatedAt,
		UpdatedAt:   shortening.UpdatedAt,
	}
}
