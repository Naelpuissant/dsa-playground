# SkipList

Basicaly a skiplist with "shortcuts" called "levels". 
If height of each nodes is 0, it's an oredered linked list.

## Bench

For a personnal project I used [huandu/skiplist](github.com/huandu/skiplist), 
I want to bench my insert/search compared this implementation (simple, no arena).
```bash
go test -bench=. ./sl -cpuprofile=./profile.out
```

current bench results
```
goos: linux
goarch: amd64
pkg: ds/sl
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkSkiplistInsert-8                7611813               159.5 ns/op
BenchmarkHuanduSkiplistInsert-8         14052273                85.25 ns/op
BenchmarkSkiplistSearch-8               54602737                21.09 ns/op
BenchmarkHuanduSkiplistSearch-8         17877283                65.48 ns/op
PASS
ok      ds/sl   4.919s
```

## TODO
* Reuse the update slice
* Pool nodes
* Optimize random height