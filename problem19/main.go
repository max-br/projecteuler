package main

import (
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

func addMonth(date *Date, n int) {
	date.month += n
	if date.month > 12 {
		date.year++
		date.month = date.month % 12
	}
}

func addDay(date *Date, n int) {
	date.day += n
	if date.month == 2 {
		if date.year%4 == 0 {
			if date.day > 29 {
				addMonth(date, 1)
				date.day = date.day % 29
			}
			return
		}
		if date.day > 28 {
			addMonth(date, 1)
			date.day = date.day % 28
		}
		return
	}
	if date.month == 4 || date.month == 6 || date.month == 9 || date.month == 11 {
		if date.day > 30 {
			addMonth(date, 1)
			date.day = date.day % 30
		}
		return
	}
	if date.day > 31 {
		addMonth(date, 1)
		date.day = date.day % 31
	}
}

func main() {
	var date Date = Date{1901, 1, 6}
	cnt := 0
	for date.year < 2001 {
		if date.day == 1 {
			cnt++
			fmt.Println(date, cnt)
		}
		addDay(&date, 7)
	}
}
