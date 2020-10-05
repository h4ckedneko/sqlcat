package sqlcat_test

import (
	"testing"

	"github.com/h4ckedneko/sqlcat"
)

func testParsed(t *testing.T, o1, o2 string) {
	if o1 != o2 {
		t.Errorf("expected order %q but got %q", o1, o2)
	}
}

func TestParseOrders(t *testing.T) {
	orders := sqlcat.ParseOrders([]string{"name", "name:asc", "pets.name:asc"})
	testParsed(t, "name", orders[0])
	testParsed(t, "name ASC", orders[1])
	testParsed(t, "pets.name ASC", orders[2])
}
