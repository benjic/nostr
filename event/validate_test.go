package event

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Validate(t *testing.T) {
	newTestEvent := func(t *testing.T) Event {
		b := NewBuilder(1, "Hello World", testPrivateKey)

		ts, err := time.Parse(time.DateOnly, "2024-10-05")
		require.NoError(t, err)
		b.SetTimestamp(ts)

		e, err := b.Build()
		require.NoError(t, err)

		return e
	}

	t.Run("validates against Build", func(t *testing.T) {
		e := newTestEvent(t)
		assert.NoError(t, Validate(e))
	})

	t.Run("mismatched signature", func(t *testing.T) {
		e := newTestEvent(t)
		e.Signature[len(e.Signature)-1] = 0
		assert.ErrorIs(t, Validate(e), ErrInvalidSignature)
	})

	t.Run("mismatched id", func(t *testing.T) {
		e := newTestEvent(t)
		e.ID[len(e.ID)-1] = 0
		assert.ErrorIs(t, Validate(e), ErrInvalidSignature)
	})

	t.Run("corrupted values", func(t *testing.T) {
		e := newTestEvent(t)
		e.Content = "bad value"
		assert.ErrorIs(t, Validate(e), ErrInvalidSignature)
	})

}
