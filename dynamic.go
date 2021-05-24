package permutation

import (
	"math"
)

func Dynamic(n, r int) int {
	P := make([][]int, n+2)

	for i := range P {

		P[i] = make([]int, r+2)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= int(math.Min(float64(i), float64(r))); j++ {
			if j == 0 {
				P[i][j] = 1
			} else {
				P[i][j] = P[i-1][j] + (j * P[i-1][j-1])
			}
			P[i][j+1] = 0
		}
	}

	return P[n][r]
}
