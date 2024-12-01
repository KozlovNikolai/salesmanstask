package main

import (
	"fmt"
	"salesmanstask/003/internal/bitree"
	"salesmanstask/003/internal/iteration"
	"salesmanstask/003/internal/methods"
	"salesmanstask/003/internal/models"
	"salesmanstask/data"
)

var Debug = false

func main() {
	for i := range data.Matrixes {
		fmt.Printf("\n#########################\n#\tMatrix: %d\t#\n#########################\n", i)
		Calculate(data.Matrixes[i])
	}
}

func Calculate(mx [][]int) {
	models.Debug = Debug
	// именуем столбцы и строки
	matrixOriginal := methods.SetNaming(mx)
	// methods.PrintMatrix(matrixOriginal)
	// fmt.Println("исходная матрица    ^^^")
	// fmt.Println("____________________________________________________________________________")

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
			//time.Sleep(1000 * time.Millisecond)
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
			fmt.Println("###############################  stop branch #############################")
			bitree.PrintTree(bitree.BT.RootNode)
		}

	}

	fmt.Println("___________________________________________________")
	bitree.PrintTree(bitree.BT.RootNode)
	fmt.Printf("\nTour from state:\n")
	for _, v := range bitree.BT.Result.Tour {
		fmt.Printf("ID:%d, W:%d, (%d,%d)\n", v.ID, v.W, v.Out, v.In)
	}
	if models.Debug {
		printAllNodes()
	}
}

func printAllNodes() {
	fmt.Println("Все узлы Маршрута:")
	for _, v := range bitree.BT.Result.Tour {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		methods.PrintMatrix(v.Mxs)
	}
	fmt.Println("Все отложенные узлы:")
	for _, v := range bitree.BT.Result.Back {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		methods.PrintMatrix(v.Mxs)
	}
}

func findInBack(res bitree.Results) ([][]int, *bitree.TreeNode, int, bitree.Results) {
	for i := 1; i < len(res.Back); i++ {
		if bitree.BT.Q > res.Back[i].W {
			fmt.Printf("Найдено в отложенных:  W:%d, %s(%d,%d), id: %d\n",
				res.Back[i].W,
				res.Back[i].Sign,
				res.Back[i].Out,
				res.Back[i].In,
				res.Back[i].ID)

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
