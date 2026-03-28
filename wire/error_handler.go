package wire

import (
	"context"

	"github.com/mondegor/go-sysmess/errors"
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/mrlog"
)

// InitErrorHandler - создаёт объект errors.Handler.
func InitErrorHandler(logger mrlog.Logger) errors.Handler {
	return errors.HandlerFunc(
		func(ctx context.Context, err error) {
			switch kind.Analyze(err) {
			case kind.User:
				// 1. пользовательские ошибки логируются только в отладочном режиме;
				logger.Debug(ctx, "ErrorHandler: user error", "error", err)
			case kind.System, kind.Internal:
				// 2. пользовательские ошибки с вложенной runtime ошибкой;
				// 3. runtime ошибки;
				logger.Error(ctx, "ErrorHandler", "error", err)
			default:
				// 4. остальные ошибки у которых нет метода Kind() (требуется найти место их возникновения и правильно обработать);
				logger.Error(ctx, "ErrorHandler: unexpected error", "error", err)
			}
		},
	)
}
