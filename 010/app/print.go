package app

import (
	"fmt"

	"github.com/fatih/color"
)

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
			} else if mx[i][j] == Inf {
				fmt.Printf(black("\t%s"), "inf")
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
