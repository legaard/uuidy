package clipboard

import (
	"fmt"
	"os/exec"
)

const (
	copyCmdArgs = "pbcopy"
)

type Writer struct{}

func NewWriter() Writer {
	return Writer{}
}

func (w Writer) Write(p []byte) (n int, err error) {
	copyCmd := exec.Command(copyCmdArgs)
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return 0, fmt.Errorf("getting command writer: %w", err)
	}

	err = copyCmd.Start()
	if err != nil {
		return 0, fmt.Errorf("starting %s command: %w", copyCmdArgs, err)
	}

	_, err = in.Write(p)
	if err != nil {
		return 0, fmt.Errorf("writing data: %w", err)
	}

	err = in.Close()
	if err != nil {
		return 0, fmt.Errorf("closing writer: %w", err)
	}

	err = copyCmd.Wait()
	if err != nil {
		return 0, fmt.Errorf("waiting for %s command: %w", copyCmdArgs, err)
	}

	return len(p), nil
}
