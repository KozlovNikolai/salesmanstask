package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func main() {
	size := 15

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = -1
			} else {
				for {
					matrix[i][j] = rand.Intn(100)
					if matrix[i][j] != 0 {
						break
					}
				}
			}

		}
	}

	PrintMatrix(matrix)

	printMatrixAsArr(matrix)
}
func PrintMatrix(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()

	//fmt.Println()
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
	fmt.Println()
}

func printMatrixAsArr(mx [][]int) {
	for i := 0; i < len(mx); i++ {
		fmt.Print("{")
		for j := 0; j < len(mx[i]); j++ {
			fmt.Printf("%d,", mx[i][j])
		}
		fmt.Print("},")
		fmt.Println()
	}
	fmt.Println()
}
