package mrmsg

import (
	"reflect"
	"testing"
)

func Test_render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message string
		args    []NamedArg
		want    string
		wantErr bool
	}{
		{
			name:    "empty message",
			message: "",
			args:    []NamedArg{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "message with arg",
			message: "test message {{ .arg1 }}",
			args:    []NamedArg{{"arg1", "value1"}},
			want:    "test message value1",
			wantErr: false,
		},
		{
			name:    "message with arg without value",
			message: "test message {{ .arg1 }}",
			args:    []NamedArg{},
			want:    "test message ",
			wantErr: false,
		},
		{
			name:    "message with a few args",
			message: "{{ .arg2 }} test {{ .arg3 }} message {{ .arg1 }}",
			args:    []NamedArg{{"arg1", "value1"}, {"arg2", "value2"}, {"arg3", "value3"}},
			want:    "value2 test value3 message value1",
			wantErr: false,
		},
		{
			name:    "message with error: 'template: :1: unclosed action'",
			message: "test message {{ .arg1",
			args:    []NamedArg{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "message is missing arg1",
			message: "test message {{ .arg2 }}",
			args:    []NamedArg{{"arg1", "value1"}, {"arg2", "value2"}},
			want:    "test message value2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := render(tt.message, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("render() got = %v, want %v", got, tt.want)
			}
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

			if got := ParseArgsNames(tt.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgsNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
