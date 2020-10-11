# sqlcat

[![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/h4ckedneko/sqlcat)
[![Latest Version](https://img.shields.io/github/v/release/h4ckedneko/sqlcat?label=latest)](https://github.com/h4ckedneko/sqlcat/releases)
[![License Name](https://img.shields.io/github/license/h4ckedneko/sqlcat?color=blue)](https://github.com/h4ckedneko/sqlcat/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/h4ckedneko/sqlcat/Testing)](https://github.com/h4ckedneko/sqlcat/actions?query=workflow:Testing)
[![Coverage Status](https://gocover.io/_badge/github.com/h4ckedneko/sqlcat)](https://gocover.io/github.com/h4ckedneko/sqlcat)
[![Report Card Status](https://goreportcard.com/badge/github.com/h4ckedneko/sqlcat)](https://goreportcard.com/report/github.com/h4ckedneko/sqlcat)

Package sqlcat is a dead simple SQL query builder for Go. It is not an ORM nor a full-featured query builder, only SELECT is supported and you still need to write some SQL. Its job is to structurally concatenate those SQL to form a complete statement.

**Features and benefits:**

-   Minimal API, it only has one struct and two helper functions.
-   Helps you write queries in a structured and declarative way.
-   Outputs prepared statement to prevent SQL injection.
-   Supports PostgreSQL positional parameters.
-   No imported external dependencies.
-   Properly tested with benchmarks.

## Installation

Make sure you have a working [Go](https://golang.org/doc/install) workspace, then:

```
go get github.com/h4ckedneko/sqlcat
```

For updating to latest stable release, do:

```
go get -u github.com/h4ckedneko/sqlcat
```

## Usage

Here is a basic example for this package:

```go
package main

import (
	"fmt"

	"github.com/h4ckedneko/sqlcat"
)

func main() {
	// Initialize the builder.
	b := &sqlcat.Builder{
		Table:   "pets",
		Columns: []string{"*"},
	}

	// Associate various context.
	b.WithOrders([]string{"name ASC"})
	b.WithLimit(30)
	b.WithOffset(30)

	// Associate a condition.
	// It prevents SQL injection.
	typ := "cat"
	cond := "type = $n"
	sqlcat.WithCondition(b, cond, typ)

	// Build the SQL query and its arguments.
	// Output sql: SELECT * FROM pets WHERE type = $1 ORDER BY name ASC LIMIT 30 OFFSET 30
	// Output args: [cat]
	sql, args := b.ToSQL()
	fmt.Println(sql)
	fmt.Println(args)
}
```

See [examples](https://github.com/h4ckedneko/sqlcat/tree/master/examples) for more advanced real-world examples.

## Performance

You can run benchmarks by yourself using `make bench` command.

```
BenchmarkBuilderToSQLBasic-2             3536516               321 ns/op              40 B/op          2 allocs/op
BenchmarkBuilderToSQLCountBasic-2        3427706               348 ns/op              96 B/op          2 allocs/op
BenchmarkBuilderToSQLComplex-2            767654              1518 ns/op             736 B/op          8 allocs/op
BenchmarkBuilderToSQLCountComplex-2       901642              1322 ns/op             736 B/op          8 allocs/op
```

## Alternatives

If you want a full-featured SQL query builder, check out these packages:

-   [goqu](https://github.com/doug-martin/goqu)
-   [squirrel](https://github.com/Masterminds/squirrel)
-   [loukoum](https://github.com/ulule/loukoum)

## Inspiration

This package is based on [Miniflux](https://github.com/miniflux/miniflux)'s code for its [EntryQueryBuilder](https://github.com/miniflux/miniflux/blob/master/storage/entry_query_builder.go).

## License

MIT © Lyntor Paul Figueroa. See [LICENSE](https://github.com/h4ckedneko/sqlcat/blob/master/LICENSE) for full license text.
