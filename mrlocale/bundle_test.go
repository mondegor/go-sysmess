package mrlocale_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"

	"github.com/mondegor/go-core/mrlocale"
)

// TestNewBundle_DefaultLanguage - проверяет выбор языка по умолчанию.
// Язык сверяется со списком поддерживаемых по разобранному тегу, поэтому
// одна и та же локаль, записанная по-разному ("ru-RU", "ru_RU", "ru-ru"),
// опознаётся как один и тот же язык.
func TestNewBundle_DefaultLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		languages []string
		def       string
		want      string
	}{
		{
			name:      "default is not specified, first language of the list wins",
			languages: []string{"ru-RU", "en-US"},
			def:       "",
			want:      "ru-RU",
		},
		{
			name:      "default matches the list entry byte to byte",
			languages: []string{"ru-RU", "en-US"},
			def:       "en-US",
			want:      "en-US",
		},
		{
			// подчёркивание - разделитель, который порождает генератор gotext,
			// поэтому такая запись попадает в конфигурацию и должна опознаваться
			name:      "underscore separated default matches hyphen separated list",
			languages: []string{"ru-RU", "en-US"},
			def:       "en_US",
			want:      "en-US",
		},
		{
			name:      "hyphen separated default matches underscore separated list",
			languages: []string{"ru_RU", "en_US"},
			def:       "en-US",
			want:      "en-US",
		},
		{
			name:      "default is case insensitive",
			languages: []string{"ru-RU", "en-US"},
			def:       "en-us",
			want:      "en-US",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opts := []mrlocale.BundleOption{
				mrlocale.WithMessageProvider(
					func(_ []language.Tag) (mrlocale.MessageProvider, error) {
						return stubProvider{}, nil
					},
				),
			}

			if tc.def != "" {
				opts = append(opts, mrlocale.WithDefaultLanguage(tc.def))
			}

			bundle, err := mrlocale.NewBundle(tc.languages, opts...)
			require.NoError(t, err)

			assert.Equal(t, tc.want, mrlocale.NewPool(bundle).Localizer().Language())
		})
	}
}

// TestNewBundle_Errors - проверяет отказы при создании бандла.
func TestNewBundle_Errors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		languages []string
		def       string
	}{
		{
			name:      "empty list of languages",
			languages: nil,
		},
		{
			name:      "language of the list is not well-formed",
			languages: []string{"ru-RU", "garbage!!!"},
		},
		{
			// регион EN не существует, поэтому такой тег отвергается
			name:      "region of the list language is unknown",
			languages: []string{"ru-RU", "en-EN"},
		},
		{
			name:      "duplicate language in the list",
			languages: []string{"ru-RU", "en-US", "ru-RU"},
		},
		{
			// разная запись одной локали - это тот же язык, поэтому тоже дубликат
			name:      "duplicate language written differently",
			languages: []string{"ru-RU", "ru_RU"},
		},
		{
			name:      "default language is not well-formed",
			languages: []string{"ru-RU", "en-US"},
			def:       "garbage!!!",
		},
		{
			name:      "default language is not in the list",
			languages: []string{"ru-RU", "en-US"},
			def:       "de-DE",
		},
		{
			// язык без региона не равен языку с регионом, поэтому не опознаётся
			name:      "default language without region does not match regional list entry",
			languages: []string{"ru-RU", "en-US"},
			def:       "en",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			opts := []mrlocale.BundleOption{
				mrlocale.WithMessageProvider(
					func(_ []language.Tag) (mrlocale.MessageProvider, error) {
						return stubProvider{}, nil
					},
				),
			}

			if tc.def != "" {
				opts = append(opts, mrlocale.WithDefaultLanguage(tc.def))
			}

			bundle, err := mrlocale.NewBundle(tc.languages, opts...)

			require.Error(t, err)
			assert.Nil(t, bundle)
		})
	}
}

// TestNewBundle_UndefinedLanguageIsSupported - фиксирует, что "und" является
// обычным языком списка: он не должен трактоваться как признак незаданного
// языка по умолчанию.
func TestNewBundle_UndefinedLanguageIsSupported(t *testing.T) {
	t.Parallel()

	bundle, err := mrlocale.NewBundle(
		[]string{"ru-RU", "und"},
		mrlocale.WithDefaultLanguage("und"),
		mrlocale.WithMessageProvider(
			func(_ []language.Tag) (mrlocale.MessageProvider, error) {
				return stubProvider{}, nil
			},
		),
	)
	require.NoError(t, err)

	assert.Equal(t, "und", mrlocale.NewPool(bundle).Localizer().Language())
}
