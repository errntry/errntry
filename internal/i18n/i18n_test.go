package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	t.Run("error", func(t *testing.T) {
		err := Error("some error occurred")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error occurred")
	})
}
