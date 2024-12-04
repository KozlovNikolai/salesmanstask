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
	fmt.Printf("Текущая матрица - id:%d, (%d,%d), w:%d\n", node.ID, node.Out, node.In, node.W)
	PrintMatrix(mx)
	if node.Out == 0 || node.In == 0 {
		rowName := 1
		for colName := 2; colName < len(mx[0]); colName++ {
			mxr := Reduce(mx, rowName, colName)
			MarkInfinityCell(mxr, rowName, colName, s)
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
			// if rowName == colName {
			// 	continue
			// }
			if mx[1][colIdx] == Inf {
				continue
			}
			PrintMatrix(mx)
			fmt.Printf("Новая матрица ID:%d\n", s.NextID)
			fmt.Printf("убираем строку/колонку: %d/%d\n", rowName, colName)
			mxr := Reduce(mx, rowName, colName)
			PrintMatrix(mxr)

			fmt.Printf("Маркируем бесконечную дугу\n")
			_, ok := MarkInfinityCell(mxr, rowName, colName, s)
			if ok {
				fmt.Println("*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_")
				fmt.Println("*_^_*_^_*_^_*_^_*_^_*_^   РЕШЕНО   _*_^_*_^_*_^_*_^_*_^_*_^_*_^_*_^_")
			}
			PrintMatrix(mxr)
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

// func MarkInfinityCell(mx [][]int, rowName, colName, rowNameP, colNameP int) [][]int {
// 	var rowIdx, colIdx int
// 	// if rowNameP == colName {
// 	if rowName == colNameP {
// 		// rowIdx, colIdx, _ = IdxByName(mx, colNameP, rowName)
// 		rowIdx, colIdx, _ = IdxByName(mx, colName, rowNameP)
// 	} else {
// 		rowIdx, colIdx, _ = IdxByName(mx, colName, rowName)
// 	}
// 	mx[rowIdx][colIdx] = Inf
// 	return mx
// }

// func MarkInfinityCell(mx [][]int, rowName, colName, rowNameP, colNameP int) [][]int {
// 	var rowIdx, colIdx int
// 	if rowNameP == colName {
// 		rowIdx, colIdx, _ = IdxByName(mx, colNameP, rowName)
// 	} else if rowName == colNameP {
// 		rowIdx, colIdx, _ = IdxByName(mx, colName, rowNameP)
// 	} else {
// 		rowIdx, colIdx, _ = IdxByName(mx, colName, rowName)
// 	}
// 	mx[rowIdx][colIdx] = Inf
// 	return mx
// }

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
	PrintMapSort(list)

	// создаем словарь с количеством вхождений узлов в тур.
	// количество вхождений может быть 1 или 2 (других быть не может)
	count := 0 // счетчик парных вхождений

	dict := make(map[int]struct {
		cnt int
		dir string
	})
	for key, val := range list {
		v, ok := dict[key]
		if ok {
			count++
		}
		dict[key] = struct {
			cnt int
			dir string
		}{v.cnt + 1, "out"}
		v, ok = dict[val]
		if ok {
			count++
		}
		dict[val] = struct {
			cnt int
			dir string
		}{v.cnt + 1, "in"}
	}
	// если счетчик парных вхождений равен длине списка лучей, значит нашли ответ, TRUE
	if count == len(list) {
		return mx, true
	}

	// если входные параметры имеют по одному вхождению, то просто переворачиваем их
	// и преобразуем в индексы текущей таблицы
	var rowIdx, colIdx int
	if dict[rowName].cnt == 1 && dict[colName].cnt == 1 {
		rowIdx, colIdx, _ = IdxByName(mx, colName, rowName)
		mx[rowIdx][colIdx] = Inf
		return mx, false
	}

	// var row, col, temp int
	for k, v := range dict {
		fmt.Printf("key:%d, val:%+v\n", k, v)
	}

	var row, col int
	for item, value := range dict {
		if value.cnt == 1 && value.dir == "out" {
			col = item
		} else if value.cnt == 1 && value.dir == "in" {
			row = item
		}
	}

	rowIdx, colIdx, _ = IdxByName(mx, row, col)
	mx[rowIdx][colIdx] = Inf
	return mx, false
}

func PrintMapSort(m map[int]int) {
	for k, v := range m {
		fmt.Printf("key:%d, val:%d\n", k, v)
	}
}
