package permutation

func BruteForce(n, r int) int {
	P := 1		// Pendefinisian variabel yang menampung hasil permutasi
	for i := 0; i < r; i++ {	// Looping sebanyak r kali 
		P *= (n - i)	// Operasi perkalian antara P dengan nilai (n-i)
	}
	return P	// Mengembalikan hasil permutasi
}
