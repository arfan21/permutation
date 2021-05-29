testAll:
	go test -v -bench=.

testBruteForce:
	go test constant.go brute_force.go  brute_force_test.go -v -bench=.

testDynamic:
	go test constant.go dynamic.go  dynamic_test.go -v -bench=.