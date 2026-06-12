package filestorage

import (
	"maps"
	"slices"
)

type (
	// PermsProviderInfo - информация о зарегистрированных ролях и правах.
	// Используется для отладки и отображения в консоли.
	// Привилегии и разрешения объединены в единый отсортированный список Rights
	// (см. PermsProvider - они являются одним концептом права).
	PermsProviderInfo struct {
		Roles  []string
		Rights []string
	}
)

// ExtractProviderInfo - извлекает данные PermsProviderInfo из PermsProvider.
// Списки отсортированы для детерминированного вывода.
func ExtractProviderInfo(provider *PermsProvider) PermsProviderInfo {
	if provider == nil {
		return PermsProviderInfo{}
	}

	return PermsProviderInfo{
		Roles:  slices.Sorted(maps.Keys(provider.roleRights)),
		Rights: slices.Sorted(maps.Keys(provider.registered)),
	}
}
