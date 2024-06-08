package mrmsg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    Data
		want string
	}{
		{
			name: "data with string value",
			d:    map[string]any{"key1": "stringValue"},
			want: "{key1: stringValue}",
		},
		{
			name: "data with int value",
			d:    map[string]any{"key1": 1234},
			want: "{key1: 1234}",
		},
		{
			name: "data with boolean value",
			d:    map[string]any{"key1": true},
			want: "{key1: true}",
		},
		{
			name: "data with structure value",
			d: map[string]any{"key1": struct {
				key1 string
				key2 int
			}{"value", 1234}},
			want: "{key1: {value 1234}}",
		},
		{
			name: "data with a few values",
			d:    map[string]any{"key1": "stringValue", "key3": false, "key2": 1234},
			want: "{key1: stringValue, key2: 1234, key3: false}",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.d.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
