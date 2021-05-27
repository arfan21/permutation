package permutation

import (
	"fmt"
	"testing"
)

//Testing fungsi Dynamic dengan inputan n dan r dari file constant
func TestDynamic(t *testing.T) {
	fmt.Printf("Dynamic P(%d, %d)\n", n, r)
	result := Dynamic(n, r)
	fmt.Println("Result :", result)

	if result != expectedOutput {
		t.Fatalf("Not equal: \n"+
			"expected: %d\n"+
			"actual  : %d", expectedOutput, result)
	}
}

//Benchmark untuk fungsi Dynamic untuk mengetahui waktu eksekusi
func BenchmarkDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(n, r)
	}
}
