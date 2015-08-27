package main

import (
	"fmt"
)

var triangle = [...]int{75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 20, 4, 82, 47, 65, 19, 1, 23, 75, 3, 34, 88, 2, 77, 73, 7, 63, 67, 99, 65, 4, 28, 6, 16, 70, 92, 41, 41, 26, 56, 83, 40, 80, 70, 33, 41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}

var numRow int = 15

func parentRight(row int, col int) (prow int, pcol int) {
	return row - 1, col
}

func coordToIndex(row int, col int) (index int) {
	return ((row-1)*row)/2 + col
}

func row(n int) []int {
	low := ((n - 1) * n) / 2
	high := (n * (n + 1)) / 2
	return triangle[low:high]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func addRow(n int) {
	row := row(n)
	for i := 0; i < len(row)-1; i++ {
		prow, pcol := parentRight(n, i)
		triangle[coordToIndex(prow, pcol)] += max(row[i], row[i+1])
	}
}

func printTriangle() {
	for i := 1; i <= numRow; i++ {
		fmt.Println(row(i))
	}
}

func main() {
	for i := numRow; i > 0; i-- {
		addRow(i)
	}
	printTriangle()
}
