package main

import (
	"fmt"
	"github.com/max-br/projecteuler/lib"
)

var primes []bool = lib.Primesieve(1000000)

func isPrime(n int) bool {
	return primes[n]
}

func evalExpr(a int, b int) (primecnt int) {
	n := 0
	candidate := n*n + a*n + b
	for candidate > 0 && isPrime(candidate) {
		primecnt++
		n++
		candidate = n*n + a*n + b
	}
	return
}

func testBValues(a int) (bmax int, maxcnt int) {
	b := -999
	for b < 1000 {
		cnt := evalExpr(a, b)
		if cnt > maxcnt {
			bmax = b
			maxcnt = cnt
		}
		b++
	}
	return
}

func main() {
	fmt.Println()
	amax, bmax, maxcnt := 0, 0, 0
	a := -999
	for a < 1000 {
		b, cnt := testBValues(a)
		if cnt > maxcnt {
			amax = a
			bmax = b
			maxcnt = cnt
		}
		a++
		fmt.Println(a, amax, bmax, maxcnt)
	}
	fmt.Println(amax * bmax)
}
