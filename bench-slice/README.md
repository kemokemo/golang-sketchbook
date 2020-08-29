# bench-slice

mattnさんの記事 [golang でパフォーマンスチューニングする際に気を付けるべきこと](https://mattn.kaoriya.net/software/lang/go/20161019124907.htm)を写経して、ベンチマークの書き方と計測方法の基礎を復習。

```sh
go test -count 10 -test.bench BenchmarkMakeSlice
```

改良前の測定結果を`slower.log`に保存、改良後の測定結果を`improved.log`に保存。

[benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat?tab=doc)ツールで比較する。


```sh
go get golang.org/x/perf/cmd/benchstat
```

```sh
benchstat slower.log improved.log
```

速度向上の効果を測定した結果が以下。

```sh
name          old time/op  new time/op  delta
MakeSlice-12   188ns ± 3%   143ns ± 1%  -23.78%  (p=0.000 n=9+8)
```
