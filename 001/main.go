package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

var matrix = [][]int{
	{0, 0, -1},
	{-1, 3, 5},
	{2, -1, 0},
}

// var matrix = [][]int{
// 	{-1, 0, 0, 2},
// 	{0, -1, 3, 5},
// 	{2, 12, 10, -1},
// 	{0, 2, 0, 0},
// }

// var matrix = [][]int{
// 	{-1, 5, 16, 14},
// 	{13, -1, 6, 9},
// 	{10, 12, -1, 11},
// 	{8, 15, 7, -1},
// }

// var matrix = [][]int{
// 	{-1, 0, 0, 2, 3},
// 	{6, -1, 9, 11, -1},
// 	{4, 14, -1, 11, 0},
// 	{2, 12, 10, -1, 0},
// 	{0, 2, 0, 0, -1},
// }

func main() {
	printMatrix(matrix)
	fmt.Println()
	mc := matrixConversion(matrix)
	printMatrixColor(mc)
	//##########################################

}

func matrixConversion(mx [][]int) [][]int {
	rc := rowsConversion(mx)
	cc := columnsConversion(rc)
	return cc
}

func rowsConversion(mx [][]int) [][]int {
	rows := len(mx)
	cols := len(mx[0])

	resMx := make([][]int, rows)
	for i := 0; i < rows; i++ {
		resMx[i] = make([]int, cols+1)
	}

	for i := 0; i < rows; i++ {
		min := math.MaxInt
		for j := 0; j < cols; j++ {
			if mx[i][j] >= 0 {
				if mx[i][j] < min {
					min = mx[i][j]
				}
			}
		}
		//fmt.Printf("min: %d\n", min)
		for j := 0; j < cols; j++ {
			if mx[i][j] >= 0 {
				resMx[i][j] = mx[i][j] - min
			} else {
				resMx[i][j] = mx[i][j]
			}
		}
		resMx[i][cols] = min
	}
	return resMx
}

func columnsConversion(mx [][]int) [][]int {
	rows := len(mx)
	cols := len(mx[0])

	resMx := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		resMx[i] = make([]int, cols)
	}

	for j := 0; j < cols-1; j++ {
		min := math.MaxInt
		for i := 0; i < rows; i++ {
			if mx[i][j] >= 0 {
				if mx[i][j] < min {
					min = mx[i][j]
				}
			}
		}
		for i := 0; i < rows; i++ {
			if mx[i][j] >= 0 {
				resMx[i][j] = mx[i][j] - min
			} else {
				resMx[i][j] = mx[i][j]
			}
		}
		resMx[rows][j] = min
	}
	var sum int
	for i := 0; i < rows; i++ {
		resMx[i][cols-1] = mx[i][cols-1]
		sum += mx[i][cols-1]
	}
	for j := 0; j < cols-1; j++ {
		sum += resMx[rows][j]
	}
	resMx[rows][cols-1] = sum
	return resMx
}

func printMatrixColor(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[i]); j++ {
			if mx[i][j] == -1 {
				fmt.Printf(black("%-4d"), mx[i][j])
			} else {
				if i == len(mx)-1 && j == len(mx[i])-1 {
					fmt.Printf(white("%-4d"), mx[i][j])
				} else if i == len(mx)-1 || j == len(mx[i])-1 {
					fmt.Printf(green("%-4d"), mx[i][j])
				} else if mx[i][j] == 0 {
					fmt.Printf(red("%-4d"), mx[i][j])
				} else {
					fmt.Printf(yel("%-4d"), mx[i][j])
				}
			}

			// fmt.Printf("%-4d", mx[i][j])
		}
		fmt.Println()
	}
}
func printMatrix(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[i]); j++ {
			if mx[i][j] == -1 {
				fmt.Printf(black("%-4d"), mx[i][j])
			} else {
				if mx[i][j] == 0 {
					fmt.Printf(red("%-4d"), mx[i][j])
				} else {
					fmt.Printf(yel("%-4d"), mx[i][j])
				}
			}

			// fmt.Printf("%-4d", mx[i][j])
		}
		fmt.Println()
	}
}
