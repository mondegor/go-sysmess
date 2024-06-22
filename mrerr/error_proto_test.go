package mrerr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrerr"
)

func TestNewProto(t *testing.T) {
	t.Parallel()

	got := mrerr.NewProto("test-code", mrerr.ErrorKindSystem, "test-message")
	assert.Equal(t, "test-code", got.Code())
	assert.Equal(t, mrerr.ErrorKindSystem, got.Kind())
	assert.ErrorContains(t, got, "test-message")
}

func TestNewProtoWithExtra(t *testing.T) {
	t.Parallel()

	got := mrerr.NewProtoWithExtra("test-code", mrerr.ErrorKindSystem, "test-message", mrerr.ProtoExtra{})
	assert.Equal(t, "test-code", got.Code())
	assert.Equal(t, mrerr.ErrorKindSystem, got.Kind())
	assert.ErrorContains(t, got, "test-message")
}

func TestProtoAppError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		message   string
		argsNames []string
		args      []any
		want      string
	}{
		{
			name:    "test1",
			message: "my-message",
			want:    "my-message",
		},
		{
			name:    "test2",
			message: "my-message {{ .key1 }} - {{ .key2 }}",
			want:    "my-message missed-arg1 - missed-arg2",
		},
		{
			name:    "test3",
			message: "my-message {{ .key1 }} - {{ .key3 }} - {{ .key2 }}",
			want:    "my-message missed-arg1 - missed-arg2 - missed-arg3",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := mrerr.NewProto("", mrerr.ErrorKindInternal, tt.message)
			got := e.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}
