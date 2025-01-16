package assert

import "testing"

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatalf("Got no error")
	}
}
