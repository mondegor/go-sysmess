package mrlang

import (
    "fmt"

    "github.com/ilyakaznacheev/cleanenv"
    "github.com/mondegor/go-sysmess/mrmsg"
)

type (
    Locale struct {
        langCode string
        messages map[string]string
        errors map[string]ErrorMessage
    }

    localeConfig struct {
        Messages map[string]string `yaml:"messages"`
        Errors map[string]ErrorMessage `yaml:"errors"`
    }
)

var (
	defaultLocale = &Locale{
        langCode: "default",
        messages: make(map[string]string, 0),
        errors: make(map[string]ErrorMessage, 0),
    }
)

func DefaultLocale() *Locale {
    return defaultLocale
}

func (l *Locale) LangCode() string {
    return l.langCode
}

func (l *Locale) TranslateMessage(id string, defaultMessage string, args ...mrmsg.NamedArg) string {
    value, ok := l.messages[id]

    if !ok {
        value = defaultMessage
    }

    if len(args) > 0 {
        value = mrmsg.Render(value, args)
    }

    return value
}

func (l *Locale) TranslateError(id string, defaultMessage string, args ...mrmsg.NamedArg) ErrorMessage {
    value, ok := l.errors[id]

    if !ok {
        value = ErrorMessage{Reason: defaultMessage}
    }

    if len(args) > 0 {
        value.Reason = mrmsg.Render(value.Reason, args)

        for i := 0; i < len(value.Details); i++ {
            value.Details[i] = mrmsg.Render(value.Details[i], args)
        }
    }

    return value
}

func newLocale(langCode string, filePath string) (*Locale, error) {
    cfg := localeConfig{}
    err := cleanenv.ReadConfig(filePath, &cfg)

    if err != nil {
        return nil, fmt.Errorf("while reading locale '%s', error '%s' occurred", filePath, err)
    }

    err = checkLocale(filePath, &cfg)

    if err != nil {
        return nil, err
    }

    return &Locale{
        langCode: langCode,
        messages: cfg.Messages,
        errors: cfg.Errors,
    }, nil
}

func checkLocale(filePath string, cfg *localeConfig) error {
    for messId, value := range cfg.Messages {
        err := mrmsg.CheckParse(value)

        if err != nil {
            return fmt.Errorf("message with id '%s' has error '%s' in locale %s", messId, err, filePath)
        }
    }

    for errId, value := range cfg.Errors {
        err := mrmsg.CheckParse(value.Reason)

        if err != nil {
            return fmt.Errorf("error.Reason with id '%s' has error '%s' in locale %s", errId, err, filePath)
        }

        for n, detail := range value.Details {
            err = mrmsg.CheckParse(detail)

            if err != nil {
                return fmt.Errorf("error.Details[%d] with id '%s' has error '%s' in locale %s", n, errId, err, filePath)
            }
        }
    }

    return nil
}
