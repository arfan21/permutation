package main

import (
	"fmt"

	"github.com/arfan21/permutation/permutation"
)

func main() {
	fmt.Println("Masukkan angka untuk P(n,r). contoh : 8 2")
	fmt.Print("Input -> ")
	var n, r int
	fmt.Scanln(&n, &r)
	fmt.Println("Hasil menggunakan Dynamic :", permutation.Dynamic(n, r))
	fmt.Println("Hasil menggunakan Brute Force :", permutation.BruteForce(n, r))
}
