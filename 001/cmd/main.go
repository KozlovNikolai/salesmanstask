package main

import (
	"fmt"
	"salesmanstask/001/internal/bitree"
	"salesmanstask/001/internal/methods"
	"salesmanstask/001/internal/models"
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

func main() {
	// именуем столбцы и строки
	names := methods.SetNaming(matrix)

	// выводим исходную матрицу:
	// fmt.Println("Source matrix:")
	methods.PrintMatrix(matrix, names)
	fmt.Println("исходная матрица    ^^^")
	fmt.Println("____________________________________________________________________________")
	var mtr [][]int
	mtr, lbtfRoot = methods.MatrixConversion(matrix, names)

	for {
		mx, _ := Step(mtr, names)
		if bt.Q < bt.Result.Tour[len(bt.Result.Tour)-1].W {
			// fmt.Printf("Break, Q: %d, Tour: %d\n", bt.Q, bt.Result.Tour[len(bt.Result.Tour)-1].W)
			break
		}
		if len(mx) == 2 {
			// fmt.Printf("Break, len(mx): %d\n", len(mx))
			// EndingBranch(mx, names)
			break
		}
		// bitree.PrintTree(bt.RootNode)

		//	fmt.Scanln()
		mtr = mx
	}

	// methods.PrintMatrix(mx, names)
	bitree.PrintTree(bt.RootNode)

	for _, v := range bt.Result.Tour {
		fmt.Printf("W:%d, (%d,%d)\n", v.W, v.Out, v.In)
	}

}

func EndingBranch(mx [][]int, names *models.NamesOfIndexes) {
	// fmt.Println("_________________________________")
	bitree.PrintTree(bt.RootNode)
	// fmt.Println("_________________________________")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if mx[i][j] != -1 {
				rowName1 := names.GetRowName(i)
				colName1 := names.GetColName(j)
				weight1 := matrix[rowName1-1][colName1-1]

				rowName2 := names.GetRowName(i ^ 1)
				colName2 := names.GetColName(j ^ 1)
				weight2 := matrix[rowName2-1][colName2-1]
				// fmt.Printf("weight-1: %d\n", weight1)
				// fmt.Printf("weight-2: %d\n", weight2)
				if weight1 < weight2 {
					bt.CreateRightNode(lbtfRoot, rowName1, colName1, true)
					bt.CreateLastNode(lbtfRoot, rowName2, colName2)
				} else {
					bt.CreateRightNode(lbtfRoot, rowName2, colName2, true)
					bt.CreateLastNode(lbtfRoot, rowName1, colName1)
				}
				break
			}
		}
	}

}

var bt *bitree.BiTree

func Step(mc [][]int, names *models.NamesOfIndexes) ([][]int, *bitree.TreeNode) {
	fmt.Println("---------START----------------")
	// получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// "lower bound of the target function" => lbtfRoot:
	// mc, lbtfRoot := methods.MatrixConversion(matrix, names)

	// fmt.Printf("H source: %d\n\n", lbtfRoot)
	if bt == nil {
		// инициализируем дерево
		bt = bitree.NewBiTree(lbtfRoot)
	}

	// ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextLeaf := methods.FindCellWithMaxMin(mc, names)
	// fmt.Printf("Next leaf: %+v\n\n", nextLeaf)

	// удаляем найденную ячейку с ее строкой и столбцом:

	mx3, nm3 := methods.RemoveCellFromMatrixByIndex(mc, names.GetRowIdx(nextLeaf.RowName), names.GetColIdx(nextLeaf.ColName), names)
	methods.PrintMatrix(mx3, nm3)
	fmt.Println("      удаление строки и столбца     ^^^")
	fmt.Println("_________________________________________________________________")

	// получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	// "lower bound of the target function" => lbtf:
	mx4, lbtfNode := methods.MatrixConversion(mx3, nm3)
	// fmt.Printf("H current: %d\n\n", lbtfNode)
	methods.PrintMatrix(mx4, nm3)
	fmt.Println("      приведение матрицы     ^^^")
	fmt.Println("_________________________________________________________________")

	var setCurrentRightNode bool
	if lbtfRoot+nextLeaf.MaxSum >= lbtfRoot+lbtfNode {
		setCurrentRightNode = true

	}
	// fmt.Printf("lbtfRoot: %d\n", lbtfRoot)
	// fmt.Printf("lbtfNode: %d\n", lbtfNode)
	// fmt.Printf("MaxSum: %d\n", nextLeaf.MaxSum)
	// fmt.Printf("Negativ: %d, Positive: %d, setCurrRight: %v\n", lbtfRoot+nextLeaf.MaxSum, lbtfRoot+lbtfNode, setCurrentRightNode)
	bt.CreateLeftNode(lbtfRoot+nextLeaf.MaxSum, nextLeaf.RowName, nextLeaf.ColName, !setCurrentRightNode)
	bt.CreateRightNode(lbtfRoot+lbtfNode, nextLeaf.RowName, nextLeaf.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		lbtfRoot = lbtfRoot + lbtfNode
	}
	fmt.Println("---------STOP----------------")
	return mx4, bt.RootNode
}
