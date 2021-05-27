package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi Dynamic dengan inputan n dan r dari file constant
func TestDynamic(t *testing.T) {
	fmt.Printf("Dynamic P(%d, %d)\n", N10, R2)
	result := Dynamic(N10, R2)
	fmt.Println("Result :", result)

	if result != expectedOutput {
		t.Fatalf("Not equal: \n"+
			"expected: %d\n"+
			"actual  : %d", expectedOutput, result)
	}
}

//Benchmark untuk fungsi Dynamic untuk mengetahui waktu eksekusi
func benchmarkDynamic(n, r int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(n, r)
	}
}

func BenchmarkDynamicN10R2(b *testing.B) {
	benchmarkDynamic(N10, R2, b)
}

func BenchmarkDynamicN5R4(b *testing.B) {
	benchmarkDynamic(N5, R4, b)
}

func BenchmarkDynamicN20R10(b *testing.B) {
	benchmarkDynamic(N20, R10, b)
}
