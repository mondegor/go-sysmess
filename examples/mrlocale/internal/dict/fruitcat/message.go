package fruitcat

import (
	"strings"
)

type (
	// FruitMessage - модель объекта используемая при его локализации.
	FruitMessage struct {
		ID          string
		Name        string
		Description string
	}
)

// String - возвращает объект в виде строки.
func (m FruitMessage) String() string {
	if m.Description == "" {
		return m.Name
	}

	return m.ID + "|" + m.Name + "|" + m.Description
}

// ParseMessage - парсинг объекта хранящегося в справочнике фруктов в FruitMessage.
func ParseMessage(translation string) (msg FruitMessage) {
	attrs := strings.SplitN(translation, "|", 3)

	if len(attrs) > 0 {
		msg.ID = attrs[0]
	}

	if len(attrs) > 1 {
		msg.Name = attrs[1]
	}

	if len(attrs) > 2 {
		msg.Description = attrs[2]
	}

	return msg
}
