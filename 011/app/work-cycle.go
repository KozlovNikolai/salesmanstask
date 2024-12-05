package app

import "fmt"

func Run(s *Store) {
	for !s.IsSolved {
		Iteration(s)
	}
}

func Iteration(s *Store) {
	// текущий узел
	node := s.Tree[s.CurrentNodeID]
	mx := node.MX
	if Debug {
		fmt.Printf("Текущая матрица - id:%d, (%d,%d), w:%d\n", node.ID, node.Out, node.In, node.W)
		PrintMatrix(mx)
	}

	if node.Out == 0 || node.In == 0 {
		rowName := 1
		for colName := 2; colName < len(mx[0]); colName++ {
			mxr := Reduce(mx, rowName, colName)
			MarkInfinityCell(mxr, rowName, colName, s)
			h := GetH(mxr, rowName, colName, s)
			s.AddNode(mxr, rowName, colName, h)
			s.InsertLeaf()
		}
	} else {
		rowName := mx[1][0]
		for colIdx := 1; colIdx < len(mx[0]); colIdx++ {
			colName, _ := ColNameByIdx(mx, colIdx)

			if mx[1][colIdx] == Inf {
				continue
			}
			if Debug {
				PrintMatrix(mx)
				fmt.Printf("Новая матрица ID:%d\n", s.NextID)
				fmt.Printf("убираем строку/колонку: %d/%d\n", rowName, colName)
			}

			mxr := Reduce(mx, rowName, colName)
			if Debug {
				PrintMatrix(mxr)
				fmt.Printf("Маркируем бесконечную дугу\n")
			}

			_, ok := MarkInfinityCell(mxr, rowName, colName, s)
			if ok {
				fmt.Println("*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_")
				fmt.Println("*_^_*_^_*_^_*_^_*_^_*_^   РЕШЕНО   _*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_")

				s.IsSolved = true
				return
			}
			if Debug {
				PrintMatrix(mxr)
			}

			h := GetH(mxr, rowName, colName, s)

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

func MarkInfinityCell(mx [][]int, rowName, colName int, s *Store) ([][]int, bool) {
	// если текущий узел - корень дерева, то просто меняем имена строки и столбца между собой
	// в корневой матрице имена столбцов всегда совпадают с индексами массива
	parentNodeID := s.Tree[s.CurrentNodeID].ParentID
	if s.CurrentNodeID == 0 {
		rowIdx, colIdx, _ := IdxByName(mx, colName, rowName)
		mx[rowIdx][colIdx] = Inf
		return mx, false
	}

	// создаем список лучей всех вышестоящих узлов, включая текущий
	list := map[int]int{
		rowName: colName,
	}
	list[s.Tree[s.CurrentNodeID].Out] = s.Tree[s.CurrentNodeID].In
	for parentNodeID != 0 {
		list[s.Tree[parentNodeID].Out] = s.Tree[parentNodeID].In
		parentNodeID = s.Tree[parentNodeID].ParentID
	}
	if Debug {
		PrintMap(list)
	}

	for j := 1; j < len(mx[0]); j++ {
		name := mx[0][j]
		count := 0
		for {
			count++
			value, ok := list[name]
			if !ok {
				break
			}
			for i := 1; i < len(mx); i++ {
				if mx[i][0] == value {
					if count == (len(s.Tree[0].MX) - 2) {
						list[mx[1][0]] = mx[0][1]

						return mx, true
					}
					mx[i][j] = Inf
				}
			}
			name = value
		}
	}

	return mx, false
}

func PrintMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("key:%d, val:%d\n", k, v)
	}
}

func MapToArray(m map[int]int, start int) []int {
	result := make([]int, len(m))
	cnt := 0
	for cnt < len(m) {
		result[cnt] = m[start]
		start = m[start]
		cnt++
	}
	return result
}
