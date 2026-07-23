package timezone

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Имена параметров заголовка часового пояса и предельное число его сегментов.
// Сегментов у годного заголовка не больше трёх (имя и два параметра), поэтому
// потолок взят с запасом в один и на разборе годных заголовков не сказывается.
// Нужен из-за того, что заголовок задаёт клиент: без потолка значение вида
// "a;a;a;..." размером с лимит net/http на заголовки давало бы тысячи разборов
// на один запрос.
const (
	acceptParamOffset = "offset"
	acceptParamDST    = "dst"

	maxAcceptItems = 4
)

// ErrInvalidAcceptTimeZone - из значения заголовка не удалось извлечь
// ни имени, ни смещения.
var ErrInvalidAcceptTimeZone = errors.New("neither name nor offset could be extracted from the accept time zone value")

type (
	// AcceptTimeZone - разобранное значение заголовка часового пояса.
	// Любая из частей может отсутствовать: имя - при пустом первом сегменте,
	// смещение - при отсутствии или негодности параметров offset и dst.
	AcceptTimeZone struct {
		Name      string        // IANA-имя из первого сегмента, "" при его отсутствии.
		Offset    time.Duration // смещение относительно UTC, годно только при HasOffset.
		IsDST     bool
		HasOffset bool // true только когда годны оба параметра offset и dst.
	}
)

// ParseAcceptTimeZone - разбирает значение заголовка часового пояса
// вида "Europe/Moscow;offset=+03:00;dst=0". Возвращает ErrInvalidAcceptTimeZone,
// если из значения не удалось извлечь ни имени, ни смещения.
//
// Первый сегмент без "=" считается IANA-именем, остальные - параметрами.
// Смещение считается заданным, только когда годны оба параметра: подбор по одному
// лишь смещению без признака летнего времени давал бы пояс наугад, а признак
// без смещения бесполезен.
//
// Разбор строгий: пробелы внутри значения не допускаются нигде - ни вокруг сегментов,
// ни вокруг "=". Заголовок задаёт не человек, а клиентский код по контракту приложения,
// поэтому единственная годная форма ровно одна, и всякая другая отбрасывается,
// а не приводится к ней.
//
// Рассматриваются только первые maxAcceptItems сегментов, остальные отбрасываются
// без разбора.
func ParseAcceptTimeZone(value string) (AcceptTimeZone, error) {
	// SplitN с запасом в один сегмент: хвост за потолком попадает в последний
	// сегмент целиком и отбрасывается вместе с ним, не порождая разборов
	// по числу точек с запятой
	items := strings.SplitN(value, ";", maxAcceptItems+1)

	if len(items) > maxAcceptItems {
		items = items[:maxAcceptItems]
	}

	var (
		result    AcceptTimeZone
		offset    time.Duration
		isDST     bool
		hasOffset bool
		hasDST    bool
	)

	for i, item := range items {
		key, val, isParam := strings.Cut(item, "=")
		if !isParam {
			// именем считается только первый сегмент: сегмент без "=" в хвосте -
			// мусор, и принять его за имя значило бы разбирать негодное значение
			if i == 0 {
				result.Name = item
			}

			continue
		}

		switch key {
		case acceptParamOffset:
			if d, err := parseOffset(val); err == nil {
				offset, hasOffset = d, true
			}
		case acceptParamDST:
			if b, err := strconv.ParseBool(val); err == nil {
				isDST, hasDST = b, true
			}
		}
	}

	if hasOffset && hasDST {
		result.Offset, result.IsDST, result.HasOffset = offset, isDST, true
	}

	if result.Name == "" && !result.HasOffset {
		return AcceptTimeZone{}, ErrInvalidAcceptTimeZone
	}

	return result, nil
}

// parseOffset - разбирает смещение относительно UTC в формате ±HH:MM
// (например: "+03:00", "-07:30"). Знак обязателен, чтобы "03:00" не принималось
// за смещение с угаданным знаком.
func parseOffset(value string) (time.Duration, error) {
	// time.ParseDuration понимает "3h30m", но не "±HH:MM", поэтому смещение
	// разбирается вручную; time.Parse со схемой "-07:00" разобрал бы его,
	// но вернул бы время, из которого смещение пришлось бы доставать Zone()
	t, err := time.Parse("-07:00", value)
	if err != nil {
		return 0, err
	}

	_, seconds := t.Zone()

	return time.Duration(seconds) * time.Second, nil
}
