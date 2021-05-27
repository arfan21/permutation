package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi BruteForce dengan inputan n dan r dari file constant
func TestBruteForce(t *testing.T) {
	fmt.Printf("Brute Force P(%d, %d)\n", N10, R2)
	result := BruteForce(N10, R2)
	fmt.Println("Result :", result)

	if result != expectedOutput {
		t.Fatalf("Not equal: \n"+
			"expected: %d\n"+
			"actual  : %d", expectedOutput, result)
	}
}

//Benchmark untuk fungsi Dynamic untuk mengetahui waktu eksekusi
func benchmarkBruteForce(n, r int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(n, r)
	}
}

func BenchmarkBruteForceN10R2(b *testing.B) {
	benchmarkBruteForce(N10, R2, b)
}

func BenchmarkBruteForceN5R4(b *testing.B) {
	benchmarkBruteForce(N5, R4, b)
}

func BenchmarkBruteForceN20R10(b *testing.B) {
	benchmarkBruteForce(N20, R10, b)
}
