package main

import (
	"os"
	"testing"
)

func Benchmark_run(b *testing.B) {
	b.ResetTimer()
	run(os.Args)
}

func Benchmark_heavyFunc(b *testing.B) {
	b.ResetTimer()
	heavyFunc()
}

func Benchmark_simepleFunc(b *testing.B) {
	b.ResetTimer()
	simpleFunc()
}
