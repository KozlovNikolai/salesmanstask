package iteration

import (
	"fmt"
	"log"
	"salesmanstask/003/internal/bitree"
	"salesmanstask/003/internal/methods"
	"salesmanstask/003/internal/models"
	"salesmanstask/data"
)

// func Iteration(matrix [][]int, node *bitree.TreeNode) (bitree.Results, bool) {
func Iteration(matrix [][]int, node *bitree.TreeNode) bool {
	models.MxRoot, _ = methods.MatrixConversion(matrix)
	// создаем узлы ветви:
	for {
		mx := Step(matrix)
		if models.Debug {
			fmt.Printf("bitree.BT.Q: %d,\nbitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W: %d\n", bitree.BT.Q, bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W)
		}
		if bitree.BT.Q < bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W {
			fmt.Printf("\nBreak, Q: %d, Tour: %d\n", bitree.BT.Q, bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W)
			return false
		}
		if len(mx) == 2 {
			fmt.Printf("\nBreak, len(mx): %d\n", len(mx))
			// EndingBranch(mx)
			EndingBranch2(mx)
			bitree.BT.Q = models.LbtfRoot
			return true
		}
		matrix = bitree.CloneMx(mx)
	}
}

func Step(mc [][]int) [][]int {
	if models.Debug {
		fmt.Println("______START__ITERATION_________")
	}

	if models.Debug {
		methods.PrintMatrix(mc)
		fmt.Println("      входящая матрица     ^^^")
		fmt.Println("_________________________________________________________________")
	}
	// // получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// // "lower bound of the target function" => lbtf:
	mx5, lbtfNode := methods.MatrixConversion(mc)
	if models.Debug {
		methods.PrintMatrix(mx5)
		fmt.Printf("H_node = %d\n", lbtfNode)
		fmt.Println("      первое приведение     ^^^")
		fmt.Println("_________________________________________________________________")
	}

	// ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextNode := methods.FindCellWithMaxMin(mx5)

	// удаляем найденную ячейку с ее строкой и столбцом:
	rowIdx, colIdx, ok := idxByName(mx5, nextNode.RowName, nextNode.ColName)
	if !ok {
		log.Println("Первый: не могу получить индексы из имени !!!")
	}
	if models.Debug {
		fmt.Printf("RowName: %d, rowIdx: %d\n", nextNode.RowName, rowIdx)
		fmt.Printf("ColName: %d, colIdx: %d\n", nextNode.ColName, colIdx)
	}
	// mx2 := bitree.CloneMx(mx5)
	// mx2[rowIdx][colIdx] = -1
	mx3 := methods.RemoveCellFromMatrixByIndex(mx5, rowIdx, colIdx)

	// rowInfIdx, colInfIdx := methods.FindInfinityCellCoords(mx3)
	// rowInfIdx, colInfIdx := methods.FindInfinityCellCoordsNew(mx3)
	rowInfIdx, colInfIdx, ok := idxByName(mx3, nextNode.ColName, nextNode.RowName)
	if models.Debug {
		// fmt.Printf("RowNameX: %d, colIdx: %d\n", nextNode.RowName, colInfIdx)
		// fmt.Printf("ColNameX: %d, rowIdx: %d\n", nextNode.ColName, rowInfIdx)
	}
	if ok {
		mx3[rowInfIdx][colInfIdx] = data.Inf
	} else {
		log.Println("Второй: не могу получить индексы из имени !!!")
	}

	if models.Debug {
		methods.PrintMatrix(mx3)
		fmt.Println(" маркируем клетку пересечения свободных колонки и строки   ^^^")
		fmt.Println("_________________________________________________________________")
	}

	// // получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// // "lower bound of the target function" => lbtf:
	mx4, lbtfNode := methods.MatrixConversion(mx3)

	if models.Debug {
		methods.PrintMatrix(mx4)
		fmt.Printf("H_node = %d\n", lbtfNode)
		fmt.Println("     второе приведение матрицы     ^^^")
		fmt.Println("_________________________________________________________________")
	}

	var setCurrentRightNode bool
	if models.LbtfRoot+nextNode.MaxSum >= models.LbtfRoot+lbtfNode {
		setCurrentRightNode = true
	}
	fmt.Printf("Change node is Right: %v\n", setCurrentRightNode)
	// bitree.BT.CreateLeftNode(mx2, models.LbtfRoot+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bitree.BT.CreateLeftNode(models.LbtfRoot+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	// bitree.BT.CreateRightNode(mx2, models.LbtfRoot+lbtfNode, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	bitree.BT.CreateRightNode(models.LbtfRoot+lbtfNode, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		models.LbtfRoot = models.LbtfRoot + lbtfNode
		fmt.Printf("Right node is: %v\n", bitree.BT.CurrentNode)
		fmt.Println("Set Right node is current")
	} else {
		models.LbtfRoot = models.LbtfRoot + nextNode.MaxSum
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Left
		fmt.Printf("Left node is: %v\n", bitree.BT.CurrentNode)
		fmt.Println("Set Left node is current")
	}
	if models.Debug {
		fmt.Println("---------STOP----------------")
		fmt.Println()
	}

	return mx4
}

