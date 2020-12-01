package sqlcat_test

import (
	"reflect"
	"testing"

	"github.com/h4ckedneko/sqlcat"
)

func testSQL(t *testing.T, s1, s2 string) {
	if s1 != s2 {
		t.Errorf("expected sql %q but got %q", s1, s2)
	}
}

func testArgs(t *testing.T, a1, a2 []interface{}) {
	if !reflect.DeepEqual(a1, a2) {
		t.Errorf("expected args %q but got %q", a1, a2)
	}
}

func TestBuilderToSQL(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets", sql)
}

func TestBuilderToSQLCount(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets) AS countq", sql)
}

func TestBuilderToSQLForColumns(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"id", "name", "type", "breed"},
	}
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT id, name, type, breed FROM pets", sql)
}

func TestBuilderToSQLCountForColumns(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"id", "name", "type", "breed"},
	}
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT id, name, type, breed FROM pets) AS countq", sql)
}

func TestBuilderToSQLForRelations(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:     "pets AS p",
		Columns:   []string{"*"},
		Relations: []string{"JOIN owners AS o ON o.id = p.owner_id"},
	}
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets AS p JOIN owners AS o ON o.id = p.owner_id", sql)
}

func TestBuilderToSQLCountForRelations(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:     "pets AS p",
		Columns:   []string{"*"},
		Relations: []string{"JOIN owners AS o ON o.id = p.owner_id"},
	}
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets AS p JOIN owners AS o ON o.id = p.owner_id) AS countq", sql)
}

func TestBuilderToSQLForConditions(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	sqlcat.WithCondition(builder, "name ILIKE $n", "%Inugami Korone%")
	sqlcat.WithCondition(builder, "(type = 'dog' OR type = 'god')")
	sqlcat.WithCondition(builder, "breed = ?", "spaniel")
	sql, args := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets WHERE name ILIKE $1 AND (type = 'dog' OR type = 'god') AND breed = ?", sql)
	testArgs(t, []interface{}{"%Inugami Korone%", "spaniel"}, args)
}

func TestBuilderToSQLCountForConditions(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	sqlcat.WithCondition(builder, "name ILIKE $n", "%Inugami Korone%")
	sqlcat.WithCondition(builder, "(type = 'dog' OR type = 'god')")
	sqlcat.WithCondition(builder, "breed = ?", "spaniel")
	sql, args := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets WHERE name ILIKE $1 AND (type = 'dog' OR type = 'god') AND breed = ?) AS countq", sql)
	testArgs(t, []interface{}{"%Inugami Korone%", "spaniel"}, args)
}

func TestBuilderToSQLForGroups(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"id", "count(owner_id)"},
		Groups:  []string{"id"},
		Having:  []string{"count(owner_id) > 1"},
	}
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT id, count(owner_id) FROM pets GROUP BY id HAVING count(owner_id) > 1", sql)
}

func TestBuilderToSQLCountForGroups(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"id", "count(owner_id)"},
		Groups:  []string{"id"},
		Having:  []string{"count(owner_id) > 1"},
	}
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT id, count(owner_id) FROM pets GROUP BY id HAVING count(owner_id) > 1) AS countq", sql)
}

func TestBuilderToSQLForOrders(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithOrders([]string{"name ASC"})
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets ORDER BY name ASC", sql)
}

func TestBuilderToSQLCountForOrders(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithOrders([]string{"name ASC"})
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets) AS countq", sql)
}

func TestBuilderToSQLForLimit(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithLimit(30)
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets LIMIT 30", sql)
}

func TestBuilderToSQLCountForLimit(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithLimit(30)
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets) AS countq", sql)
}

func TestBuilderToSQLForOffset(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithOffset(30)
	sql, _ := builder.ToSQL()
	testSQL(t, "SELECT * FROM pets OFFSET 30", sql)
}

func TestBuilderToSQLCountForOffset(t *testing.T) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	builder.WithOffset(30)
	sql, _ := builder.ToSQLCount()
	testSQL(t, "SELECT count(*) FROM (SELECT * FROM pets) AS countq", sql)
}

func BenchmarkBuilderToSQLBasic(b *testing.B) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	for i := 0; i < b.N; i++ {
		builder.ToSQL()
	}
}

func BenchmarkBuilderToSQLCountBasic(b *testing.B) {
	builder := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}
	for i := 0; i < b.N; i++ {
		builder.ToSQLCount()
	}
}

func BenchmarkBuilderToSQLComplex(b *testing.B) {
	builder := &sqlcat.Builder{
		Table:      "pets AS p",
		Columns:    []string{"p.id", "count(p.owner_id)"},
		Relations:  []string{"JOIN owners AS o ON o.id = p.owner_id"},
		Conditions: []string{"p.name ILIKE $1", "(p.type = 'dog' OR p.type = 'god')", "p.breed = $2"},
		Arguments:  []interface{}{"%Inugami Korone%", "spaniel"},
		Groups:     []string{"p.id"},
		Having:     []string{"count(p.owner_id) > 1"},
		Orders:     []string{"name ASC"},
		Limit:      30,
		Offset:     30,
	}
	for i := 0; i < b.N; i++ {
		builder.ToSQL()
	}
}

func BenchmarkBuilderToSQLCountComplex(b *testing.B) {
	builder := &sqlcat.Builder{
		Table:      "pets AS p",
		Columns:    []string{"p.id", "count(p.owner_id)"},
		Relations:  []string{"JOIN owners AS o ON o.id = p.owner_id"},
		Conditions: []string{"p.name ILIKE $1", "(p.type = 'dog' OR p.type = 'god')", "p.breed = $2"},
		Arguments:  []interface{}{"%Inugami Korone%", "spaniel"},
		Groups:     []string{"p.id"},
		Having:     []string{"count(p.owner_id) > 1"},
		Orders:     []string{"name ASC"},
		Limit:      30,
		Offset:     30,
	}
	for i := 0; i < b.N; i++ {
		builder.ToSQLCount()
	}
}
