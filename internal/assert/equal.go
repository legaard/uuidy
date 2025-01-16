package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, expected T, actual T) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected value to be equal: expected %+v, got %+v", expected, actual)
	}
}
