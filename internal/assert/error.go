package assert

import "testing"

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Expected no error: got %v", err)
	}
}

func NoErrorf(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err != nil {
		t.Logf(msg, args...)
		t.Fatalf("Expected no error: got %v", err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatalf("Expected error: got <nil>")
	}
}

func Errorf(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err == nil {
		t.Logf(msg, args...)
		t.Fatalf("Expected error: got <nil>")
	}
}
