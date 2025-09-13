package mrmsg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrmsg"
)

func Test_NewMessageReplacer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		args    []any
		want    string
	}{
		{
			name:    "empty message",
			message: "",
			args:    nil,
			want:    "",
		},
		{
			name:    "message with arg",
			message: "test message {Arg1}",
			args:    []any{"value1"},
			want:    "test message value1",
		},
		{
			name:    "message with arg without value",
			message: "test message {Arg1}",
			args:    nil,
			want:    "test message !MISSINGARG",
		},
		{
			name:    "message with a few placeholders",
			message: "{Arg1} test {Arg2} message {Arg3}",
			args:    []any{"value1", "value2", "value3"},
			want:    "value1 test value2 message value3",
		},
		{
			name:    "message with missed }",
			message: "message {Arg1",
			args:    []any{"value1"},
			want:    "message {Arg1",
		},
		{
			name:    "message with missed } 2",
			message: "message {Arg1 {Arg2}",
			args:    []any{"value1", "value2"},
			want:    "message {Arg1 value1",
		},
		{
			name:    "message with missed {",
			message: "message {Arg1} Arg2}",
			args:    []any{"value1", "value2"},
			want:    "message value1 Arg2}",
		},
		{
			name:    "message is missing arg1",
			message: "test message {Arg2}",
			args:    []any{"value1", "value2"},
			want:    "test message value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			replacer := mrmsg.NewMessageReplacer("{", "}", tt.message)

			got, err := replacer.Replace(tt.args)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
