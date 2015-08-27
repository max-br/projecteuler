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
	var abundantnumber []int
	for i := 1; i <= 28123; i++ {
		if divisorSum(i) > i {
			abundantnumber = append(abundantnumber, i)
		}
	}

	fmt.Println()

	var numbers [28124]bool

	for i, ano := range abundantnumber {
		for _, ano2 := range abundantnumber[i:] {
			if ano+ano2 > 28123 {
				break
			}
			numbers[ano+ano2] = true
		}
	}

	sum := 0
	for i, val := range numbers {
		if !val {
			fmt.Println(i, val)
			sum += i
		}
	}
	fmt.Println(sum)
}
