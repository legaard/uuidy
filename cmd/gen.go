package cmd

import "io"

//go:generate go run github.com/matryer/moq@latest -pkg cmd_test -stub -out mocks_test.go . Writer

type Writer interface {
	io.Writer
}
