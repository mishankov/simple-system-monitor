package testutils

import "testing"

func Assert(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %v want %v", got, want)
	}
}
