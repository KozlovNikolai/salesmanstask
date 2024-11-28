package main

/*
При запросе программы - скопировать и вставить в терминал:
5 5
1 2 3 4 5
6 1 8 9 0
1 2 3 4 5
6 7 8 9 0
1 2 3 4 5

*/
import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Print(`Ищем путь от верхнего левого угла к нижнему правому за максимальную стоимость.

Введи размер матрицы и ее данные через пробел, например так:
5 5
1 2 3 4 5
6 1 8 9 0
1 2 3 4 5
6 7 8 9 0
1 2 3 4 5

Давай, начинай вводить:
`)
	var n int
	var m int
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Printf("\nmatrix:\n")
	print(matrix)

	// создаем динамическую матрицу
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// создаем матрицу пути
	pm := make([][][2]int, n)
	for i := 0; i < n; i++ {
		pm[i] = make([][2]int, m)
	}

	// заполняем первый элемент
	dp[0][0] = matrix[0][0]
	pm[0][0] = [2]int{0, 0}

	// заполняем верхнюю строку
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
		pm[0][j] = [2]int{0, j - 1}
	}

	// заполняем левый столбец
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
		pm[i][0] = [2]int{i - 1, 0}
	}

	// ищем максимумы
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
				pm[i][j] = [2]int{i - 1, j}
			} else {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
				pm[i][j] = [2]int{i, j - 1}
			}
		}
	}
	fmt.Printf("\nстоимость по нарастающей:\n")
	print(dp)
	fmt.Printf("максимальная стоимость пути: %d\n", dp[n-1][m-1])

	// восстанавливаем путь
	i, j := n-1, m-1
	path := [][2]int{{i, j}}
	for pm[i][j] != [2]int{0, 0} {
		path = append([][2]int{pm[i][j]}, path...)
		i, j = pm[i][j][0], pm[i][j][1]
	}
	path = append([][2]int{{0, 0}}, path...)
	fmt.Printf("\npath: %+v\n", path)

	printPath(matrix, path)
}

func print(t [][]int) {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			fmt.Printf("%-3d", t[i][j])
		}
		fmt.Println()
	}
}

func printPath(matrix [][]int, path [][2]int) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("\npath:\n")
	_ = path
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if compare(path, [2]int{i, j}) {
				fmt.Printf(red("%-3d"), matrix[i][j])
			} else {
				fmt.Printf("%-3d", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

func compare(path [][2]int, point [2]int) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == point {
			return true
		}
	}
	return false
}
