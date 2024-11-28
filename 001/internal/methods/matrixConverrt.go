package methods

import (
	"math"
	"salesmanstask/001/internal/models"
)

func MatrixConversion(mx [][]int, names *models.NamesOfIndexes) ([][]int, int) {
	rc := rowsConversion(mx)
	cc, sum := columnsConversion(rc)
	//PrintMatrixColor(cc, names)
	//fmt.Println()

	matrix := cutFromConversionMatrix(cc)
	return matrix, sum
}

func cutFromConversionMatrix(mx [][]int) [][]int {
	rows := len(mx)
	cols := len(mx[0])
	for i := range mx {
		mx[i] = mx[i][:cols-1]
	}
	mx = mx[:rows-1]
	return mx
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

func columnsConversion(mx [][]int) ([][]int, int) {
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
	return resMx, sum
}
