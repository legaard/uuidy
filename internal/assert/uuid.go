package assert

import (
	"testing"

	"github.com/gofrs/uuid/v5"
)

func UUIDVersion(t *testing.T, actual string, version uint32) {
	t.Helper()

	parsed, err := uuid.FromString(actual)
	if err != nil {
		t.Fatalf("Value not a UUID: %s", err)
	}

	Equal(t, version, uint32(parsed.Version()))
}

func UUID(t *testing.T, actual string) {
	t.Helper()

	_, err := uuid.FromString(actual)
	if err != nil {
		t.Fatalf("Value not a UUID: %s", err)
	}
}
