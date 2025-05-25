package mrargs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrargs"
)

func TestData_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		group mrargs.Group
		want  string
	}{
		{
			name:  "data with string value",
			group: map[string]any{"key1": "stringValue"},
			want:  `{"key1": "stringValue"}`,
		},
		{
			name:  "data with int value",
			group: map[string]any{"key1": 1234},
			want:  `{"key1": 1234}`,
		},
		{
			name:  "data with boolean value",
			group: map[string]any{"key1": true},
			want:  `{"key1": true}`,
		},
		{
			name:  "data with boolean value",
			group: map[string]any{"key1": nil},
			want:  `{"key1": null}`,
		},
		{
			name: "data with structure value",
			group: map[string]any{"key1": struct {
				key2 string
				key3 int
			}{`value and "quotes"`, 1234}},
			want: `{"key1": "{key2:value and \"quotes\" key3:1234}"}`,
		},
		{
			name:  "data with a few values",
			group: map[string]any{"key1": `stringValue and "quotes"`, "key3": false, "key2": 1234.56, "key4": nil, "key5": 1234.0},
			want:  `{"key1": "stringValue and \"quotes\"", "key2": 1234.56, "key3": false, "key4": null, "key5": 1234}`,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.group.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
