package strategy

import (
	"math"
	"math/rand/v2"
	"time"
)

const (
	// minTimePeriod - минимально допустимый период тика.
	minTimePeriod = time.Millisecond
)

type (
	// Strategy - интерфейс стратегии определения периода тиков.
	// Метод Period() возвращает актуальную длительность периода,
	// которая может меняться в зависимости от реализации стратегии.
	Strategy interface {
		Period() time.Duration
	}
)

// calcPeriod - вычисляет период со случайным отклонением.
// Генерируется случайное значение в диапазоне
// [value-value*ratio, value+value*ratio] с равномерным распределением.
// Значение value может быть отрицательным (используется в стратегии ускорения).
func calcPeriod(value time.Duration, ratio float64) time.Duration {
	delta := time.Duration(float64(value) * ratio)
	minInterval := value - delta
	maxInterval := value + delta

	if minInterval > maxInterval {
		minInterval, maxInterval = maxInterval, minInterval
	}

	// ширина диапазона вычисляется в беззнаковой арифметике,
	// чтобы корректно обработать отрицательные границы без переполнения int64
	widthInterval := uint64(maxInterval) - uint64(minInterval) //nolint:gosec

	// если widthInterval не помещается в int64, то диапазон считается
	// некорректным и подменяется безопасным значением [1ms..1min]
	if widthInterval >= math.MaxInt64 {
		minInterval, widthInterval = minTimePeriod, uint64(time.Minute)
	}

	return minInterval + time.Duration(rand.Int64N(int64(widthInterval)+1)) //nolint:gosec
}

// addPeriod - складывает две длительности с защитой от переполнения int64
// (saturating-сложение): при переполнении возвращается граничное значение.
func addPeriod(a, b time.Duration) time.Duration {
	sum := a + b

	switch {
	case a > 0 && b > 0 && sum <= 0:
		return math.MaxInt64
	case a < 0 && b < 0 && sum >= 0:
		return math.MinInt64
	default:
		return sum
	}
}

// fixedPeriod - предотвращает нулевые и отрицательные значения.
func fixedPeriod(value time.Duration) time.Duration {
	if value <= 0 {
		return minTimePeriod
	}

	return value
}

// fixedCoefficient - предотвращает значения больше или равные 1 и отрицательные значения.
func fixedCoefficient(value float64) float64 {
	if value < 0 || math.IsNaN(value) {
		value = 0
	}

	if value >= 1 || math.IsInf(value, 0) {
		value = 0.99
	}

	return value
}
