package suint64_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/util/slices/suint64"
)

func TestSortedUint64_FilterFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    []uint64
		want []uint64
	}{
		{
			name: "test1",
			s:    nil,
			want: nil,
		},
		{
			name: "test2",
			s:    []uint64{1000},
			want: []uint64{1000},
		},
		{
			name: "test3",
			s:    []uint64{1000, 1001},
			want: []uint64{1000, 1001},
		},
		{
			name: "test4",
			s:    []uint64{1000, 0, 1001},
			want: []uint64{1000, 1001},
		},
		{
			name: "test5",
			s:    []uint64{0, 1000},
			want: []uint64{1000},
		},
		{
			name: "test6",
			s:    []uint64{1000, 0},
			want: []uint64{1000},
		},
		{
			name: "test7",
			s:    []uint64{0, 0, 1000, 0, 0, 0, 1001, 0, 0, 0, 1002, 0, 0},
			want: []uint64{1000, 1001, 1002},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, suint64.FilterFunc(tt.s, func(el uint64) bool { return el > 0 }))
		})
	}
}

func TestSortedUint64_SortedUnique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    []uint64
		want []uint64
	}{
		{
			name: "test1",
			s:    nil,
			want: nil,
		},
		{
			name: "test2",
			s:    []uint64{1000},
			want: []uint64{1000},
		},
		{
			name: "test3",
			s:    []uint64{1000, 1001},
			want: []uint64{1000, 1001},
		},
		{
			name: "test4",
			s:    []uint64{1001, 1000},
			want: []uint64{1000, 1001},
		},
		{
			name: "test5",
			s:    []uint64{1001, 1000, 1002},
			want: []uint64{1000, 1001, 1002},
		},
		{
			name: "test6",
			s:    []uint64{1002, 1001, 1000},
			want: []uint64{1000, 1001, 1002},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, suint64.SortedUnique(tt.s))
		})
	}
}

func TestSortedUint64_BinaryIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		s     []uint64
		value uint64
		want  int
	}{
		{
			name:  "test1",
			s:     nil,
			value: 1000,
			want:  -1,
		},
		{
			name:  "test2",
			s:     []uint64{1000},
			value: 1000,
			want:  0,
		},
		{
			name:  "test3",
			s:     []uint64{1000, 1001, 1002},
			value: 1001,
			want:  1,
		},
		{
			name:  "test4",
			s:     []uint64{1000, 1001, 1002},
			value: 1003,
			want:  -1,
		},
		{
			name:  "test5",
			s:     []uint64{1000, 1002, 1003},
			value: 1001,
			want:  -1,
		},
		{
			name:  "test6",
			s:     []uint64{1000, 1001, 1002},
			value: 999,
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, suint64.BinaryIndex(tt.s, tt.value))
		})
	}
}

func TestSortedUint64_BinaryAppend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		s     []uint64
		value uint64
		want  []uint64
	}{
		{
			name:  "test1",
			s:     nil,
			value: 1000,
			want:  []uint64{1000},
		},
		{
			name:  "test2",
			s:     []uint64{1000},
			value: 1000,
			want:  []uint64{1000, 1000},
		},
		{
			name:  "test3",
			s:     []uint64{999, 1000},
			value: 1000,
			want:  []uint64{999, 1000, 1000},
		},
		{
			name:  "test4",
			s:     []uint64{999},
			value: 1000,
			want:  []uint64{999, 1000},
		},
		{
			name:  "test5",
			s:     []uint64{1001},
			value: 1000,
			want:  []uint64{1000, 1001},
		},
		{
			name:  "test6",
			s:     []uint64{999, 1001},
			value: 1000,
			want:  []uint64{999, 1000, 1001},
		},
		{
			name:  "test7",
			s:     []uint64{999, 1000, 1001},
			value: 1000,
			want:  []uint64{999, 1000, 1000, 1001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := suint64.BinaryAppend(tt.s, tt.value)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSortedUint64_UniqueBinaryAppend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		s     []uint64
		value uint64
		want  []uint64
	}{
		{
			name:  "test1",
			s:     nil,
			value: 1000,
			want:  []uint64{1000},
		},
		{
			name:  "test2",
			s:     []uint64{1000},
			value: 1000,
			want:  []uint64{1000},
		},
		{
			name:  "test3",
			s:     []uint64{999, 1000},
			value: 1000,
			want:  []uint64{999, 1000},
		},
		{
			name:  "test4",
			s:     []uint64{999, 1000, 1001},
			value: 1000,
			want:  []uint64{999, 1000, 1001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := suint64.UniqueBinaryAppend(tt.s, tt.value)

			assert.Equal(t, tt.want, got)
		})
	}
}
