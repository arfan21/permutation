package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBruteForce(t *testing.T) {
	fmt.Printf("Brute Force P(%d, %d)\n",n , r)
	result := BruteForce(n, r)
	fmt.Println("Result :",result)
	assert.Equal(t, result, expectedOutput)
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForce(n, r)
	}
}
