package main

import (
	"fmt"
	"salesmanstask/003/internal/bitree"
	"salesmanstask/003/internal/iteration"
	"salesmanstask/003/internal/methods"
	"salesmanstask/003/internal/models"
	"time"
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
	mtr, models.LbtfRoot = methods.MatrixConversion(matrixOriginal)
	if Debug {
		methods.PrintMatrix(mtr)
		fmt.Printf("H_root = %d\n", models.LbtfRoot)
		fmt.Println("приведенная матрица    ^^^")
		fmt.Println("____________________________________________________________________________")
	}

	/* создаем корневой узел дерева с параметрами:
	Q           критерий кратчайшего пути
	State       мапа с узлами дерева и копиями матриц отложенных узлов
	Count       счетчик узлов дерева
	Result      структура с результатами одной итерации (Маршрут и отложенные узлы с весам, приведенные матрицы узлов)
	CurrentNode текущий узел дерева
	RootNode    корневой узел дерева */
	bitree.BT = bitree.NewBiTree(matrixOriginal, models.LbtfRoot)

	// начинаем итерации создания ветвей:
	for {
		if models.Debug {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("###############################  NEW BRANCH #############################")
		}
		results, ok := iteration.Iteration(mtr, bitree.BT.RootNode)
		// break
		if ok {
			weight := 0
			mtr, bitree.BT.CurrentNode, weight, results = findInBack(results)

			if mtr == nil {
				fmt.Printf("!!! matrix is NIL !!!\n")
				break
			}
			if bitree.BT.RootNode == nil {
				fmt.Printf("!!! new root Node is NIL !!!\n")
				break
			}
			models.LbtfRoot = weight
		} else {
			fmt.Printf("NOT OK !!!\n")
			break
		}
		if models.Debug {
			fmt.Println("___________________________________________________")
			fmt.Printf("\nQ from object: %d\n", bitree.BT.Q)
			fmt.Printf("\nData from result:\n")
			fmt.Printf("\nTour:\n")
			for _, v := range results.Tour {
				fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
			}
			fmt.Printf("\nBack:\n")
			for _, v := range results.Back {
				fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
			}
		}
		// if true{
		// 	bitree.BT.Q = models.LbtfRoot
		// }
		if models.Debug {
			fmt.Println("###############################  stop branch #############################")
		}
		bitree.PrintTree(bitree.BT.RootNode)
	}

	fmt.Println("___________________________________________________")
	bitree.PrintTree(bitree.BT.RootNode)
	fmt.Printf("\nTour from state:\n")
	for _, v := range bitree.BT.Result.Tour {
		fmt.Printf("W:%d, (%d,%d)\n", v.W, v.Out, v.In)
	}

	printAllNodes()

}

func printAllNodes() {
	for _, v := range bitree.BT.Result.Back {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		methods.PrintMatrix(v.Mxs)
	}
}

func findInBack(res bitree.Results) ([][]int, *bitree.TreeNode, int, bitree.Results) {
	for i := 1; i < len(res.Back); i++ {
		if bitree.BT.Q > res.Back[i].W {
			fmt.Printf("Найдено в отложенных:  W:%d, %s(%d,%d), id: %d\n", res.Back[i].W, res.Back[i].Sign, res.Back[i].Out, res.Back[i].In, res.Back[i].ID)
			//models.LbtfRoot = res.Back[i].W
			matrix := res.Back[i].Mxs
			node := res.Back[i].Node
			w := res.Back[i].W

			res.Back[i] = res.Back[len(res.Back)-1]
			res.Back = res.Back[:len(res.Back)-1]

			return matrix, node, w, res
		}
	}
	return nil, nil, 0, bitree.Results{}
}
