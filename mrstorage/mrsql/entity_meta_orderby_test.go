package mrsql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrstorage/mrsql"
	"github.com/mondegor/go-core/mrtype"
)

// Make sure the mrsql.EntityMetaOrderBy conforms with the mrtype.ListSorter interface.
func TestEntityMetaOrderByImplementsListSorter(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrtype.ListSorter)(nil), &mrsql.EntityMetaOrderBy{})
}
