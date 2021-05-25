package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamic(t *testing.T) {
	fmt.Printf("Dynamic P(%d, %d)\n",n , r)
	result := Dynamic(n, r)
	fmt.Println("Result :",result)
	assert.Equal(t, result, 90)
}

func BenchmarkDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dynamic(n, r)
	}
}
