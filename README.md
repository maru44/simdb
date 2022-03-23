# simdb

WIP

```
go test -bench . -benchmem ./tests/bench
goos: darwin
goarch: arm64
pkg: github.com/maru44/simdb/tests/bench
BenchmarkInsert-8                4806235               306.1 ns/op           230 B/op          0 allocs/op
BenchmarkInsert_SyncMap-8        2393808               456.9 ns/op           199 B/op          5 allocs/op
BenchmarkGet-8                  12298252               112.9 ns/op             0 B/op          0 allocs/op
BenchmarkLoad-8                 12409465               116.5 ns/op             0 B/op          0 allocs/op
BenchmarkLoad_SyncMap-8          7687611               173.7 ns/op             0 B/op          0 allocs/op
BenchmarkDelete-8               11132658               124.2 ns/op             0 B/op          0 allocs/op
BenchmarkDelete_SyncMap-8        7858884               175.8 ns/op             0 B/op          0 allocs/op
```
