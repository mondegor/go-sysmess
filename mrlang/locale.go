package mrlang

import (
	"context"
	"fmt"

	"github.com/mondegor/go-sysmess/mrmsg"
)

type (
	Locale struct {
		langID   uint16
		langCode string
		messages map[string]string
		errors   map[string]mrmsg.ErrorMessage
	}

	localeConfig struct {
		LangID   uint16                        `yaml:"lang_id"`
		Messages map[string]string             `yaml:"messages"` // code -> message
		Errors   map[string]mrmsg.ErrorMessage `yaml:"errors"`   // code -> {reason, []details}
	}
)

var (
	stubLocale = &Locale{
		langID:   0,
		langCode: "stub-locale",
		messages: make(map[string]string, 0),
		errors:   make(map[string]mrmsg.ErrorMessage, 0),
	}
)

func (l *Locale) LangID() uint16 {
	return l.langID
}

func (l *Locale) LangCode() string {
	return l.langCode
}

func (l *Locale) WithContext(ctx context.Context) context.Context {
	return WithContext(ctx, l)
}

func (l *Locale) TranslateMessage(code, defaultMessage string, args ...mrmsg.NamedArg) string {
	value, ok := l.messages[code]

	if !ok {
		value = defaultMessage
	}

	if len(args) > 0 {
		value = mrmsg.Render(value, args)
	}

	return value
}

func (l *Locale) HasErrorCode(code string) bool {
	_, ok := l.errors[code]

	return ok
}

func (l *Locale) TranslateError(code, defaultMessage string, args ...mrmsg.NamedArg) mrmsg.ErrorMessage {
	value, ok := l.errors[code]

	if !ok {
		value = mrmsg.ErrorMessage{Reason: defaultMessage}
	}

	if len(args) > 0 {
		value.Reason = mrmsg.Render(value.Reason, args)

		for i := 0; i < len(value.Details); i++ {
			value.Details[i] = mrmsg.Render(value.Details[i], args)
		}
	}

	return value
}

func newLocale(langCode, filePath string) (*Locale, error) {
	cfg := localeConfig{}

	if err := parseFile(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing locale file '%s': %w", filePath, err)
	}

	if err := checkLocale(filePath, &cfg); err != nil {
		return nil, err
	}

	return &Locale{
		langID:   cfg.LangID,
		langCode: langCode,
		messages: cfg.Messages,
		errors:   cfg.Errors,
	}, nil
}

func checkLocale(filePath string, cfg *localeConfig) error {
	if cfg.LangID <= 0 {
		return fmt.Errorf("lang_id cannot be '%d' in locale %s", cfg.LangID, filePath)
	}

	for messCode, value := range cfg.Messages {
		if err := mrmsg.CheckParse(value); err != nil {
			return fmt.Errorf("message with code '%s' has error '%s' in locale %s", messCode, err, filePath)
		}
	}

	for errCode, value := range cfg.Errors {
		if err := mrmsg.CheckParse(value.Reason); err != nil {
			return fmt.Errorf("error.Reason with code '%s' has error '%s' in locale %s", errCode, err, filePath)
		}

		for n, detail := range value.Details {
			if err := mrmsg.CheckParse(detail); err != nil {
				return fmt.Errorf("error.Details[%d] with code '%s' has error '%s' in locale %s", n, errCode, err, filePath)
			}
		}
	}

	return nil
}
