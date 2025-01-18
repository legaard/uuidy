package assert

import (
	"testing"

	"github.com/gofrs/uuid/v5"
)

func UUIDVersion(t *testing.T, version uint32, value string) {
	t.Helper()

	parsed, err := uuid.FromString(value)
	if err != nil {
		t.Fatalf("Expected UUID: %s", err)
	}

	Equal(t, version, uint32(parsed.Version()))
}

func UUIDVersionf(t *testing.T, version uint32, value string, msg string, args ...any) {
	t.Helper()

	parsed, err := uuid.FromString(value)
	if err != nil {
		t.Logf(msg, args...)
		t.Fatalf("Expected UUID: %s", err)
	}

	Equal(t, version, uint32(parsed.Version()))
}

func UUID(t *testing.T, value string) {
	t.Helper()

	_, err := uuid.FromString(value)
	if err != nil {
		t.Fatalf("Expected UUID: %s", err)
	}
}

func UUIDf(t *testing.T, value string, msg string, args ...any) {
	t.Helper()

	_, err := uuid.FromString(value)
	if err != nil {
		t.Logf(msg, args...)
		t.Fatalf("Expected UUID: %s", err)
	}
}
