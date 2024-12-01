package iteration

import (
	"fmt"
	"salesmanstask/003/internal/bitree"
	"salesmanstask/003/internal/methods"
	"salesmanstask/003/internal/models"
)

// func Iteration(matrix [][]int, node *bitree.TreeNode) (bitree.Results, bool) {
func Iteration(matrix [][]int, node *bitree.TreeNode) bool {
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
		if len(mx) == 3 {
			fmt.Printf("\nBreak, len(mx): %d\n", len(mx))
			EndingBranch(mx)
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
	rowIdx, colIdx := idxByName(mx5, nextNode.RowName, nextNode.ColName)
	if models.Debug {
		fmt.Printf("RowName: %d, rowIdx: %d\n", nextNode.RowName, rowIdx)
		fmt.Printf("ColName: %d, colIdx: %d\n", nextNode.ColName, colIdx)
	}
	mx2 := bitree.CloneMx(mx5)
	mx2[rowIdx][colIdx] = -1
	mx3 := methods.RemoveCellFromMatrixByIndex(mx5, rowIdx, colIdx)

	rowInfIdx, colInfIdx := methods.FindInfinityCellCoords(mx3)
	mx3[rowInfIdx][colInfIdx] = -1

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

	bitree.BT.CreateLeftNode(mx2, models.LbtfRoot+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bitree.BT.CreateRightNode(mx2, models.LbtfRoot+lbtfNode, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		models.LbtfRoot = models.LbtfRoot + lbtfNode
	} else {
		models.LbtfRoot = models.LbtfRoot + nextNode.MaxSum
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Left
	}
	if models.Debug {
		fmt.Println("---------STOP----------------")
		fmt.Println()
	}

	return mx4
}

func idxByName(m [][]int, rowName, colName int) (rowIdx, colIdx int) {
	for i, v := range m {
		if v[0] == rowName {
			rowIdx = i
			break
		}
	}
	for j, v := range m[0] {
		if v == colName {
			colIdx = j
			break
		}
	}
	return
}

func EndingBranch(mx [][]int) {
	m := bitree.BT.Result.Back[0].Mxs
	if models.Debug {
		fmt.Println("origin:")
		methods.PrintMatrix(m)
	}

	for i := 1; i < 2; i++ {
		for j := 1; j < 3; j++ {
			if mx[i][j] != -1 {
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

				rowIdx, colIdx := idxByName(m, rowName1, colName1)
				weight1 := m[rowIdx][colIdx]
				rowIdx, colIdx = idxByName(m, rowName2, colName2)
				weight2 := m[rowIdx][colIdx]
				if weight1 < weight2 {
					bitree.BT.CreateRightNode(mx, models.LbtfRoot, rowName1, colName1, true)
					bitree.BT.CreateLastNode(mx, models.LbtfRoot, rowName2, colName2)
				} else {
					bitree.BT.CreateRightNode(mx, models.LbtfRoot, rowName2, colName2, true)
					bitree.BT.CreateLastNode(mx, models.LbtfRoot, rowName1, colName1)
				}
				break
			}
		}
	}
}
