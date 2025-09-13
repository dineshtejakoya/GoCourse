package advanced

import (
	"math/rand"
	"testing"
)

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size %d, received %d", size, len(slice))
	}
}

// Now Benchmark the random slice Generator
// We have to start our benchmark function with Benchmark keyword
func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		SumSlice(slice)
	}
}
