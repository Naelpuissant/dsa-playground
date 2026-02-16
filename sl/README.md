# SkipList

Basicaly a skiplist with "shortcuts" called "levels". 
If height of each nodes is 0, it's an oredered linked list.

## Bench

For a personnal project I used [huandu/skiplist](github.com/huandu/skiplist),
I want to bench my insert/search compared this implementation (simple, no arena).
```bash
go test -benchtime=10000000x -bench=. ./sl -cpuprofile=./profile.out
```

current bench results
```
goos: linux
goarch: amd64
pkg: ds/sl
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkSkiplistInsert-8               10000000                85.86 ns/op
BenchmarkHuanduSkiplistInsert-8         10000000               117.7 ns/op
BenchmarkSkiplistSearch-8               10000000                86.60 ns/op
BenchmarkHuanduSkiplistSearch-8         10000000               297.8 ns/op
PASS
ok      ds/sl   6.051s
```
Search results looks kinda sus, how huandu's sl be so slow ?


## TODO
* Allow concurrency
* Optimize random height
