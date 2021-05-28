package permutation

import (
	"math"
)

func Dynamic(n, r int) int {
	// Pendefinisian array 2 dimensi untuk menyimpan nilai hasil permutasi
	P := make([][]int, n+2)

	for i := range P {

		P[i] = make([]int, r+2)
	}

	// Looping untuk array dimensi pertama, dilakukan sebanyak n+1 kali
	for i := 0; i <= n; i++ {
		// Looping untuk array dimensi kedua
		// Untuk setiap j, dilakukan looping sebanyak nilai minimum antara i dan r
		// Hal ini karena untuk setiap j > i maka hasil permutasi akan selalu 0
		for j := 0; j <= int(math.Min(float64(i), float64(r))); j++ {
			if j == 0 { // Basis ketika P(i,0)
				P[i][j] = 1
			} else { // Menghitung hasil P(i,j) menggunakan nilai yang telah disimpan pada array
				P[i][j] = P[i-1][j] + (j * P[i-1][j-1])
			}
			// Untuk setiap j > i maka hasil permutasi akan selalu 0
			P[i][j+1] = 0
		}
	}
	// Mengembalikan hasil P(n,r)
	return P[n][r]
}
