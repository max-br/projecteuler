package main

import (
	"fmt"
)

func spiralSequence(n int) (seq []int) {
	center := 1
	seq = append(seq, center)
	distance := 2
	num := center
	for distance < n {
		for j := 0; j < 4; j++ {
			num = num + distance
			seq = append(seq, num)
		}
		distance += 2
	}
	return
}

func sumSlice(s []int) (sum int) {
	for _, val := range s {
		sum += val
	}
	return
}

func main() {
	seq := spiralSequence(1001)
	fmt.Println(sumSlice(seq))
}
