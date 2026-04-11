package suint64

import "slices"

// FilterFunc - фильтрует слайс uint64 на месте (без выделения нового слайса).
// Удаляет элементы, для которых check возвращает false.
// Возвращает усечённый слайс.
func FilterFunc(s []uint64, check func(el uint64) bool) []uint64 {
	for i := 0; i < len(s); i++ {
		if check(s[i]) {
			continue
		}

		s2 := s[i:]
		for i2 := 1; i2 < len(s2); i2++ {
			if check(s2[i2]) {
				s[i] = s2[i2]
				i++
			}
		}

		clear(s[i:]) // zero/nil out the obsolete elements, for GC

		return s[:i]
	}

	return s
}

// SortedUnique - возвращает уникальный отсортированный массив элементов (не копию).
func SortedUnique(s []uint64) []uint64 {
	slices.Sort(s)

	return slices.Compact(s)
}

// BinaryIndex - возвращает индекс массива, где содержится указанный элемент массива.
// Указанный массив обязан быть предварительно отсортирован.
func BinaryIndex(s []uint64, value uint64) int {
	i := binaryIndex(s, value)

	if i == -1 || s[i] != value {
		return -1
	}

	return i
}

// BinaryContains - сообщает, содержится ли указанный элемент в массиве.
// Указанный массив обязан быть предварительно отсортирован.
func BinaryContains(s []uint64, value uint64) bool {
	return BinaryIndex(s, value) >= 0
}

// BinaryAppend - добавляет элемент в массив оставляя последний отсортированным.
// Указанный массив обязан быть предварительно отсортирован.
func BinaryAppend(s []uint64, value uint64) []uint64 {
	return binaryAppend(s, value, false)
}

// UniqueBinaryAppend - добавляет элемент в массив, если такого ещё не было,
// оставляя массив отсортированным.
// Указанный массив обязан быть предварительно отсортирован.
func UniqueBinaryAppend(s []uint64, value uint64) []uint64 {
	return binaryAppend(s, value, true)
}

func binaryAppend(s []uint64, value uint64, unique bool) []uint64 {
	i := binaryIndex(s, value)

	if unique && i != -1 && value == s[i] {
		return s
	}

	if i == -1 || i == len(s)-1 && value > s[i] {
		return append(s, value)
	}

	c := append(s, 0) //nolint:gocritic
	copy(c[i+1:], s[i:])
	c[i] = value

	return c
}

func binaryIndex(s []uint64, value uint64) int {
	i, j := 0, len(s)

	for i < j {
		// avoid overflow
		h := int(uint(i+j) >> 1)

		switch {
		case s[h] < value:
			i = h + 1
		case s[h] > value:
			j = h
		default:
			return h
		}
	}

	if i < len(s) {
		return i
	}

	return i - 1
}
