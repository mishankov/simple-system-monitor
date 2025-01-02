package testutils

import "testing"

func Assert(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %v want %v", got, want)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Error was not expected:", err)
	}
}
