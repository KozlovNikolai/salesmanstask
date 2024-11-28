package methods

import (
	"fmt"
	"salesmanstask/001/internal/models"
)

func RemoveCellFromMatrixByIndex(mx [][]int, indexRow int, indexCol int, names *models.NamesOfIndexes) ([][]int, *models.NamesOfIndexes) {
	rowInf, colInf := findInfinityCellCoords(mx, indexRow, indexCol)
	mx[rowInf][colInf] = -1
	mt, nm := RemoveRowFromMatrixByIndex(mx, indexRow, names)
	resultMx, nms := RemoveColFromMatrixByIndex(mt, indexCol, nm)
	return resultMx, nms
}

func RemoveRowFromMatrixByIndex(mx [][]int, index int, names *models.NamesOfIndexes) ([][]int, *models.NamesOfIndexes) {
	fmt.Printf("names before delete %+v\n", names)
	name, names := names.RemoveRowByIndex(index)
	fmt.Printf("Deleted row idx: %d, name: %d\n", index, name)
	fmt.Printf("names after delete %+v\n", names)
	//_ = name
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows-1; i++ {
		if i < index {
			resultMx = append(resultMx, mx[i])
		} else {
			resultMx = append(resultMx, mx[i+1])
		}
	}
	return resultMx, names
}

func RemoveColFromMatrixByIndex(mx [][]int, index int, names *models.NamesOfIndexes) ([][]int, *models.NamesOfIndexes) {
	fmt.Printf("names before delete %+v\n", names)
	name, names := names.RemoveColByIndex(index)
	fmt.Printf("Deleted column idx: %d, name: %d\n", index, name)
	fmt.Printf("names after delete %+v\n", names)
	//_ = name
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows; i++ {
		resultMx = append(resultMx, mx[i][:index])
		resultMx[i] = append(resultMx[i], mx[i][index+1:]...)
	}
	return resultMx, names
}

func findInfinityCellCoords(mx [][]int, rowDel, colDel int) (rowInfName, colInfName int) {
	for i := 0; i < len(mx); i++ {
		if mx[i][colDel] == -1 {
			rowInfName = i
			break
		}
	}
	for j := 0; j < len(mx[0]); j++ {
		if mx[rowDel][j] == -1 {
			colInfName = j
		}
	}
	return
}
