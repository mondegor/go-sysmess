package strategy_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrprocess/period/strategy"
)

const minTimePeriod = time.Millisecond

func TestStaticPeriod(t *testing.T) {
	t.Parallel()

	// staticPeriod пропускает значение через fixedPeriod: нулевые и
	// отрицательные значения подменяются минимально допустимым периодом
	type testCase struct {
		name  string
		value time.Duration
		want  time.Duration
	}

	tests := []testCase{
		{name: "zero", value: 0, want: minTimePeriod},
		{name: "negative", value: -5 * time.Second, want: minTimePeriod},
		{name: "positive", value: 5 * time.Second, want: 5 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, strategy.NewStaticPeriod(tt.value).Period())
		})
	}
}

func TestDispersionPeriod_Deterministic(t *testing.T) {
	t.Parallel()

	// при ratio = 0 (а также при отрицательном/NaN ratio, который fixedCoefficient
	// приводит к 0) случайное отклонение отсутствует и результат детерминирован
	type testCase struct {
		name  string
		value time.Duration
		ratio float64
		want  time.Duration
	}

	tests := []testCase{
		{name: "zero ratio", value: time.Second, ratio: 0, want: time.Second},
		{name: "negative ratio is clamped to zero", value: time.Second, ratio: -1, want: time.Second},
		{name: "nan ratio is clamped to zero", value: time.Second, ratio: math.NaN(), want: time.Second},
		{name: "non positive period is clamped", value: 0, ratio: 0, want: minTimePeriod},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, strategy.NewDispersionPeriod(tt.value, tt.ratio).Period())
		})
	}
}

func TestDispersionPeriod_WithinBounds(t *testing.T) {
	t.Parallel()

	// результат всегда попадает в диапазон [value-value*ratio, value+value*ratio]
	type testCase struct {
		name     string
		value    time.Duration
		ratio    float64
		min, max time.Duration
	}

	tests := []testCase{
		{name: "half dispersion", value: time.Hour, ratio: 0.5, min: 30 * time.Minute, max: 90 * time.Minute},
		{name: "small ratio", value: time.Minute, ratio: 0.1, min: 54 * time.Second, max: 66 * time.Second},
		// ratio >= 1 приводится fixedCoefficient к 0.99
		{name: "ratio above one is clamped", value: time.Hour, ratio: 1.5, min: 36 * time.Second, max: 2 * time.Hour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := strategy.NewDispersionPeriod(tt.value, tt.ratio)
			for i := 0; i < 1000; i++ {
				got := p.Period()
				assert.GreaterOrEqual(t, got, tt.min)
				assert.LessOrEqual(t, got, tt.max)
			}
		})
	}
}

func TestDispersionPeriod_ExtremeValuesDoNotPanic(t *testing.T) {
	t.Parallel()

	// граничные значения не должны приводить к панике rand.Int64N(<=0)
	values := []time.Duration{math.MaxInt64, math.MaxInt64 - 1}
	ratios := []float64{0, 0.5, 0.99, 1.5}

	for _, value := range values {
		for _, ratio := range ratios {
			assert.NotPanics(t, func() {
				_ = strategy.NewDispersionPeriod(value, ratio).Period()
			})
		}
	}
}

func TestDelayedPeriod_DeterministicDelayAndAcceleration(t *testing.T) {
	t.Parallel()

	const base = 2 * time.Second

	t.Run("delay decays towards base", func(t *testing.T) {
		t.Parallel()

		// delayed=1s, decay=0.5, ratio=0: 3s, 2.5s, 2.25s, ...
		p := strategy.NewDelayedPeriod(time.Second, 0, 0.5, strategy.NewStaticPeriod(base))
		assert.Equal(t, 3*time.Second, p.Period())
		assert.Equal(t, 2500*time.Millisecond, p.Period())
		assert.Equal(t, 2250*time.Millisecond, p.Period())
	})

	t.Run("acceleration grows towards base", func(t *testing.T) {
		t.Parallel()

		// регрессия: при отрицательном delayed и decay>0 период монотонно растёт к base,
		// а не скачет случайным мусором (до фикса calcPeriod ломал ускорение)
		p := strategy.NewDelayedPeriod(-time.Second, 0, 0.5, strategy.NewStaticPeriod(base))
		assert.Equal(t, time.Second, p.Period())
		assert.Equal(t, 1500*time.Millisecond, p.Period())
		assert.Equal(t, 1750*time.Millisecond, p.Period())
	})
}

func TestDelayedPeriod_HugeDelayDoesNotOverflow(t *testing.T) {
	t.Parallel()

	// регрессия: saturating-сложение addPeriod не должно переполнять Duration в минус
	p := strategy.NewDelayedPeriod(math.MaxInt64, 0, 0.5, strategy.NewStaticPeriod(time.Hour))

	assert.NotPanics(t, func() {
		got := p.Period()
		assert.Positive(t, got)
	})
}

func TestDelayedPeriod_NilStrategyFallsBackToMinPeriod(t *testing.T) {
	t.Parallel()

	// при nil-стратегии используется статичный минимальный период
	p := strategy.NewDelayedPeriod(0, 0, 0, nil)
	assert.Equal(t, minTimePeriod, p.Period())
}
