package main

import (
	"testing"
)

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}
