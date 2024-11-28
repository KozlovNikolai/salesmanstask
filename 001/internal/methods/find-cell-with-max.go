package methods

import (
	"fmt"
	"math"
	"salesmanstask/001/internal/models"
)

type nullCell struct {
	row int
	col int
}
type valsOfCell struct {
	minOfRows int
	minOfCols int
}

// FindCellWithMaxЬшт ищет ячейку из нулевых ячеек, где
// сумма минимальных значений строки и столбца - максимальна
func FindCellWithMaxMin(mx [][]int, names *models.NamesOfIndexes) models.CellWithMaxMin {
	// определяем размер матрицы
	rowsLen := len(mx)
	colsLen := len(mx[0])

	// создаем список значений нулевых ячеек размером минимум по количеству колонок
	list := make(map[nullCell]valsOfCell)
	var minRow int
	var minCol int

	for i := 0; i < rowsLen; i++ {
		for j := 0; j < colsLen; j++ {
			if mx[i][j] == 0 {
				minRow = findMinFromArray(mx[i], j)
				colArr := make([]int, rowsLen)
				for n := range mx {
					colArr[n] = mx[n][j]
				}
				minCol = findMinFromArray(colArr, i)

				// записываем результаты проходов через нулевую ячейку в список:
				list[nullCell{i, j}] = valsOfCell{
					minOfRows: minRow,
					minOfCols: minCol,
				}
			}
		}

	}

	result := models.CellWithMaxMin{}

	for i, v := range list {
		fmt.Printf("[%+v], %+v\n", i, v)
		if v.minOfCols+v.minOfRows > result.MaxSum {
			result = models.CellWithMaxMin{
				RowName: names.NamesOfRows[i.row],
				ColName: names.NamesOfCols[i.col],
				MaxSum:  v.minOfCols + v.minOfRows,
			}
		}
	}
	fmt.Printf("Max:%d, (%d,%d)\n", result.MaxSum, result.RowName, result.ColName)
	return result
}

func findMinFromArray(arr []int, exclude int) int {
	min := math.MaxInt
	for i, v := range arr {
		if i != exclude && v >= 0 && v < min {
			min = v
		}
	}

	// fmt.Printf("arr: %+v\n", arr)
	// fmt.Printf("excl: %d\n", exclude)
	// fmt.Printf("min: %d\n\n", min)

	return min
}
