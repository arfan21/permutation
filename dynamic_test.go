package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamic(t *testing.T) {
	result := Dynamic(10, 2)
	fmt.Println(result)
	assert.Equal(t, result, 90)
}

func BenchmarkDynamic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Dynamic(10, 2)
	}
}
