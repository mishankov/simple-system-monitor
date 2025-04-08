package env_test

import (
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/simple-system-monitor/internal/env"
)

func TestEnv(t *testing.T) {
	env := env.New()

	t.Run("test string", func(t *testing.T) {
		t.Setenv("MY_VAR", "value")

		val := env.GetStringOrDefault("MY_VAR", "other value")
		assert.Equal(t, val, "value")
	})

	t.Run("test string default", func(t *testing.T) {
		val := env.GetStringOrDefault("EMPTY_VAR", "other value")
		assert.Equal(t, val, "other value")
	})

	t.Run("test int", func(t *testing.T) {
		t.Setenv("MY_VAR", "1")

		val, err := env.GetIntOrDefault("MY_VAR", 3)
		assert.NoError(t, err)
		assert.Equal(t, val, 1)
	})

	t.Run("test int default", func(t *testing.T) {
		val, err := env.GetIntOrDefault("EMPTY_VAR", 3)
		assert.NoError(t, err)
		assert.Equal(t, val, 3)
	})

	t.Run("test int error", func(t *testing.T) {
		t.Setenv("MY_VAR", "not int")

		val, err := env.GetIntOrDefault("MY_VAR", 3)
		if err == nil {
			t.Error("Expected error, got nil")
		}
		assert.Equal(t, val, 0)
	})
}
