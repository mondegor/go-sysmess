package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrpostgres/builder/helper"
	"github.com/mondegor/go-sysmess/mrstorage"
)

// Make sure the helper.SQLSet conforms with the mrstorage.SQLSetHelper interface.
func TestSQLSetImplementsSQLSetHelper(t *testing.T) {
	assert.Implements(t, (*mrstorage.SQLSetHelper)(nil), &helper.SQLSet{})
}
