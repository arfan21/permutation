package permutation

func BruteForce(n, r int) int {
	P := 1
	for i := 0; i < r; i++ {
		P *= (n - i)
	}
	return P
}
