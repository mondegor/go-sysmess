package helper

import (
	"errors"
	"fmt"

	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// ErrorStatusMapper - маппинг кодов пользовательских ошибок в HTTP-статусы ответа.
	// Позволяет настроить соответствие между конкретными кодами ошибок и статусами,
	// а также задать статусы по умолчанию для системных, внутренних и необработанных ошибок.
	ErrorStatusMapper struct {
		defaultStatus    int
		systemStatus     int
		internalStatus   int
		unexpectedStatus int
		code2status      map[string]int
	}
)

// NewErrorStatusMapper - создаёт объект ErrorStatusMapper.
// codeStatus - плоский слайс пар "код(string)/статус(int)", например: ["code1", 400, "code2", 404].
func NewErrorStatusMapper(
	defaultStatus, systemStatus, internalStatus, unexpectedStatus int,
	codeStatus []any,
) (*ErrorStatusMapper, error) {
	code2status := make(map[string]int, len(codeStatus)/2)

	for i := 0; i < len(codeStatus); i += 2 {
		code, ok := codeStatus[i].(string)
		if !ok {
			return nil, fmt.Errorf("codeStatus[%d] contains invalid code type, expected string", i)
		}

		status, ok := codeStatus[i+1].(int)
		if !ok {
			return nil, fmt.Errorf("codeStatus[%d] contains invalid status type, expected int", i+1)
		}

		code2status[code] = status
	}

	return &ErrorStatusMapper{
		defaultStatus:    defaultStatus,
		systemStatus:     systemStatus,
		internalStatus:   internalStatus,
		unexpectedStatus: unexpectedStatus,
		code2status:      code2status,
	}, nil
}

// ErrorStatus - возвращает HTTP-статус на основе типа и кода ошибки.
// Алгоритм:
//   - kind.User: ищет код ошибки в маппинге code2status, проверяя цепочку обёрнутых ошибок.
//     Если код не найден, возвращает defaultStatus.
//   - kind.System: возвращает systemStatus.
//   - kind.Internal: возвращает internalStatus.
//   - Остальные (без метода Kind()): возвращает unexpectedStatus.
func (m *ErrorStatusMapper) ErrorStatus(err error) int {
	switch kind.Analyze(err) {
	case kind.User:
		for {
			if e, ok := err.(interface{ Code() string }); ok {
				if status, ok := m.code2status[e.Code()]; ok {
					return status
				}

				if err = errors.Unwrap(err); err != nil {
					continue
				}
			}

			return m.defaultStatus
		}
	case kind.System:
		return m.systemStatus
	case kind.Internal:
		return m.internalStatus
	default:
		// если ошибка явно необработанна разработчиком (не имеет метода Kind()),
		// то отображается указанный m.unexpectedStatus
		return m.unexpectedStatus
	}
}
