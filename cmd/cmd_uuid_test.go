package cmd_test

import (
	"fmt"
	"github.com/legaard/uuidy/cmd"
	"github.com/legaard/uuidy/internal/assert"
	"strings"
	"testing"
	"time"

	"github.com/gofrs/uuid/v5"
)

func TestV1Cmd(t *testing.T) {
	t.Run(`use is "v1"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V1Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v1", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V1Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, len(writerMock.WriteCalls()), 1)

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 1, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V1Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 1, actual)
		}
	})
}

func TestV3Cmd(t *testing.T) {
	t.Run(`use is "v3 [value]"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V3Cmd(uuid.NamespaceDNS)
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v3 [value]", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 3, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 3, actual)
		}
	})

	t.Run("return error on invalid namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = "invalid"
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = uuid.Must(uuid.NewV4()).String()
			sut        = cmd.V3Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 3, string(actual))
	})
}

func TestV4Cmd(t *testing.T) {
	t.Run(`use is "v4"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V4Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v4", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V4Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 4, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V4Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 4, actual)
		}
	})
}

func TestV5Cmd(t *testing.T) {
	t.Run(`use is "v5 [value]"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V5Cmd(uuid.NamespaceDNS)
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v5 [value]", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 5, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 5, actual)
		}
	})

	t.Run("return error on invalid namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = "invalid"
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with namespace", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			ns         = uuid.Must(uuid.NewV4()).String()
			sut        = cmd.V5Cmd(uuid.NamespaceDNS)
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNamespace, ns)

		// act
		err := sut.RunE(sut, []string{"testing"})

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 5, string(actual))
	})
}

func TestV6Cmd(t *testing.T) {
	t.Run(`use is "v6"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V6Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v6", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V6Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 6, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V6Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 6, actual)
		}
	})
}

func TestV7Cmd(t *testing.T) {
	t.Run(`use is "v7"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.V7Cmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "v7", actual)
	})

	t.Run("generate UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 7, string(actual))
	})

	t.Run("generate multiple UUIDs", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			number     = 10
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagNumber, fmt.Sprintf("%d", number))

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, number, len(writerMock.WriteCalls()))
		for _, call := range writerMock.WriteCalls() {
			actual := strings.ReplaceAll(string(call.P), "\n", "") // remove new lines
			assert.UUIDVersion(t, 7, actual)
		}
	})

	t.Run("return error on invalid epoch", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			epoch      = "invalid"
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagEpoch, epoch)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.Error(t, err)
	})

	t.Run("generate uuid with epoch", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			epoch      = time.Now().Format(time.RFC3339Nano)
			sut        = cmd.V7Cmd()
		)
		sut.SetOut(writerMock)
		_ = sut.Flags().Set(cmd.FlagEpoch, epoch)

		// act
		err := sut.RunE(sut, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.UUIDVersion(t, 7, string(actual))
	})
}

func TestNullCmd(t *testing.T) {
	t.Run(`use is "null"`, func(t *testing.T) {
		// arrange
		var (
			sut = cmd.NullCmd()
		)

		// act
		actual := sut.Use

		// assert
		assert.Equal(t, "null", actual)
	})

	t.Run("generate null UUID", func(t *testing.T) {
		// arrange
		var (
			writerMock = &WriterMock{}
			sut        = cmd.NullCmd()
		)
		sut.SetOut(writerMock)

		// act
		sut.Run(sut, nil)

		// assert
		assert.Equal(t, 1, len(writerMock.WriteCalls()))

		actual := writerMock.WriteCalls()[0].P
		assert.Equal(t, "00000000-0000-0000-0000-000000000000", string(actual))
	})
}
