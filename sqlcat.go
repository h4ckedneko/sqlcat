package sqlcat

import (
	"strconv"
	"strings"
)

const (
	sepWS    = " "
	sepComma = ", "
	sepAnd   = " AND "
)

// A Builder builds an SQL query by concatenating its fields.
type Builder struct {
	Table      string
	Columns    []string
	Relations  []string
	Conditions []string
	Arguments  []interface{}
	Orders     []string
	Limit      int
	Offset     int
}

// ToSQL returns the built SQL query and its arguments.
func (b *Builder) ToSQL() (string, []interface{}) {
	return buildSQL(b, false)
}

// ToSQLCount is like ToSQL, but for counting.
func (b *Builder) ToSQLCount() (string, []interface{}) {
	return buildSQL(b, true)
}

// WithOrders associates an orders into the query.
func (b *Builder) WithOrders(orders []string) {
	if len(orders) > 0 {
		b.Orders = orders
	}
}

// WithLimit associates a limit into the query.
func (b *Builder) WithLimit(limit int) {
	if limit > 0 {
		b.Limit = limit
	}
}

// WithOffset associates an offset into the query.
func (b *Builder) WithOffset(offset int) {
	if offset > 0 {
		b.Offset = offset
	}
}

func buildSQL(b *Builder, count bool) (string, []interface{}) {
	var sb strings.Builder
	if count {
		sb.WriteString("SELECT count(*)")
		sb.WriteString(buildSQLFrom(b))
		sb.WriteString(buildSQLJoin(b))
		sb.WriteString(buildSQLWhere(b))
	} else {
		sb.WriteString("SELECT" + buildSQLExp(b))
		sb.WriteString(buildSQLFrom(b))
		sb.WriteString(buildSQLJoin(b))
		sb.WriteString(buildSQLWhere(b))
		sb.WriteString(buildSQLOrderBy(b))
		sb.WriteString(buildSQLLimit(b))
		sb.WriteString(buildSQLOffset(b))
	}
	return sb.String(), b.Arguments
}

func buildSQLFrom(b *Builder) string {
	if b.Table == "" {
		return ""
	}
	return " FROM " + b.Table
}

func buildSQLExp(b *Builder) string {
	if len(b.Columns) < 1 {
		return ""
	}
	return " " + strings.Join(b.Columns, sepComma)
}

func buildSQLJoin(b *Builder) string {
	if len(b.Relations) < 1 {
		return ""
	}
	return " " + strings.Join(b.Relations, sepWS)
}

func buildSQLWhere(b *Builder) string {
	if len(b.Conditions) < 1 {
		return ""
	}
	return " WHERE " + strings.Join(b.Conditions, sepAnd)
}

func buildSQLOrderBy(b *Builder) string {
	if len(b.Orders) < 1 {
		return ""
	}
	return " ORDER BY " + strings.Join(b.Orders, sepComma)
}

func buildSQLLimit(b *Builder) string {
	if b.Limit < 1 {
		return ""
	}
	return " LIMIT " + strconv.Itoa(b.Limit)
}

func buildSQLOffset(b *Builder) string {
	if b.Offset < 1 {
		return ""
	}
	return " OFFSET " + strconv.Itoa(b.Offset)
}

// WithCondition associates a condition into the query. In PostgreSQL,
// you can use positional parameters by inserting `$n` to the condition.
// It will be replaced by a number staring from the current arguments size.
func WithCondition(b *Builder, cond string, args ...interface{}) {
	if len(args) > 0 {
		for _, arg := range args {
			b.Arguments = append(b.Arguments, arg)
			pos := "$" + strconv.Itoa(len(b.Arguments))
			cond = strings.Replace(cond, "$n", pos, 1)
		}
	}
	b.Conditions = append(b.Conditions, cond)
}
