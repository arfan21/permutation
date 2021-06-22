package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi BruteForce dengan inputan n dan r dari file constant
func TestBruteForce(t *testing.T) {
	fmt.Printf("Brute Force P(%d, %d)\n", N12, R2)
	result := BruteForce(N12, R2)
	fmt.Println("Result :", result)

	if result != expectedOutput {
		t.Fatalf("Not equal: \n"+
			"expected: %d\n"+
			"actual  : %d", expectedOutput, result)
	}
}

//Benchmark untuk fungsi BruteForce untuk mengetahui waktu eksekusi
func benchmarkBruteForce(n, r int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(n, r)
	}
}

func BenchmarkBruteForceN12R2(b *testing.B) {
	benchmarkBruteForce(N12, R2, b)
}

func BenchmarkBruteForceN12R8(b *testing.B) {
	benchmarkBruteForce(N12, R8, b)
}

func BenchmarkBruteForceN12R12(b *testing.B) {
	benchmarkBruteForce(N12, R12, b)
}

func BenchmarkBruteForceN5R3(b *testing.B) {
	benchmarkBruteForce(N5, R3, b)
}

func BenchmarkBruteForceN12R3(b *testing.B) {
	benchmarkBruteForce(N12, R3, b)
}

func BenchmarkBruteForceN18R3(b *testing.B) {
	benchmarkBruteForce(N18, R3, b)
}
