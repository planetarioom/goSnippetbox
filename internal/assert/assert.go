package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, want T) {
	t.Helper()

	if actual != want {
		t.Errorf("got %v; want %v", actual, want)
	}
}

func StringContains(t *testing.T, actual, expected string) {
	t.Helper()
	
	if !strings.Contains(actual, expected) {
		t.Errorf("got %s; expected %s", actual, expected)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got %v, expected this shit: nil", actual)
	}
}