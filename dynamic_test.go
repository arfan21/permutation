package permutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamic(t *testing.T) {
	result := Dynamic(5, 4)
	fmt.Println(result)
	assert.Equal(t, result, 120)
}

func BenchmarkDynamic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Dynamic(5, 4)
	}
}
