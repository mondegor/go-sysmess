package ordered_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/util/slices/ordered"
)

func TestOrdered_FilterFunc(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		s    []uint64
		want []uint64
	}

	tests := []testCase{
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

			assert.Equal(t, tt.want, ordered.FilterFunc(tt.s, func(el uint64) bool { return el > 0 }))
		})
	}
}

func TestOrdered_SortedUnique(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		s    []uint64
		want []uint64
	}

	tests := []testCase{
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

			assert.Equal(t, tt.want, ordered.SortedUnique(tt.s))
		})
	}
}

func TestOrdered_SortedUniqueClone(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		s    []uint64
		want []uint64
	}

	tests := []testCase{
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
			s:    []uint64{1002, 1000, 1001, 1000},
			want: []uint64{1000, 1001, 1002},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			orig := slices.Clone(tt.s)

			got := ordered.SortedUniqueClone(tt.s)

			assert.Equal(t, tt.want, got)
			// исходный слайс не должен измениться
			assert.Equal(t, orig, tt.s)
		})
	}
}

func TestOrdered_SortedUniqueClone_Generic_Uint8(t *testing.T) {
	t.Parallel()

	s := []uint8{30, 10, 20, 10}

	got := ordered.SortedUniqueClone(s)

	assert.Equal(t, []uint8{10, 20, 30}, got)
	assert.Equal(t, []uint8{30, 10, 20, 10}, s) // вход не изменился
}

func TestOrdered_BinaryIndex(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		s     []uint64
		value uint64
		want  int
	}

	tests := []testCase{
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

			assert.Equal(t, tt.want, ordered.BinaryIndex(tt.s, tt.value))
		})
	}
}

func TestOrdered_BinaryAppend(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		s     []uint64
		value uint64
		want  []uint64
	}

	tests := []testCase{
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

			got := ordered.BinaryAppend(tt.s, tt.value)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOrdered_UniqueBinaryAppend(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		s     []uint64
		value uint64
		want  []uint64
	}

	tests := []testCase{
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

			got := ordered.UniqueBinaryAppend(tt.s, tt.value)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOrdered_BinaryContains(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		s     []uint64
		value uint64
		want  bool
	}

	tests := []testCase{
		{
			name:  "empty slice",
			s:     nil,
			value: 1000,
			want:  false,
		},
		{
			name:  "single element found",
			s:     []uint64{1000},
			value: 1000,
			want:  true,
		},
		{
			name:  "element in the middle",
			s:     []uint64{1000, 1001, 1002},
			value: 1001,
			want:  true,
		},
		{
			name:  "value greater than all",
			s:     []uint64{1000, 1001, 1002},
			value: 1003,
			want:  false,
		},
		{
			name:  "value less than all",
			s:     []uint64{1000, 1001, 1002},
			value: 999,
			want:  false,
		},
		{
			name:  "gap between elements",
			s:     []uint64{1000, 1002, 1003},
			value: 1001,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ordered.BinaryContains(tt.s, tt.value))
		})
	}
}

// TestOrdered_Generic_Uint8 - проверяет работу обобщённых функций
// на типе, отличном от uint64.
func TestOrdered_Generic_Uint8(t *testing.T) {
	t.Parallel()

	s := []uint8{0, 10, 0, 30, 20, 0}

	s = ordered.FilterFunc(s, func(el uint8) bool { return el > 0 })
	assert.Equal(t, []uint8{10, 30, 20}, s)

	s = ordered.SortedUnique(s)
	assert.Equal(t, []uint8{10, 20, 30}, s)

	s = ordered.UniqueBinaryAppend(s, 25)
	assert.Equal(t, []uint8{10, 20, 25, 30}, s)

	s = ordered.BinaryAppend(s, 25)
	assert.Equal(t, []uint8{10, 20, 25, 25, 30}, s)

	assert.Equal(t, 1, ordered.BinaryIndex(s, uint8(20)))
	assert.True(t, ordered.BinaryContains(s, uint8(30)))
	assert.False(t, ordered.BinaryContains(s, uint8(99)))
}

// TestOrdered_Boundary_Uint8 - проверяет корректность сравнений и вставки
// у граничных значений малого беззнакового типа (около math.MaxUint8),
// где у обобщённой версии нет запаса по разрядности.
func TestOrdered_Boundary_Uint8(t *testing.T) {
	t.Parallel()

	s := []uint8{0, 1, 254}

	// вставка максимума типа в конец отсортированного массива
	s = ordered.UniqueBinaryAppend(s, 255)
	assert.Equal(t, []uint8{0, 1, 254, 255}, s)

	// повторная вставка максимума не дублирует его
	s = ordered.UniqueBinaryAppend(s, 255)
	assert.Equal(t, []uint8{0, 1, 254, 255}, s)

	assert.True(t, ordered.BinaryContains(s, uint8(0)))
	assert.True(t, ordered.BinaryContains(s, uint8(255)))
	assert.Equal(t, 3, ordered.BinaryIndex(s, uint8(255)))
}
