package mrmsg

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"
)

const (
	leftDelim  = "{{"
	rightDelim = "}}"
)

var (
	regexpArgName = regexp.MustCompile(`^\.[A-Za-z][A-Za-z0-9]*$`)
)

func Render(message string, args []NamedArg) string {
	value, err := render(message, args)

	if err != nil {
		return message
	}

	return value
}

func render(message string, args []NamedArg) (string, error) {
	if message == "" {
		return "", nil
	}

	templ, err := template.New("").Parse(message)

	if err != nil {
		return "", err
	}

	data := make(map[string]string, len(args))

	for _, item := range args {
		data[item.name] = item.valueString()
	}

	var msg bytes.Buffer

	err = templ.Execute(&msg, data)

	if err != nil {
		return "", err
	}

	return msg.String(), nil
}

func CheckParse(message string) error {
	var args []NamedArg

	for _, arg := range ParseArgsNames(message) {
		args = append(args, NamedArg{arg, arg})
	}

	_, err := render(message, args)

	return err
}

func ParseArgsNames(message string) []string {
	var argsNames []string
	var keys map[string]bool

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
			keys = make(map[string]bool, 1)
		} else {
			if _, ok := keys[name]; ok {
				continue
			}
		}

		argsNames = append(argsNames, name)
		keys[name] = true
	}

	return argsNames
}
