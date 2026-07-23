package timezone_test

// база часовых поясов встраивается в тестовый бинарник, чтобы тесты
// проходили в минимальных образах, где она отсутствует в системе.
import (
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/util/timezone"
)

func TestLocationList_LocationByName(t *testing.T) {
	t.Parallel()

	list := timezone.NewLocationList([]string{"Europe/Moscow"})

	tests := []struct {
		name    string
		value   string
		want    *time.Location
		wantErr bool
	}{
		{
			name:  "registered iana name",
			value: "Europe/Moscow",
			want:  mustLoadLocation(t, "Europe/Moscow"),
		},
		{
			name:  "utc is registered by default",
			value: "UTC",
			want:  time.UTC,
		},
		{
			// пояс процесса наружу не отдаётся: в список он не входит,
			// поэтому промах сводится к поясу по умолчанию
			name:    "local is not registered",
			value:   "Local",
			want:    mustLoadLocation(t, "Europe/Moscow"),
			wantErr: true,
		},
		{
			// промах сводится к поясу по умолчанию, а не к UTC: в этом списке
			// они различаются, поэтому подмена одного другим была бы заметна
			name:    "empty value returns the default zone and error",
			value:   "",
			want:    mustLoadLocation(t, "Europe/Moscow"),
			wantErr: true,
		},
		{
			name:    "known but not registered name returns the default zone and error",
			value:   "Asia/Tokyo",
			want:    mustLoadLocation(t, "Europe/Moscow"),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := list.LocationByName(tc.value)

			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

// TestNewLocationList_SkipsInvalidNames - фиксирует, что негодные имена
// пропускаются, а не прерывают создание списка: имя, стоящее после негодного,
// должно быть зарегистрировано. Отвергать такой список - забота
// config.ValidateTimeZones, конструктор ошибку не возвращает.
func TestNewLocationList_SkipsInvalidNames(t *testing.T) {
	t.Parallel()

	list := timezone.NewLocationList([]string{"", "Nowhere/Nowhere", "Europe/Moscow"})

	t.Run("valid name after invalid ones is registered", func(t *testing.T) {
		t.Parallel()

		got, err := list.LocationByName("Europe/Moscow")

		require.NoError(t, err)
		assert.Equal(t, mustLoadLocation(t, "Europe/Moscow"), got)
	})

	t.Run("valid name after invalid ones is indexed by offset", func(t *testing.T) {
		t.Parallel()

		// Москва: +03:00 круглый год, перехода на летнее время нет
		got, ok := list.NameByOffset(3*time.Hour, false)

		require.True(t, ok)
		assert.Equal(t, "Europe/Moscow", got)
	})

	t.Run("unknown name is not registered", func(t *testing.T) {
		t.Parallel()

		got, err := list.LocationByName("Nowhere/Nowhere")

		require.Error(t, err)
		// негодные имена списком не зарегистрированы, поэтому поясом
		// по умолчанию стало первое годное имя
		assert.Equal(t, mustLoadLocation(t, "Europe/Moscow"), got)
	})

	t.Run("utc is registered by default", func(t *testing.T) {
		t.Parallel()

		utc, err := list.LocationByName("UTC")
		require.NoError(t, err)
		assert.Equal(t, time.UTC, utc)
	})

	t.Run("local is not registered", func(t *testing.T) {
		t.Parallel()

		// пояс процесса наружу не отдаётся даже как всегда доступное имя
		_, err := list.LocationByName("Local")
		require.Error(t, err)
	})
}

func TestLocationList_Default(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		names []string
		want  *time.Location
	}{
		{
			name:  "first name of the list",
			names: []string{"Europe/Moscow", "Asia/Tokyo"},
			want:  mustLoadLocation(t, "Europe/Moscow"),
		},
		{
			name:  "empty list falls back to utc",
			names: nil,
			want:  time.UTC,
		},
		{
			// Local наружу не отдаётся, поэтому поясом по умолчанию
			// становится следующее годное имя списка
			name:  "local is skipped",
			names: []string{"Local", "Europe/Moscow"},
			want:  mustLoadLocation(t, "Europe/Moscow"),
		},
		{
			name:  "unloadable names are skipped",
			names: []string{"", "Nowhere/Bad", "Asia/Tokyo"},
			want:  mustLoadLocation(t, "Asia/Tokyo"),
		},
		{
			name:  "entirely unusable list falls back to utc",
			names: []string{"", "Nowhere/Bad"},
			want:  time.UTC,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, timezone.NewLocationList(tc.names).Default())
		})
	}
}

// mustLoadLocation - загружает часовой пояс или прерывает тест.
func mustLoadLocation(t *testing.T, name string) *time.Location {
	t.Helper()

	loc, err := time.LoadLocation(name)
	require.NoError(t, err)

	return loc
}
