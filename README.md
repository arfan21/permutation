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
make testAll
```

### Test Brute Force Only

```sh
make testBruteForce
```

### Test Dynamic Only

```sh
make testDynamic
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
BenchmarkBruteForceN10R2
BenchmarkBruteForceN10R2-8      748677180                1.441 ns/op
BenchmarkBruteForceN5R4
BenchmarkBruteForceN5R4-8       391819078                2.776 ns/op
BenchmarkBruteForceN20R10
BenchmarkBruteForceN20R10-8     318623853                3.755 ns/op
BenchmarkDynamicN10R2
BenchmarkDynamicN10R2-8          1894484               692.1 ns/op
BenchmarkDynamicN5R4
BenchmarkDynamicN5R4-8           2467027               507.0 ns/op
BenchmarkDynamicN20R10
BenchmarkDynamicN20R10-8          493983              2819 ns/op
PASS
ok      github.com/arfan21/permutation  10.422s
```
