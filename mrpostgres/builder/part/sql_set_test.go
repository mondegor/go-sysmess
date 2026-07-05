package part_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrpostgres/builder/part"
	"github.com/mondegor/go-core/mrstorage"
)

// Make sure the builder.SQLSetBuilder conforms with the mrstorage.SQLSetBuilder interface.
func TestSQLSetBuilderImplementsSQLSetBuilder(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrstorage.SQLSetBuilder)(nil), &part.SQLSetBuilder{})
}
