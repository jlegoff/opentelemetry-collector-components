package apmprocessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDbTableFromSqlMultipleTables(t *testing.T) {
	table, exists := NewSQLParser().ParseDbTableFromSQL("Select owners.* from Owners, users")
	assert.Equal(t, true, exists)
	assert.Equal(t, "owners", table)
}

func TestParseDbTableFromSqlJoin(t *testing.T) {
	table, exists := NewSQLParser().ParseDbTableFromSQL("Select * from users u join company c on u.company_id = c.id")
	assert.Equal(t, true, exists)
	assert.Equal(t, "users", table)
}
