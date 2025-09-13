package main

import "testing"

func Add(a, b int) int {
	return a + b
}

// ========== BENCHMARKING

func BenchmarkAddSmallInput(b *testing.B) {
	//for i := 0; i < b.N ; i++
	for range b.N {
		Add(2, 3)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	//for i := 0; i < b.N ; i++
	for range b.N {
		Add(2, 3)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	//for i := 0; i < b.N ; i++
	for range b.N {
		Add(2, 3)
	}
}
