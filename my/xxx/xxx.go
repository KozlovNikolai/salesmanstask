package main

import (
	"fmt"

	"github.com/fatih/color"
)

var matrix = [][]int{
	{-1, 1, 2, 3, 4},
	{14, -1, 15, 16, 5},
	{13, 20, -1, 17, 6},
	{12, 19, 18, -1, 7},
	{11, 10, 9, 8, -1},
}

func main() {
	PrintMatrix(matrix)

	mt := CloneMx(matrix)

	PrintMatrix(mt)

	mt[2][2] = 111
	PrintMatrix(mt)
	PrintMatrix(matrix)
}

func CloneMx(mx [][]int) [][]int {
	lenRows := len(mx)
	lenCols := len(mx[0])
	mxClone := make([][]int, lenRows)
	for i := range mxClone {
		mxClone[i] = make([]int, lenCols)
	}
	for i := 0; i < lenRows; i++ {
		for j := 0; j < lenCols; j++ {
			mxClone[i][j] = mx[i][j]
		}
	}
	return mxClone
}
func PrintMatrix(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()

	fmt.Println()
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[i]); j++ {
			if i == 0 || j == 0 {
				fmt.Printf(green("\t%d"), mx[i][j])
			} else if mx[i][j] == -1 {
				fmt.Printf(black("\t%d"), mx[i][j])
			} else {
				if mx[i][j] == 0 {
					fmt.Printf(red("\t%d"), mx[i][j])
				} else {
					fmt.Printf(yel("\t%d"), mx[i][j])
				}
			}
		}
		fmt.Println()
	}
}
