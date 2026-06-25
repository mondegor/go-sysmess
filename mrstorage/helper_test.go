package mrstorage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrstorage"
)

func TestNonZeroLimit(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		value int
		want  string
	}

	tests := []testCase{
		{
			name:  "negative",
			value: -10,
			want:  " LIMIT 1",
		},
		{
			name:  "zero",
			value: 0,
			want:  " LIMIT 1",
		},
		{
			name:  "one",
			value: 1,
			want:  " LIMIT 1",
		},
		{
			name:  "positive",
			value: 25,
			want:  " LIMIT 25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := mrstorage.NonZeroLimit(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