func idxByName(m [][]int, rowName, colName int) (rowIdx, colIdx int, ok bool) {
	for i, v := range m {
		if v[0] == rowName {
			rowIdx = i
			break
		}
	}
	if rowIdx == 0 {
		return 0, 0, false
	}
	for j, v := range m[0] {
		if v == colName {
			colIdx = j
			break
		}
	}
	if colIdx == 0 {
		return 0, 0, false
	}
	return rowIdx, colIdx, true
}

func EndingBranch(mx [][]int) {
	// m := bitree.BT.Result.Back[0].Mxs
	if models.Debug {
		fmt.Println("origin:")
		// methods.PrintMatrix(m)
		methods.PrintMatrix(models.MxRoot)
	}

	for i := 1; i < 2; i++ {
		for j := 1; j < 3; j++ {
			if mx[i][j] != data.Inf {
				var rowName2, colName2 int
				rowName1 := mx[i][0]
				colName1 := mx[0][j]
				if i == 1 && j == 1 {
					rowName2 = mx[i+1][0]
					colName2 = mx[0][j+1]
				} else if i == 2 && j == 2 {
					rowName2 = mx[i-1][0]
					colName2 = mx[0][j-1]
				} else if i == 1 && j == 2 {
					rowName2 = mx[i+1][0]
					colName2 = mx[0][j-1]
				} else if i == 2 && j == 1 {
					rowName2 = mx[i-1][0]
					colName2 = mx[0][j+1]
				}

				// rowIdx, colIdx := idxByName(m, rowName1, colName1)
				rowIdx, colIdx, ok := idxByName(models.MxRoot, rowName1, colName1)
				if !ok {
					log.Println("Ending branch: не могу получить индексы из имени !!!")
				}
				// weight1 := m[rowIdx][colIdx]
				weight1 := models.MxRoot[rowIdx][colIdx]
				// rowIdx, colIdx = idxByName(m, rowName2, colName2)
				rowIdx, colIdx, ok = idxByName(models.MxRoot, rowName2, colName2)
				if !ok {
					log.Println("Ending branch: не могу получить индексы из имени !!!")
				}
				// weight2 := m[rowIdx][colIdx]
				weight2 := models.MxRoot[rowIdx][colIdx]
				if weight1 < weight2 {
					// bitree.BT.CreateRightNode(mx, models.LbtfRoot, rowName1, colName1, true)
					bitree.BT.CreateRightNode(models.LbtfRoot, rowName1, colName1, true)
					bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
					// bitree.BT.CreateLastNode(mx, models.LbtfRoot, rowName2, colName2)
					bitree.BT.CreateLastNode(models.LbtfRoot, rowName2, colName2)
				} else {
					// bitree.BT.CreateRightNode(mx, models.LbtfRoot, rowName2, colName2, true)
					bitree.BT.CreateRightNode(models.LbtfRoot, rowName2, colName2, true)
					bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
					// bitree.BT.CreateLastNode(mx, models.LbtfRoot, rowName1, colName1)
					bitree.BT.CreateLastNode(models.LbtfRoot, rowName1, colName1)
				}
				break
			}
		}
	}
}

func EndingBranch2(mx [][]int) {
	rowIdx, colIdx, ok := idxByName(models.MxRoot, mx[1][0], mx[0][1])
	if !ok {
		log.Println("Ending branch: не могу получить индексы из имени !!!")
	}
	bitree.BT.CreateRightNode(models.LbtfRoot, rowIdx, colIdx, true)
}
