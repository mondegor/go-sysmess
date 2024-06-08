package mrerr

import (
	"errors"
	"testing"

	"github.com/mondegor/go-sysmess/mrmsg"
	"github.com/stretchr/testify/assert"
)

func TestAppError_WithAttr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		paramName  string
		paramValue any
		want       []mrmsg.NamedArg
	}{
		{
			name:       "test1",
			paramName:  "",
			paramValue: nil,
			want: []mrmsg.NamedArg{
				{
					Name:  attrNameByDefault,
					Value: nil,
				},
			},
		},
		{
			name:       "test2",
			paramName:  "test-name",
			paramValue: "test-value",
			want: []mrmsg.NamedArg{
				{
					Name:  "test-name",
					Value: "test-value",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &AppError{}

			got := e.WithAttr(tt.paramName, tt.paramValue)
			assert.Equal(t, tt.want, got.attrs)
		})
	}
}

func TestAppError_InstanceID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		instanceID    string
		errInstanceID *string
		want          string
	}{
		{
			name:          "test1",
			instanceID:    "",
			errInstanceID: nil,
			want:          "",
		},
		{
			name:          "test2",
			instanceID:    "test-id",
			errInstanceID: nil,
			want:          "test-id",
		},
		{
			name:          "test3",
			instanceID:    "",
			errInstanceID: func(s string) *string { return &s }("test-id"),
			want:          "test-id",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &AppError{
				instanceID:    tt.instanceID,
				errInstanceID: tt.errInstanceID,
			}
			got := e.InstanceID()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAppError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		message    string
		argsNames  []string
		args       []any
		instanceID string
		attrs      []mrmsg.NamedArg
		err        error
		want       string
	}{
		{
			name:       "test1",
			message:    "my-message",
			argsNames:  nil,
			args:       nil,
			instanceID: "",
			attrs:      nil,
			err:        nil,
			want:       "my-message",
		},
		{
			name:       "test2",
			message:    "my-message {{ .key1 }} - {{ .key2 }}",
			argsNames:  []string{"key1", "key2"},
			args:       []any{"value1", "value2"},
			instanceID: "test-id",
			attrs: []mrmsg.NamedArg{
				{
					Name:  "attr1",
					Value: "value2",
				},
				{
					Name:  "attr1",
					Value: "value2",
				},
			},
			err:  errors.New("test external error"),
			want: "[test-id] my-message value1 - value2 (attr1=value2, attr1=value2): test external error",
		},
		{
			name:      "test3",
			message:   "my-message {{ .key1 }} - {{ .key2 }}",
			argsNames: []string{"key1", "key2"},
			args:      []any{"value1", "value2", "value3"},
			want:      "[WARNING!!! too many arguments in error message] my-message value1 - value2",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &AppError{
				pureError: pureError{
					message:   tt.message,
					argsNames: tt.argsNames,
					args:      tt.args,
				},
				instanceID: tt.instanceID,
				attrs:      tt.attrs,
				err:        tt.err,
			}
			got := e.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}
