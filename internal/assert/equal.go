package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, actual T, expected T) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Values not equal: expected %+v, got %+v", actual, expected)
	}
}
