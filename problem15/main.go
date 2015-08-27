package main

import (
	"fmt"
)

const limit int = 20

var board [limit * limit]int

func noFields(n int) (ret int) {
	if n == 1 {
		return 1
	}
	return n*n - noFields(n-1)
}

var cnt int = 0

func calcSteps(rightsteps int, leftsteps int) {
	if rightsteps == 0 || leftsteps == 0 {
		cnt++
		return
	}
	calcSteps(rightsteps-1, leftsteps)
	calcSteps(rightsteps, leftsteps-1)
}

func main() {
	cnt = 0
	calcSteps(20, 20)
	fmt.Println(cnt)
}
