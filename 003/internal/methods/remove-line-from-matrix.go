package methods

import (
	"fmt"
	"salesmanstask/003/internal/models"
)

func RemoveCellFromMatrixByIndex(mx [][]int, idxRow int, idxCol int) [][]int {
	// rowInfIdx, colInfIdx := FindInfinityCellCoords(mx, idxRow, idxCol)
	// mx[rowInfIdx][colInfIdx] = -1
	mt := RemoveRowFromMatrixByIndex(mx, idxRow)
	resultMx := RemoveColFromMatrixByIndex(mt, idxCol)
	return resultMx
}

func RemoveRowFromMatrixByIndex(mx [][]int, nameIndex int) [][]int {
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows-1; i++ {
		if i < nameIndex {
			resultMx = append(resultMx, mx[i])
		} else {
			resultMx = append(resultMx, mx[i+1])
		}
	}
	return resultMx
}

func RemoveColFromMatrixByIndex(mx [][]int, nameIndex int) [][]int {
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows; i++ {
		resultMx = append(resultMx, mx[i][:nameIndex])
		resultMx[i] = append(resultMx[i], mx[i][nameIndex+1:]...)
	}
	return resultMx
}

// func FindInfinityCellCoords(mx [][]int, rowDel, colDel int) (rowInfName, colInfName int) {
// 	for i := 0; i < len(mx); i++ {
// 		if mx[i][colDel] == -1 {
// 			rowInfName = i
// 			break
// 		}
// 	}
// 	for j := 0; j < len(mx[0]); j++ {
// 		if mx[rowDel][j] == -1 {
// 			colInfName = j
// 		}
// 	}
// 	return
// }

func FindInfinityCellCoords(mx [][]int) (rowInfName, colInfName int) {
	if models.Debug {
		PrintMatrix(mx)
		fmt.Println("удаляем строку и столбец   ^^^")
		fmt.Println("___________________________________________________")
	}

	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			// fmt.Printf("mx[%d,%d]=%d\n", i, j, mx[i][j])
			if mx[i][j] == -1 {
				break
			}
			if j == len(mx[0])-1 {
				rowInfName = i
			}
		}
	}
	for j := 0; j < len(mx[0]); j++ {
		for i := 1; i < len(mx); i++ {
			if mx[i][j] == -1 {
				break
			}
			if i == len(mx)-1 {
				colInfName = j
			}
		}
	}
	return
}
