package methods

import (
	"fmt"
	"math"
	"salesmanstask/003/internal/models"
)

// FindCellWithMaxЬшт ищет ячейку из нулевых ячеек, где
// сумма минимальных значений строки и столбца - максимальна
func FindCellWithMaxMin(mx [][]int) models.CellWithMaxMin {
	// определяем размер матрицы
	rowsLen := len(mx)
	colsLen := len(mx[0])

	// создаем список значений нулевых ячеек размером минимум по количеству колонок
	// list := make(map[nullCell]valsOfCell)
	var minRow int
	var minCol int
	result := models.CellWithMaxMin{}
	// идем по строкам исключая строку с заголовками
	for i := 1; i < rowsLen; i++ {
		// идем по элементам строки исключая заголовок строки
		for j := 1; j < colsLen; j++ {
			// если элемент равен нулю, то
			if mx[i][j] == 0 {
				// находим минимальное значение в строке
				minRow = findMinFromArray(mx[i], j)

				// создаем и заполняем массив значениями из колонки
				colArr := make([]int, rowsLen)
				for n := range mx {
					colArr[n] = mx[n][j]
				}
				// находим минимальное значение в колонке
				minCol = findMinFromArray(colArr, i)

				if minCol+minRow > result.MaxSum {
					result = models.CellWithMaxMin{
						RowName: mx[i][0],
						ColName: mx[0][j],
						MaxSum:  minCol + minRow,
					}
				}
			}
		}
	}

	if models.Debug {
		fmt.Printf("Max:%d, (%d,%d)\n", result.MaxSum, result.RowName, result.ColName)
	}

	return result
}

func findMinFromArray(arr []int, exclude int) int {
	min := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if i != exclude && arr[i] >= 0 && arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
