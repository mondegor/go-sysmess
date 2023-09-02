package mrlang

import (
    "fmt"

    "github.com/ilyakaznacheev/cleanenv"
    "github.com/mondegor/go-sysmess/mrmsg"
)

type (
    locale struct {
        langCode string
        messages map[string]Message
        errors map[string]ErrorMessage
    }

    LocaleConfig struct {
        Messages map[string]Message `yaml:"messages"`
        Errors map[string]ErrorMessage `yaml:"errors"`
    }
)

func newLocale(langCode string, filePath string) (*locale, error) {
    cfg := LocaleConfig{}
    err := cleanenv.ReadConfig(filePath, &cfg)

    if err != nil {
        return nil, fmt.Errorf("while reading locale '%s', error '%s' occurred", filePath, err)
    }

    for id, value := range cfg.Messages {
        var args []mrmsg.NamedArg

        for _, arg := range mrmsg.ParseArgsNames(string(value)) {
            args = append(args, mrmsg.NewArg(arg, arg))
        }

        _, err = mrmsg.Render(string(value), args)

        if err != nil {
            return nil, fmt.Errorf("message with id '%s' has error '%s' in locale %s", id, err, filePath)
        }
    }

    return &locale{
        langCode: langCode,
        messages: cfg.Messages,
        errors: cfg.Errors,
    }, nil
}

func (l locale) LangCode() string {
    return l.langCode
}

func (l locale) TranslateMessage(id string, defaultMessage string, args ...mrmsg.NamedArg) Message {
    value, ok := l.messages[id]

    if !ok {
        value = Message(defaultMessage)
    }

    if len(args) > 0 {
        value = Message(mrmsg.MustRender(string(value), args))
    }

    return value
}

func (l locale) TranslateError(id string, defaultMessage string, args ...mrmsg.NamedArg) ErrorMessage {
    value, ok := l.errors[id]

    if !ok {
        value = ErrorMessage{Reason: defaultMessage}
    }

    if len(args) > 0 {
        value.Reason = mrmsg.MustRender(value.Reason, args)

        for i := 0; i < len(value.Details); i++ {
            value.Details[i] = mrmsg.MustRender(value.Details[i], args)
        }
    }

    return value
}
