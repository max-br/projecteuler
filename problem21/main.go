package main

import (
	"fmt"
)

func divisorSum(n int) (sum int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}

func main() {
	var divsum map[int]int = make(map[int]int)
	for i := 1; i < 10000; i++ {
		divsum[i] = divisorSum(i)
	}

	var amicables []int

	for key, val := range divsum {
		if divsum[val] == key && divsum[val] != val {
			amicables = append(amicables, key)
		}
	}

	sum := 0
	for _, val := range amicables {
		fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)
}
