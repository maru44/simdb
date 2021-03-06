# simdb

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/maru44/scheman/blob/master/LICENSE)
![ActionsCI](https://github.com/maru44/simdb/workflows/Test%20Lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/maru44/simdb)](https://goreportcard.com/report/github.com/maru44/simdb)

A simple IMDB (in-memory database) generator for go.

We generate type-safe and thread-safe IMDB.

We generate table with columns and also implement some methods like `Get`, `Insert`, `Update`, `Delete`, etc...

## Usage

Basic Informations.

#### Table Info

| attribute  | explanation                                               | type      |          |
| ---------- | --------------------------------------------------------- | --------- | -------- |
| name       | name of table                                             | string    | required |
| is_private | whether private or not                                    | bool      |
| is_pointer | whether the value is pointer or not                       | bool      |
| key_type   | type of key (this table's primary and an only unique key) | string    | required |
| columns    | array of columns                                          | []Columns |          |

#### Column Info

| attribute  | explanation            | type   |          |
| ---------- | ---------------------- | ------ | -------- |
| name       | name of column         | string | required |
| type       | type of column         | string | requried |
| is_private | whether private or not | bool   |          |

#### Setting args

| arg     | explanation                   | positional or optional |                                                       |
| ------- | ----------------------------- | ---------------------- | ----------------------------------------------------- |
| file    | generated file name           | positional (1st arg)   | default is `db.go`                                    |
| dir     | generated directory name      | optional               | default is current directory                          |
| package | package name of generated     | optional               | requried (default is `main`) [can set in config file] |
| config  | config name withou extensions | optional               | required (default is `simdb`)                         |

### Examples

**_foo/gen.go_**

```go
package gen

//go:generate go run github.com/maru44/simdb db.go --dir=bar --package=bench --config=conf

```

**_foo/conf.yaml_**

```yaml
name: bench
is_private: true
key_type: int
columns:
  - name: name
    type: string
  - name: email
    type: string
  - name: age
    type: uint
    is_private: true
  - name: is_vald
    type: bool
```

In this case, the generated file would be like following.

**_foo/bar/db.go_**

```go
// Code generated by "github.com/maru44/simdb/gen"; DO NOT EDIT.

package bench

import (
	"fmt"
	"sync"
)

type (
	bench struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	benchs struct {
		data map[int]bench
		sync.RWMutex
	}
)

func NewBenchs() benchs {
	return benchs{
		data: map[int]bench{},
	}
}

func (t *benchs) List() map[int]bench {
	t.RLock()
	defer t.RUnlock()
	return t.data
}

...

```

## Performance

I measured bench mark.
It depends on struct size, so for your information only.

### Not pointer

I used this struct to measure bench mark.

```go
type (
	bench struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	benchs struct {
		data map[int]bench
		sync.RWMutex
	}
)
```

A little bit better than `sync.Map`.

```
goos: darwin
goarch: arm64
pkg: github.com/maru44/simdb/_tests/bench
BenchmarkInsert-8                4867156               235.6 ns/op           227 B/op          0 allocs/op
BenchmarkInsert_SyncMap-8        2842941               363.6 ns/op           184 B/op          5 allocs/op
BenchmarkGet-8                  12589046               113.1 ns/op             0 B/op          0 allocs/op
BenchmarkLoad-8                 12701738               111.3 ns/op             0 B/op          0 allocs/op
BenchmarkLoad_SyncMap-8          8800521               153.4 ns/op             0 B/op          0 allocs/op
BenchmarkDelete-8               12029200               120.3 ns/op             0 B/op          0 allocs/op
BenchmarkDelete_SyncMap-8       10508106               133.4 ns/op             0 B/op          0 allocs/op
```

### Pointer

I used this struct to measure bench mark.

```go
type (
	Pt struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	Pts struct {
		data map[int]*Pt
		sync.RWMutex
	}
)
```

Even in the case of pointer, a little bit better than `sync.Map`.

```
goos: darwin
goarch: arm64
pkg: github.com/maru44/simdb/tests/bench
BenchmarkInsertPt-8             10145770               171.2 ns/op            67 B/op          0 allocs/op
BenchmarkInsertPt_SyncMap-8      3924448               388.6 ns/op           174 B/op          4 allocs/op
BenchmarkGetPt-8                23999539                71.32 ns/op            0 B/op          0 allocs/op
BenchmarkLoadPt-8               23471499                65.13 ns/op            0 B/op          0 allocs/op
BenchmarkLoadPt_SyncMap-8        9493228               153.6 ns/op             0 B/op          0 allocs/op
BenchmarkDeletePt-8             15360548                90.44 ns/op            0 B/op          0 allocs/op
BenchmarkDeletePt_SyncMap-8     10498572               133.5 ns/op             0 B/op          0 allocs/op
```

### Comparing IsPointer or not

**pointer**

```go
type (
	Pt struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	Pts struct {
		data map[int]*Pt
		sync.RWMutex
	}
)
```

**not pointer**

```go
type (
	bench struct {
		Name    string
		Email   string
		age     uint
		IsValid bool
	}

	benchs struct {
		data map[int]bench
		sync.RWMutex
	}
)
```

In this struct size, pointer is better than not pointer.

```
goos: darwin
goarch: arm64
pkg: github.com/maru44/simdb/tests/bench
BenchmarkInsert-8                4867156               235.6 ns/op           227 B/op          0 allocs/op
BenchmarkInsertPt-8             10145770               171.2 ns/op            67 B/op          0 allocs/op
BenchmarkGet-8                  12589046               113.1 ns/op             0 B/op          0 allocs/op
BenchmarkLoad-8                 12701738               111.3 ns/op             0 B/op          0 allocs/op
BenchmarkGetPt-8                23999539                71.32 ns/op            0 B/op          0 allocs/op
BenchmarkLoadPt-8               23471499                65.13 ns/op            0 B/op          0 allocs/op
BenchmarkDelete-8               12029200               120.3 ns/op             0 B/op          0 allocs/op
BenchmarkDeletePt-8             15360548                90.44 ns/op            0 B/op          0 allocs/op
BenchmarkUpdate-8                9903958               132.9 ns/op             0 B/op          0 allocs/op
BenchmarkUpdatePt-8             15739252                94.40 ns/op            0 B/op          0 allocs/op
BenchmarkUpdateName-8            7570064               175.3 ns/op             0 B/op          0 allocs/op
BenchmarkUpdateNamePt-8         18335049                82.61 ns/op            0 B/op          0 allocs/op
```
