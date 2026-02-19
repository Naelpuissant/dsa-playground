# SkipList

Basicaly an ordered linked list with "shortcuts" called "levels". 
If height of each nodes is 0, it's an oredered linked list.
Thread-safe, mutex locking.

## Bench

For a personnal project I used [huandu/skiplist](github.com/huandu/skiplist),
I want to bench my insert/search compared this implementation (simple, no arena).
```bash
go test -benchtime=5s -bench=. ./sl
```

current bench results
```
goos: linux
goarch: amd64
pkg: ds/sl
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkSkiplistInsert-8               19274250               335.0 ns/op
BenchmarkHuanduSkiplistInsert-8         14218832               528.8 ns/op
BenchmarkSkiplistSearch-8               25595133               226.9 ns/op
BenchmarkHuanduSkiplistSearch-8         16565250               368.3 ns/op
PASS
ok      ds/sl   27.591s
```
Looks like my thread safe implementation beats what I previously used, yahou!

## TODO
* [X] Allow concurrency
* [X] Optimize random height
