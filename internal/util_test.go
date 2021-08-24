package internal_test

import (
	"context"
	"testing"

	"github.com/berquerant/ga/internal"
	"github.com/stretchr/testify/assert"
)

func TestIsDone(t *testing.T) {
	t.Run("done", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.TODO())
		cancel()
		assert.True(t, internal.IsDone(ctx))
	})
	t.Run("not done", func(t *testing.T) {
		assert.False(t, internal.IsDone(context.TODO()))
	})
}
