package config

import (
	"errors"
	"fmt"
	"time"
)

// ValidateTimeZones - валидирует указанный список имён часовых поясов:
// имена проверяются на непустоту, уникальность и существование в базе
// часовых поясов. Проверка останавливается на первом негодном имени
// и сообщает о нём ошибкой. Пустой список ошибкой не считается.
//
// Это единственное место, где список отвергается: timezone.NewLocationList
// ошибку не возвращает и негодные имена молча пропускает, поэтому вызов
// этой функции при загрузке конфигурации обязателен - иначе опечатка
// в имени пояса обнаружится не на старте, а на первом обращении к поясу,
// которого не оказалось в списке.
//
// ВНИМАНИЕ: проверка существования требует базу часовых поясов в системе;
// в минимальных образах (scratch, distroless) её нужно встроить
// в бинарник импортом `_ "time/tzdata"`, иначе будут отвергнуты все
// IANA-имена. Значения "UTC" и "Local" доступны всегда.
func ValidateTimeZones(names []string) error {
	uniqNames := make(map[string]bool, len(names))

	for _, name := range names {
		if name == "" {
			return errors.New("timezone name is required")
		}

		if uniqNames[name] {
			return fmt.Errorf("duplicate timezone name '%s'", name)
		}

		if _, err := time.LoadLocation(name); err != nil {
			return fmt.Errorf("error loading timezone (name='%s'): %w", name, err)
		}

		uniqNames[name] = true
	}

	return nil
}
