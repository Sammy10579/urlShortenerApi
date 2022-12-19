package shorten

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Run("return an alphanumetric short identifier", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}

		testCases := []testCase{
			{
				id:       1024,
				expected: "Mv",
			},
			{
				id:       0,
				expected: "",
			},
		}

		for _, tc := range testCases {
			actual := Shorten(tc.id)
			assert.Equal(t, tc.expected, actual)
		}
	})

	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, "Mv", Shorten(1024))
		}
	})
}
