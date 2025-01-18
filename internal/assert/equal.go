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

func Equalf[T any](t *testing.T, expected T, actual T, msg string, args ...any) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Logf(msg, args...)
		t.Fatalf("Expected value to be equal: expected %+v, got %+v", expected, actual)
	}
}

func NotEqual[T any](t *testing.T, expected T, actual T) {
	t.Helper()

	if reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected value to be equal: expected %+v, got %+v", expected, actual)
	}
}

func NotEqualf[T any](t *testing.T, expected T, actual T, msg string, args ...any) {
	t.Helper()

	if reflect.DeepEqual(actual, expected) {
		t.Logf(msg, args...)
		t.Fatalf("Expected value to be not equal: expected %+v, got %+v", expected, actual)
	}
}
