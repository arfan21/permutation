package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi BruteForce dengan inputan n dan r dari file constant
func TestBruteForce(t *testing.T) {
	fmt.Printf("Brute Force P(%d, %d)\n", n, r)
	result := BruteForce(n, r)
	fmt.Println("Result :", result)

	if result != expectedOutput {
		t.Fatalf("Not equal: \n"+
			"expected: %d\n"+
			"actual  : %d", expectedOutput, result)
	}
}

//Benchmark untuk fungsi Dynamic untuk mengetahui waktu eksekusi
func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(n, r)
	}
}
