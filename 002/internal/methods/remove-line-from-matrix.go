package methods

func RemoveCellFromMatrixByIndex(mx [][]int, nameRow int, nameCol int) [][]int {
	if nameRow < 0 {
		nameRow = -nameRow
	}
	if nameCol < 0 {
		nameCol = -nameCol
	}

	rowInf, colInf := findInfinityCellCoords(mx, nameRow, nameCol)
	mx[rowInf][colInf] = -1
	mt := RemoveRowFromMatrixByIndex(mx, nameRow)
	resultMx := RemoveColFromMatrixByIndex(mt, nameCol)
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
