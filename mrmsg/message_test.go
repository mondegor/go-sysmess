package mrmsg_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrmsg"
)

func Test_render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		args    []mrmsg.NamedArg
		want    string
		wantErr bool
	}{
		{
			name:    "empty message",
			message: "",
			args:    []mrmsg.NamedArg{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "message with arg",
			message: "test message {{ .arg1 }}",
			args:    []mrmsg.NamedArg{{"arg1", "value1"}},
			want:    "test message value1",
			wantErr: false,
		},
		{
			name:    "message with arg without value",
			message: "test message {{ .arg1 }}",
			args:    []mrmsg.NamedArg{},
			want:    "test message ",
			wantErr: false,
		},
		{
			name:    "message with a few args",
			message: "{{ .arg2 }} test {{ .arg3 }} message {{ .arg1 }}",
			args:    []mrmsg.NamedArg{{"arg1", "value1"}, {"arg2", "value2"}, {"arg3", "value3"}},
			want:    "value2 test value3 message value1",
			wantErr: false,
		},
		{
			name:    "message with error: 'template: :1: unclosed action'",
			message: "test message {{ .arg1",
			args:    []mrmsg.NamedArg{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "message is missing arg1",
			message: "test message {{ .arg2 }}",
			args:    []mrmsg.NamedArg{{"arg1", "value1"}, {"arg2", "value2"}},
			want:    "test message value2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := mrmsg.RenderWithNamedArgs(tt.message, tt.args)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseArgsNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		want    []string
	}{
		{
			name:    "empty message",
			message: "",
			want:    nil,
		},
		{
			name:    "message without args",
			message: "test message",
			want:    nil,
		},
		{
			name:    "message with bad arg 1",
			message: "test message {{.arg1}}",
			want:    nil,
		},
		{
			name:    "message with bad arg 2",
			message: "test message {{ arg1 }}",
			want:    nil,
		},
		{
			name:    "message with one arg",
			message: "test message {{ .arg1 }}",
			want:    []string{"arg1"},
		},
		{
			name:    "message with two args",
			message: "test message {{ .arg1 }}, {{ .arg2 }}",
			want:    []string{"arg1", "arg2"},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := mrmsg.ParseArgsNames(tt.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgsNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
