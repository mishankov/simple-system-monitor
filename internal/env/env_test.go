package env_test

import (
	"os"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/env"
	"github.com/mishankov/simple-system-monitor/internal/testutils"
)

func TestEnv(t *testing.T) {
	env := env.New()

	t.Run("test string", func(t *testing.T) {
		os.Setenv("MY_VAR", "value")

		val := env.GetStringOrDefault("MY_VAR", "other value")
		testutils.Assert(t, val, "value")
	})

	t.Run("test string default", func(t *testing.T) {
		val := env.GetStringOrDefault("EMPTY_VAR", "other value")
		testutils.Assert(t, val, "other value")
	})

	t.Run("test int", func(t *testing.T) {
		os.Setenv("MY_VAR", "1")

		val, err := env.GetIntOrDefault("MY_VAR", 3)
		testutils.AssertError(t, err)
		testutils.Assert(t, val, 1)
	})

	t.Run("test int default", func(t *testing.T) {
		val, err := env.GetIntOrDefault("EMPTY_VAR", 3)
		testutils.AssertError(t, err)
		testutils.Assert(t, val, 3)
	})

	t.Run("test int error", func(t *testing.T) {
		os.Setenv("MY_VAR", "not int")

		val, err := env.GetIntOrDefault("MY_VAR", 3)
		if err == nil {
			t.Error("Expected error, got nil")
		}
		testutils.Assert(t, val, 0)
	})
}
