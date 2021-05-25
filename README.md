# Tugas Besar Strategi Algoritma

> Menganlisa Time Complexity menghitung permutasi menggunakan pendekatan Dynamic Programming dan Brute Force

---

KELOMPOK :

-   ARFAN JADULHAQ - 1301194277
-   MUCHAMMAD ALFI KAROM - 1301190363
-   MUHAMAD ELANG RAMADHAN - 1301190458

## Usage

Pastikan sudah terinstall golang

### Using Go

```sh
go test -v -bench=.
```

### Using Make

```sh
make test
```

Output example:

```sh
=== RUN   TestBruteForce
Brute Force P(10, 2)
Result : 90
--- PASS: TestBruteForce (0.00s)
=== RUN   TestDynamic
Dynamic P(10, 2)
Result : 90
--- PASS: TestDynamic (0.00s)
goos: linux
goarch: amd64
pkg: github.com/arfan21/permutation
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
BenchmarkBruteForce
BenchmarkBruteForce-8           720393622                1.389 ns/op
BenchmarkDynamic
BenchmarkDynamic-8               1862026               649.3 ns/op
PASS
ok      github.com/arfan21/permutation  3.044s
```
