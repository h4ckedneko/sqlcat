package sqlcat

import "strings"

// ParseOrders parses an orders REST API query and returns
// its equivalent SQL grammar. Each order should follow
// the pattern `column:direction` which is URL-safe.
//
// For example:
//
// 	name          ➜ name
// 	name:asc      ➜ name ASC
// 	pets.name:asc ➜ pets.name ASC
//
func ParseOrders(orders []string) []string {
	for i, o := range orders {
		op := strings.Split(o, ":")
		if len(op) > 1 {
			op[1] = strings.ToUpper(op[1])
		}
		orders[i] = strings.Join(op, sepWS)
	}
	return orders
}
