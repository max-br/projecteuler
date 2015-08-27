package main

import (
	"fmt"
)

func collatz(n int) (seq []int) {
	for n != 1 {
		seq = append(seq, n)
		if n%2 == 0 {
			n = n / 2
			continue
		}
		n = 3*n + 1
	}
	seq = append(seq, n)
	return
}

func main() {
	maxlen := 0
	for i := 1; i < 1000000; i++ {
		cseq := collatz(i)
		if len(cseq) > maxlen {
			fmt.Println(i, len(cseq))
			maxlen = len(cseq)
		}
	}
}
