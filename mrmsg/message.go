package mrmsg

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strings"
)

const (
	leftDelim  = "{{"
	rightDelim = "}}"
)

var regexpArgName = regexp.MustCompile(`^\.[A-Za-z][A-Za-z0-9]*$`)

// Render - возвращает сформированное сообщение со вставленными в него значениями переменных.
func Render(message string, vars map[string]string) (string, error) {
	if message == "" {
		return "", nil
	}

	// TODO: можно закешировать парсинг сообщений

	templ, err := template.New("").Parse(message)
	if err != nil {
		return "", fmt.Errorf("parse message error: %w", err)
	}

	var msg bytes.Buffer

	if err = templ.Execute(&msg, vars); err != nil {
		return "", fmt.Errorf("execute template error: %w", err)
	}

	return msg.String(), nil
}

// MustRender - возвращает сформированное сообщение со вставленными в него значениями переменных.
// В случае ошибки логирует причину и возвращает сообщение без вставленных переменных.
func MustRender(message string, vars map[string]string) string {
	value, err := Render(message, vars)
	if err != nil {
		log.Print(fmt.Errorf("MustRender: '%s' message rendering failed: %w", message, err).Error())

		return message
	}

	return value
}

// RenderWithNamedArgs - возвращает сформированное сообщение со вставленными в него именованными параметрами.
func RenderWithNamedArgs(message string, args []NamedArg) (string, error) {
	vars := make(map[string]string, len(args))

	for _, arg := range args {
		vars[arg.Name] = arg.ValueString()
	}

	return Render(message, vars)
}

// MustRenderWithNamedArgs - возвращает сформированное сообщение со вставленными в него именованными параметрами.
// В случае ошибки логирует причину и возвращает сообщение без вставленных параметров.
func MustRenderWithNamedArgs(message string, args []NamedArg) string {
	value, err := RenderWithNamedArgs(message, args)
	if err != nil {
		log.Print(fmt.Errorf("MustRenderWithNamedArgs: '%s' message rendering failed: %w", message, err).Error())

		return message
	}

	return value
}

// RenderWithData - возвращает сформированное сообщение со вставленными в него данными.
func RenderWithData(message string, data Data) (string, error) {
	vars := make(map[string]string, len(data))

	for key := range data {
		vars[key] = ToString(data[key])
	}

	return Render(message, vars)
}

// MustRenderWithData - возвращает сформированное сообщение со вставленными в него данными.
// В случае ошибки логирует причину и возвращает сообщение без вставленных данных.
func MustRenderWithData(message string, data Data) string {
	value, err := RenderWithData(message, data)
	if err != nil {
		log.Print(fmt.Errorf("MustRenderWithData: '%s' message rendering failed: %w", message, err).Error())

		return message
	}

	return value
}

// CheckRender - если указанное сообщение содержит параметры,
// то проверяется их корректность.
func CheckRender(message string) error {
	names := ParseArgsNames(message)
	vars := make(map[string]string, len(names))

	for _, name := range names {
		vars[name] = name
	}

	if _, err := Render(message, vars); err != nil {
		return fmt.Errorf("check render error: %w", err)
	}

	return nil
}

// ParseArgsNames - извлечение параметров из указанного сообщения.
func ParseArgsNames(message string) []string {
	var (
		argsNames []string
		keys      map[string]bool
	)

	for {
		pos1 := strings.Index(message, leftDelim)

		if pos1 < 0 {
			break
		}

		message = message[pos1+len(leftDelim):]
		pos2 := strings.Index(message, rightDelim)

		// 4 = space + . + char + space
		if pos2 < 0+4 {
			break
		}

		name := message[:pos2]
		message = message[pos2+len(rightDelim):]
		last := len(name) - 1

		if name[0] != ' ' || name[last] != ' ' { // required spaces
			continue
		}

		name = name[1:last] // trim spaces

		if !regexpArgName.MatchString(name) {
			continue
		}

		name = name[1:] // left trim .

		if keys == nil {
			keys = make(map[string]bool)
		} else {
			// если название аргумента уже попадалось, то оно пропускается
			if _, ok := keys[name]; ok {
				continue
			}
		}

		argsNames = append(argsNames, name)
		keys[name] = true
	}

	return argsNames
}
