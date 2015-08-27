package main

import (
	"fmt"
)

func divCntNaive(n int) (cnt int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return
}

func tNoDivisorCount(n int) (tno int, cnt int) {
	tno = (n * (n + 1)) / 2
	divn := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divn[i] = true
		}
	}
	divsn := make(map[int]bool)
	for i := 1; i <= n+1; i++ {
		if (n+1)%i == 0 {
			divsn[i] = true
		}
	}
	alldivs := make(map[int]bool)
	for keyn, _ := range divn {
		for keysn, _ := range divsn {
			alldivs[keyn*keysn] = true
		}
	}
	for key, val := range alldivs {
		if val {
			if tno%key == 0 {
				cnt++
			}
		}
	}
	return
}

func main() {
	for i := 1; ; i++ {
		tno, cnt := tNoDivisorCount(i)
		fmt.Println(i, tno, cnt)
		if cnt >= 500 {
			break
		}
	}
}
