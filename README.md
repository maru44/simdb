# simdb

WIP

```
go test -bench . -benchmem ./tests/bench
goos: darwin
goarch: arm64
pkg: github.com/maru44/simdb/tests/bench
BenchmarkInsert-8                4448840               263.0 ns/op           248 B/op          0 allocs/op
BenchmarkInsert_SyncMap-8        2779429               434.4 ns/op           186 B/op          5 allocs/op
BenchmarkGet-8                  12139834               113.8 ns/op             0 B/op          0 allocs/op
BenchmarkGet_SyncMap-8           7816982               173.2 ns/op             0 B/op          0 allocs/op
```
