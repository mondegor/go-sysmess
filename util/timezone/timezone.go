package timezone

import (
	"errors"
	"fmt"
	"time"
)

// Имена часовых поясов, которые доступны всегда и не требуют базы часовых поясов.
const (
	nameUTC   = "UTC"
	nameLocal = "Local"
)

type (
	// LocationList - хранит предзагруженные часовые пояса, доступные приложению,
	// и индекс для их подбора по смещению. И пояса, и индекс готовятся один раз
	// при создании списка, чтобы прикладной код не платил за это на каждом обращении:
	// time.LoadLocation читает базу часовых поясов с диска, а индекс строится
	// обходом года по каждому поясу, причём обход обходится дороже загрузки
	// примерно впятеро (см. BenchmarkStd_LoadLocation и BenchmarkNewLocationList).
	LocationList struct {
		locations   map[string]*time.Location
		offsetIndex map[offsetKey]string
	}
)

// NewLocationList - создаёт объект LocationList на основе списка имён часовых поясов.
// Поддерживаются IANA-имена (например: "Europe/Moscow"), а также "UTC" и "Local",
// которые регистрируются всегда и указывать их в списке не требуется.
//
// Ошибку не возвращает: пустые и не найденные в базе часовых поясов имена молча
// пропускаются, список создаётся из оставшихся, а не попавший в него пояс отличим
// только по результату обращения (LocationByName - ошибкой, NameByOffset - false).
// Поэтому вызов config.ValidateTimeZones при загрузке конфигурации обязателен:
// это единственное место, где негодный список отвергается.
//
// ВНИМАНИЕ: для IANA-имён требуется база часовых поясов в системе;
// в минимальных образах (scratch, distroless) её нужно встроить
// в бинарник импортом `_ "time/tzdata"`.
//
// ВАЖНО: стоимость создания линейна по числу имён и составляет порядка 55 мкс
// на пояс (BenchmarkNewLocationList_ByCount), поэтому список рассчитан
// на однократное создание при старте приложения, а не на создание по требованию.
func NewLocationList(names []string) *LocationList {
	locations := make(map[string]*time.Location, len(names)+2)

	locations[nameUTC] = time.UTC
	locations[nameLocal] = time.Local

	// индекс наполняется в порядке поступления имён, поэтому при совпадении
	// пары (смещение, признак летнего времени) выигрывает последнее имя;
	// UTC регистрируется после списка и потому свою пару не уступает никому
	offsetIndex := make(map[offsetKey]string, len(names)*2+1)

	// момент фиксируется один раз, чтобы все пояса списка индексировались
	// по одному годовому окну, а не каждый по своему
	now := time.Now()

	for _, name := range names {
		// пустое имя пропускается отдельно: time.LoadLocation("") молча
		// возвращает UTC без ошибки, и незаполненная настройка
		// зарегистрировалась бы поясом с пустым именем
		if name == "" {
			continue
		}

		// сюда имя доходит только при пропущенной валидации
		// либо при отсутствии самой базы часовых поясов в образе
		loc, err := time.LoadLocation(name)
		if err != nil {
			continue
		}

		locations[name] = loc

		// Local в индекс не попадает даже при явном указании в списке:
		// он совпал бы со смещением процесса, но как результат подбора
		// бесполезен, вызывающему требуется IANA-имя
		if name == nameLocal {
			continue
		}

		addToOffsetIndex(offsetIndex, name, loc, now)
	}

	// UTC регистрируется напрямую, а не обходом года: состояние у него одно
	// и известно заранее, поэтому обход дал бы ту же единственную пару
	offsetIndex[offsetKey{offset: 0, isDST: false}] = nameUTC

	return &LocationList{
		locations:   locations,
		offsetIndex: offsetIndex,
	}
}

// LocationByName - возвращает часовой пояс по указанному имени,
// если пояс не зарегистрирован в списке, то возвращается time.UTC и ошибка.
func (l *LocationList) LocationByName(value string) (*time.Location, error) {
	if value == "" {
		return time.UTC, errors.New("arg 'value' is empty")
	}

	if loc, ok := l.locations[value]; ok {
		return loc, nil
	}

	return time.UTC, fmt.Errorf("timezone not found for arg '%s'", value)
}
