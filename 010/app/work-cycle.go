package app

import "fmt"

func Run(s *Store) {
	//for {
	Iteration(s)
	//}
}

func Iteration(s *Store) {
	// текущий узел
	node := s.Tree[s.CurrentNodeID]
	mx := node.MX
	fmt.Printf("Текущая матрица: \n")
	PrintMatrix(mx)
	if node.Out == 0 || node.In == 0 {
		rowName := 1
		for colName := 2; colName < len(mx[0]); colName++ {
			mxr := Reduce(mx, rowName, colName)
			MarkInfinityCell(mxr, rowName, colName, node.Out, node.In)
			h := GetH(mxr, rowName, colName, s)
			// fmt.Printf("H: %d\n", h)
			// PrintMatrix(mxr)
			s.AddNode(mxr, rowName, colName, h)
			s.InsertLeaf()
		}
	} else {
		rowName := mx[1][0]
		for colIdx := 1; colIdx < len(mx[0]); colIdx++ {
			colName, _ := ColNameByIdx(mx, colIdx)
			if rowName == colName {
				continue
			}
			mxr := Reduce(mx, rowName, colName)
			MarkInfinityCell(mxr, rowName, colName, node.Out, node.In)
			h := GetH(mxr, rowName, colName, s)
			// fmt.Printf("H: %d\n", h)
			// PrintMatrix(mxr)
			s.AddNode(mxr, rowName, colName, h)
			s.InsertLeaf()
		}
	}

	s.RemoveLeaf()
}

func Reduce(mx [][]int, rowName, colName int) [][]int {
	mxClone := CloneMx(mx)
	idxRow, idxCol, ok := IdxByName(mxClone, rowName, colName)
	if !ok {
		return nil
	}
	mxRemRow := RemoveRowFromMatrixByIndex(mxClone, idxRow)
	mxRemCol := RemoveColFromMatrixByIndex(mxRemRow, idxCol)
	//	resultMx := MarkInfinityCell(mxRemCol, rowName, colName)

	return mxRemCol
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

func MarkInfinityCell(mx [][]int, rowName, colName, rowNameP, colNameP int) [][]int {
	var rowIdx, colIdx int
	if rowNameP == colName {
		rowIdx, colIdx, _ = IdxByName(mx, colNameP, rowName)
	} else {
		rowIdx, colIdx, _ = IdxByName(mx, colName, rowName)
	}
	mx[rowIdx][colIdx] = Inf
	return mx
}
