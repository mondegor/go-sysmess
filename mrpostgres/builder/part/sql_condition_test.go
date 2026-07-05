package part_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrpostgres/builder/part"
	"github.com/mondegor/go-sysmess/mrstorage"
)

// Make sure the builder.SQLConditionBuilder conforms with the mrstorage.SQLConditionBuilder interface.
func TestSQLConditionBuilderImplementsSQLConditionBuilder(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrstorage.SQLConditionBuilder)(nil), &part.SQLConditionBuilder{})
}
