package main

import (
	"fmt"
	"salesmanstask/002/internal/bitree"
	"salesmanstask/002/internal/methods"
	"salesmanstask/002/internal/models"
)

// var matrix = [][]int{
// 	{0, 0, -1},
// 	{-1, 3, 5},
// 	{2, -1, 0},
// }

// var matrix = [][]int{
// 	{-1, 0, 0, 2},
// 	{0, -1, 3, 5},
// 	{2, 12, 10, -1},
// 	{0, 2, 0, 0},
// }

// var matrix = [][]int{
// 	{-1, 5, 16, 14},
// 	{0, -1, 6, 9},
// 	{0, 12, -1, 11},
// 	{0, 15, 7, -1},
// }

// var matrix = [][]int{
// 	{-1, 5, 16, 14},
// 	{13, -1, 6, 9},
// 	{10, 12, -1, 11},
// 	{8, 15, 7, -1},
// }
// var matrix = [][]int{
// 	{-1, 0, 0, 2, 3},
// 	{6, -1, 9, 11, -1},
// 	{4, 14, -1, 11, 0},
// 	{2, 12, 10, -1, 0},
// 	{0, 2, 0, 0, -1},
// }

var matrix = [][]int{
	{-1, 1, 2, 3, 4},
	{14, -1, 15, 16, 5},
	{13, 20, -1, 17, 6},
	{12, 19, 18, -1, 7},
	{11, 10, 9, 8, -1},
}
var lbtfRoot int

var Debug = true

func main() {
	models.Debug = Debug
	// именуем столбцы и строки
	matrixOriginal := methods.SetNaming(matrix)

	if Debug {
		methods.PrintMatrix(matrixOriginal)
		fmt.Println("исходная матрица    ^^^")
		fmt.Println("____________________________________________________________________________")
	}

	var mtr [][]int
	mtr, lbtfRoot = methods.MatrixConversion(matrixOriginal)

	if Debug {
		methods.PrintMatrix(mtr)
		fmt.Printf("H_root = %d\n", lbtfRoot)
		fmt.Println("приведенная матрица    ^^^")
		fmt.Println("____________________________________________________________________________")

	}
	// for {
	// for bt.Q > lbtfRoot {
	for {
		mx := Step(mtr)
		if bt.Q < bt.Result.Tour[len(bt.Result.Tour)-1].W {
			fmt.Printf("Break, Q: %d, Tour: %d\n", bt.Q, bt.Result.Tour[len(bt.Result.Tour)-1].W)
			break
		}
		if len(mx) == 3 {
			fmt.Printf("Break, len(mx): %d\n", len(mx))
			EndingBranch(mx, matrixOriginal)
			break
		}
		mtr = mx
	}
	bt.Q = lbtfRoot
	fmt.Printf("Q: %d\n", bt.Q)
	// }
	// minimum := math.MaxInt
	// idx := 0
	// for i, v := range bt.Result.Back {
	// 	if v.W < minimum {
	// 		minimum = v.W
	// 		idx = i
	// 	}
	// }
	// if minimum < bt.Q {
	// 	bt.RootNode = bt.Result.Back[idx].Node
	// 	mtr, _ = methods.MatrixConversion(matrixOriginal)
	// }
	// }

	bitree.PrintTree(bt.RootNode)
	fmt.Printf("\nTour:\n")
	for _, v := range bt.Result.Tour {
		fmt.Printf("W:%d, (%d,%d)\n", v.W, v.Out, v.In)
	}
}

func EndingBranch(mx [][]int, m [][]int) {
	// fmt.Println("_________________________________")
	//bitree.PrintTree(bt.RootNode)
	// fmt.Println("_________________________________")
	if Debug {
		fmt.Printf("mx:\n%+v\n", mx)
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
				// fmt.Printf("weight-1: %d\n", weight1)
				// fmt.Printf("weight-2: %d\n", weight2)
				if weight1 < weight2 {
					bt.CreateRightNode(lbtfRoot, rowName1, colName2, true)
					bt.CreateLastNode(lbtfRoot, rowName2, colName2)
				} else {
					bt.CreateRightNode(lbtfRoot, rowName2, colName2, true)
					bt.CreateLastNode(lbtfRoot, rowName1, colName1)
				}
				break
			}
			if Debug {
				fmt.Println("пропускаем:")
			}

		}
	}
}

var bt *bitree.BiTree

func Step(mc [][]int) [][]int {
	if Debug {
		fmt.Println("---------START----------------")
	}

	// получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// "lower bound of the target function" => lbtfRoot:
	// mc, lbtfRoot := methods.MatrixConversion(matrix, names)

	// fmt.Printf("H source: %d\n\n", lbtfRoot)
	if bt == nil {
		// инициализируем дерево
		bt = bitree.NewBiTree(lbtfRoot)
	}

	// // ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextNode := methods.FindCellWithMaxMin(mc)
	// fmt.Printf("Next leaf: %+v\n\n", nextLeaf)

	// удаляем найденную ячейку с ее строкой и столбцом:
	rowIdx, colIdx := idxByName(mc, nextNode.RowName, nextNode.ColName)
	if Debug {
		fmt.Printf("RowName: %d, rowIdx: %d\n", nextNode.RowName, rowIdx)
		fmt.Printf("ColName: %d, colIdx: %d\n", nextNode.ColName, colIdx)
	}

	mx3 := methods.RemoveCellFromMatrixByIndex(mc, rowIdx, colIdx)

	if Debug {
		methods.PrintMatrix(mx3)
		fmt.Println("      удаление строки и столбца     ^^^")
		fmt.Println("_________________________________________________________________")

	}

	// // получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// // "lower bound of the target function" => lbtf:
	mx4, lbtfNode := methods.MatrixConversion(mx3)
	// fmt.Printf("H current: %d\n\n", lbtfNode)

	if Debug {
		methods.PrintMatrix(mx4)
		fmt.Printf("H_node = %d\n", lbtfNode)
		fmt.Println("      приведение матрицы     ^^^")
		fmt.Println("_________________________________________________________________")

	}

	var setCurrentRightNode bool
	if lbtfRoot+nextNode.MaxSum >= lbtfRoot+lbtfNode {
		setCurrentRightNode = true

	}
	// fmt.Printf("lbtfRoot: %d\n", lbtfRoot)
	// fmt.Printf("lbtfNode: %d\n", lbtfNode)
	// fmt.Printf("MaxSum: %d\n", nextLeaf.MaxSum)
	// fmt.Printf("Negativ: %d, Positive: %d, setCurrRight: %v\n", lbtfRoot+nextLeaf.MaxSum, lbtfRoot+lbtfNode, setCurrentRightNode)
	bt.CreateLeftNode(lbtfRoot+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bt.CreateRightNode(lbtfRoot+lbtfNode, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		lbtfRoot = lbtfRoot + lbtfNode
	}
	if Debug {
		fmt.Println("---------STOP----------------")
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
