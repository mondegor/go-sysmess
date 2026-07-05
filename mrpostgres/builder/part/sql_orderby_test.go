package part_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrpostgres/builder/part"
	"github.com/mondegor/go-core/mrstorage"
)

// Make sure the builder.SQLOrderByBuilder conforms with the mrstorage.SQLOrderByBuilder interface.
func TestSQLOrderByBuilderImplementsSQLOrderByBuilder(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrstorage.SQLOrderByBuilder)(nil), &part.SQLOrderByBuilder{})
}
