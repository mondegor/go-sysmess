package mraccess

import (
	"strings"

	"github.com/mondegor/go-sysmess/mraccess/provider/filestorage"
	"github.com/mondegor/go-sysmess/mrlog"
)

// InitPermsProvider - создаёт и инициализирует провайдер ролей и прав из файлового хранилища.
// Параметры:
//   - rolesDirPath - путь к директории с файлами ролей;
//   - roles - список ролей.
func InitPermsProvider(
	logger mrlog.Logger,
	rolesDirPath string,
	roles []string,
) (*filestorage.PermsProvider, error) {
	mrlog.Info(logger, "Create and init roles and permissions for app")

	provider, err := filestorage.NewPermsProvider(rolesDirPath, roles)
	if err != nil {
		return nil, err
	}

	info := filestorage.ExtractProviderInfo(provider)

	mrlog.Info(logger, "Registered roles: "+strings.Join(info.Roles, ", "))
	mrlog.Debug(logger, "Registered rights:\n - "+strings.Join(info.Rights, ",\n - "))

	return provider, nil
}
