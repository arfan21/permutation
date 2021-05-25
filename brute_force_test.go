package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBruteForce(t *testing.T) {
	result := BruteForce(10, 2)
	fmt.Println(result)
	assert.Equal(t, result, 90)
}

func BenchmarkBruteForce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BruteForce(10, 2)
	}
}
