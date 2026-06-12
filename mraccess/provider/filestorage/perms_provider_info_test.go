package filestorage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mraccess/provider/filestorage"
)

func TestExtractProviderInfo_Nil(t *testing.T) {
	t.Parallel()

	info := filestorage.ExtractProviderInfo(nil)

	assert.Empty(t, info.Roles)
	assert.Empty(t, info.Rights)
}

func TestExtractProviderInfo_SortedAndUnion(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	// порядок ролей в файлах/правах намеренно несортирован
	writeRole(t, dir, "manager", []string{"orders.manage"}, []string{"orders.view"})
	writeRole(t, dir, "admin", []string{"users.manage"}, []string{"orders.manage"}) // дублирует право manager'а

	provider, err := filestorage.NewPermsProvider(dir, []string{"manager", "admin"})
	require.NoError(t, err)

	info := filestorage.ExtractProviderInfo(provider)

	// детерминированный отсортированный вывод
	assert.Equal(t, []string{"admin", "manager"}, info.Roles)
	// union прав без дублей, отсортирован
	assert.Equal(t, []string{"orders.manage", "orders.view", "users.manage"}, info.Rights)
}
