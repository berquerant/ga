package ga_test

import (
	"errors"
	"testing"

	"github.com/berquerant/ga"
	"github.com/stretchr/testify/assert"
)

func TestGAErrorTrace(t *testing.T) {
	t.Run("flat", func(t *testing.T) {
		err := ga.NewGAError(nil, ga.UnknownError, "msg")
		assert.Equal(t, "[UnknownError] msg <nil>", err.Trace())
	})
	t.Run("wrapped", func(t *testing.T) {
		err := ga.NewGAError(
			ga.NewGAError(errors.New("internal"), ga.UnknownError, "wrapped"),
			ga.UnknownError,
			"msg",
		)
		assert.Equal(t, `[UnknownError] msg
[UnknownError] wrapped internal`, err.Trace())
	})
}

func TestGAErrorError(t *testing.T) {
	t.Run("flat", func(t *testing.T) {
		err := ga.NewGAError(nil, ga.UnknownError, "msg")
		assert.Equal(t, "msg", err.Error())
	})
	t.Run("wrapped", func(t *testing.T) {
		err := ga.NewGAError(
			ga.NewGAError(nil, ga.UnknownError, "wrapped"),
			ga.UnknownError,
			"msg",
		)
		assert.Equal(t, "wrapped", err.Error())
	})
}
