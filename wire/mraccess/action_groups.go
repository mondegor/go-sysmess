package mraccess

import (
	"context"
	"strings"

	"github.com/mondegor/go-core/mraccess"
	"github.com/mondegor/go-core/mraccess/config"
)

type (
	logger interface {
		Debug(ctx context.Context, msg string, args ...any)
	}
)

// InitActionGroups - создаёт и инициализирует группу действий (ActionGroup) из конфигурации.
func InitActionGroups(logger logger, groups []config.ActionGroup) (name2group map[string]mraccess.ActionGroup) {
	name2group = make(map[string]mraccess.ActionGroup, len(groups))

	for _, group := range groups {
		basePath := strings.TrimRight("/"+strings.Trim(group.BasePath, "/"), "/") + "/" // "/" or "/abc/"

		logger.Debug(
			context.Background(),
			"Init actionGroups with privilege and base path",
			"name", group.Name,
			"privilege", group.Privilege,
			"basePath", basePath,
		)

		name2group[group.Name] = mraccess.ActionGroup{
			Name:      group.Name,
			Privilege: group.Privilege,
			BasePath:  basePath,
		}
	}

	return name2group
}
