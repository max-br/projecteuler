package main

import (
	"fmt"
)

func pow2(n int) (p int) {
	p = 2
	for n > 0 {
		p *= 2
		n--
	}
	return
}

func digitSum(n int) (s int) {
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return
}

func main() {
	for i := 0; i < 50; i++ {
		pow := pow2(i)
		fmt.Println(i, pow, digitSum(pow))
	}
}
