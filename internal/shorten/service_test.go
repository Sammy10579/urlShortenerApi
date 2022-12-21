package shorten

import (
	"context"
	. "github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"zipUrl/internal/model"
	"zipUrl/internal/storage/shortening"
)

func TestService_Shorten(t *testing.T) {
	t.Run("generates shortening for a given URL", func(t *testing.T) {
		var (
			svc   = NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://google.com"}
		)

		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		require.NotEmpty(t, shortening.Identifier)
		assert.Equal(t, input.RawURL, shortening.OriginalURL)
		assert.NotZero(t, shortening.CreatedAt)
	})

	t.Run("uses custom identifier if provided", func(t *testing.T) {
		const identifier = "google"

		var (
			svc   = NewService(shortening.NewInMemory())
			input = model.ShortenInput{
				RawURL:     "https://www.google.com",
				Identifier: Some(identifier),
			}
		)

		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		assert.Equal(t, identifier, shortening.Identifier)
		assert.Equal(t, "https://www.google.com", shortening.OriginalURL)
		assert.NotZero(t, shortening.CreatedAt)
	})

	t.Run("returns error if identifier is already taken", func(t *testing.T) {
		const identifier = "google"

		var (
			svc   = NewService(shortening.NewInMemory())
			input = model.ShortenInput{
				RawURL:     "https://www.google.com",
				Identifier: Some(identifier),
			}
		)

		_, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		_, err = svc.Shorten(context.Background(), input)
		assert.ErrorIs(t, err, model.ErrIdentifierExists)
	})
}
