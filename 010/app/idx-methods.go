package app

import "fmt"

func IdxByName(m [][]int, rowName, colName int) (rowIdx, colIdx int, ok bool) {
	for i, v := range m {
		if v[0] == rowName {
			rowIdx = i
			break
		}
	}
	if rowIdx == 0 {
		fmt.Printf("не могу получить индекс по имени строки: %d\n", rowName)
		return 0, 0, false
	}
	for j, v := range m[0] {
		if v == colName {
			colIdx = j
			break
		}
	}
	if colIdx == 0 {
		fmt.Printf("не могу получить индекс по имени колонки: %d\n", colName)
		return 0, 0, false
	}
	return rowIdx, colIdx, true
}

func NameByIdx(mx [][]int, rowIdx, colIdx int) (rowName, colName int, ok bool) {
	if rowIdx > len(mx)-1 {
		fmt.Printf("не могу получить имя по индексу строки: %d\n", rowIdx)
		return 0, 0, false
	}
	if colIdx > len(mx[0])-1 {
		fmt.Printf("не могу получить имя по индексу колонки: %d\n", colIdx)
		return 0, 0, false
	}
	return mx[rowIdx][0], mx[0][colIdx], true
}

func ColNameByIdx(mx [][]int, colIdx int) (colName int, ok bool) {
	if colIdx > len(mx[0])-1 {
		fmt.Printf("не могу получить имя по индексу колонки: %d\n", colIdx)
		return 0, false
	}
	return mx[0][colIdx], true
}
