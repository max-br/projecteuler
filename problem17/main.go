package main

import (
	"fmt"
)

func digitLetterCnt(n int) int {
	if n == 1 || n == 2 || n == 6 {
		return 3
	}
	if n == 4 || n == 5 || n == 9 {
		return 4
	}
	if n == 3 || n == 7 || n == 8 {
		return 5
	}
	return 0
}

func cntLetters(n int) (cnt int) {
	if n/1000 == 1 {
		return 3 + 8 // one thousand
	}
	if n/100 >= 1 {
		cnt += 7 // hundred
		cnt += digitLetterCnt(n / 100)
		n = n % 100
	}

	return
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(i, cntLetters(i))
	}

}
