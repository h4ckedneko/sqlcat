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
	Groups     []string
	Having     []string
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
	if !count {
		sb.WriteString("SELECT" + buildSQLExp(b))
		sb.WriteString(buildSQLFrom(b))
		sb.WriteString(buildSQLJoin(b))
		sb.WriteString(buildSQLWhere(b))
		sb.WriteString(buildSQLGroupBy(b))
		sb.WriteString(buildSQLHaving(b))
		sb.WriteString(buildSQLOrderBy(b))
		sb.WriteString(buildSQLLimit(b))
		sb.WriteString(buildSQLOffset(b))
	} else {
		sb.WriteString("SELECT count(*) FROM (")
		sb.WriteString("SELECT" + buildSQLExp(b))
		sb.WriteString(buildSQLFrom(b))
		sb.WriteString(buildSQLJoin(b))
		sb.WriteString(buildSQLWhere(b))
		sb.WriteString(buildSQLGroupBy(b))
		sb.WriteString(buildSQLHaving(b))
		sb.WriteString(") AS countq")
	}
	return sb.String(), b.Arguments
}

func buildSQLFrom(b *Builder) string {
	if b.Table != "" {
		return " FROM " + b.Table
	}
	return ""
}

func buildSQLExp(b *Builder) string {
	if len(b.Columns) > 0 {
		return " " + strings.Join(b.Columns, sepComma)
	}
	return ""
}

func buildSQLJoin(b *Builder) string {
	if len(b.Relations) > 0 {
		return " " + strings.Join(b.Relations, sepWS)
	}
	return ""
}

func buildSQLWhere(b *Builder) string {
	if len(b.Conditions) > 0 {
		return " WHERE " + strings.Join(b.Conditions, sepAnd)
	}
	return ""
}

func buildSQLGroupBy(b *Builder) string {
	if len(b.Groups) > 0 {
		return " GROUP BY " + strings.Join(b.Groups, sepComma)
	}
	return ""
}

func buildSQLHaving(b *Builder) string {
	if len(b.Having) > 0 {
		return " HAVING " + strings.Join(b.Having, sepAnd)
	}
	return ""
}

func buildSQLOrderBy(b *Builder) string {
	if len(b.Orders) > 0 {
		return " ORDER BY " + strings.Join(b.Orders, sepComma)
	}
	return ""
}

func buildSQLLimit(b *Builder) string {
	if b.Limit > 0 {
		return " LIMIT " + strconv.Itoa(b.Limit)
	}
	return ""
}

func buildSQLOffset(b *Builder) string {
	if b.Offset > 0 {
		return " OFFSET " + strconv.Itoa(b.Offset)
	}
	return ""
}

// WithCondition associates a condition into the query.
// For PostgreSQL, you can use positional parameters in a
// condition by placing `$n` which automatically increments.
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
