package mrsql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrstorage/mrsql"
	"github.com/mondegor/go-sysmess/mrtype"
)

// Make sure the mrsql.EntityMetaOrderBy conforms with the mrtype.ListSorter interface.
func TestEntityMetaOrderByImplementsListSorter(t *testing.T) {
	assert.Implements(t, (*mrtype.ListSorter)(nil), &mrsql.EntityMetaOrderBy{})
}
