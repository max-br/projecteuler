package main

import (
	"fmt"
)

func cntCycle(d int) (cnt int) {
	var res map[int]bool = make(map[int]bool)
	n := 10
	for {
		cnt++
		if n%d == 0 {
			break
		}
		if res[n%d] {
			cnt--
			break
		} else {
			res[n%d] = true
		}
		n = (n % d) * 10
	}
	return
}

func main() {
	max := 0
	for i := 1; i <= 1000; i++ {
		cnt := cntCycle(i)

		fmt.Println(i, cnt)
		if cnt > max {
			max = cnt
		}
	}
}
