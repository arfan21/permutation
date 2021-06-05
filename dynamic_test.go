package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi Dynamic dengan inputan n dan r dari file constant
func TestDynamic(t *testing.T) {
	fmt.Printf("Dynamic P(%d, %d)\n", N12, R2)
	result := Dynamic(N12, R2)
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

func BenchmarkDynamicN12R2(b *testing.B) {
	benchmarkDynamic(N12, R2, b)
}

func BenchmarkDynamicN12R8(b *testing.B) {
	benchmarkDynamic(N12, R8, b)
}

func BenchmarkDynamicN12R12(b *testing.B) {
	benchmarkDynamic(N12, R12, b)
}

func BenchmarkDynamicN5R3(b *testing.B) {
	benchmarkDynamic(N5, R3, b)
}

func BenchmarkDynamicN12R3(b *testing.B) {
	benchmarkDynamic(N12, R3, b)
}

func BenchmarkDynamicN18R3(b *testing.B) {
	benchmarkDynamic(N18, R3, b)
}
