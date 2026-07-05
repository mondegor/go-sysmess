package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrpostgres/builder/helper"
	"github.com/mondegor/go-sysmess/mrstorage"
)

// Make sure the helper.SQLCondition conforms with the mrstorage.SQLConditionHelper interface.
func TestSQLConditionImplementsSQLConditionHelper(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrstorage.SQLConditionHelper)(nil), &helper.SQLCondition{})
}
