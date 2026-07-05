package mrtype

import "time"

type (
	// RangeInt64 - целочисленный интервал [Min, Max].
	RangeInt64 struct {
		Min int64
		Max int64
	}

	// RangeFloat64 - вещественный интервал [Min, Max].
	RangeFloat64 struct {
		Min float64
		Max float64
	}

	// RangeDateTime - временной интервал [From, To].
	RangeDateTime struct {
		From time.Time
		To   time.Time
	}
)

// Transform - преобразовывает в RangeFloat64 с умножением полей на указанный коэффициент
// (для приведения к необходимой единицы измерения).
func (r RangeInt64) Transform(coefficient float64) RangeFloat64 {
	return RangeFloat64{
		Min: float64(r.Min) * coefficient,
		Max: float64(r.Max) * coefficient,
	}
}
