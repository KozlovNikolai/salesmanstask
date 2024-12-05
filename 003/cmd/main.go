package main

import (
	"fmt"
	"math"
	"salesmanstask/003/internal/bitree"
	"salesmanstask/003/internal/iteration"
	"salesmanstask/003/internal/methods"
	"salesmanstask/003/internal/models"
	"salesmanstask/data"
	"time"
)

var Debug = true

func main() {
	for i := range data.Matrixes {
		fmt.Printf("\n#########################\n#\tMatrix: %d\t#\n#########################\n", i)
		t := time.Now()
		out := 0
		Calculate(data.Matrixes[i], out)

		ts := time.Since(t)
		fmt.Printf("Time latency: %v\n", ts)
	}
}

func Calculate(mx [][]int, out int) {
	models.Debug = Debug
	// устанавливаем город отправления
	if out != 0 {
		for i := range mx {
			if mx[i][out-1] != data.Inf {
				mx[i][out-1] = 0
			}

		}
	}
	// именуем столбцы и строки
	matrixOriginal := methods.SetNaming(mx)

	if Debug {
		methods.PrintMatrix(matrixOriginal)
		fmt.Println("исходная матрица    ^^^")
		fmt.Println("____________________________________________________________________________")
	}

	models.MxRoot, models.LbtfRoot = methods.MatrixConversion(matrixOriginal)
	if Debug {
		// methods.PrintMatrix(mtr)
		methods.PrintMatrix(models.MxRoot)
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
	var toursArray []bitree.Node
	prevQ := math.MaxInt
	weight := 0
	// начинаем итерации создания ветвей:
	for {
		if models.Debug {
			fmt.Println("###############################  NEW BRANCH #############################")
		}

		// ok := iteration.Iteration(mtr, bitree.BT.RootNode)
		ok := iteration.Iteration(models.MxRoot, bitree.BT.RootNode)
		if ok {
			// weight := 0
			fmt.Printf("Current Q: %d\n", bitree.BT.Q)
			// fmt.Printf("Previous Q: %d\n", bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W)
			fmt.Printf("Previous Q: %d\n", prevQ)
			if bitree.BT.Q < prevQ {
				prevQ = bitree.BT.Q
				toursArray = toursArray[:0]
				toursArray = append(toursArray, bitree.BT.Result.Tour...)

			}
			var row, col int
			bitree.BT.CurrentNode, weight, row, col = findInBack()
			fmt.Printf("findBack - current Node: %v, weight: %d, row: %d, col: %d\n", bitree.BT.CurrentNode, weight, row, col)
			// if mtr == nil {
			if bitree.BT.CurrentNode == nil {
				fmt.Printf("!!! current node is NIL !!!\n")
				break
			}
			if weight == 0 {
				fmt.Printf("!!! weight is Null !!!\n")
				break
			}
			if row == 0 || col == 0 {
				fmt.Printf("!!! Row or Col is Null !!!\n")
				break
			}
			models.LbtfRoot = weight

			models.MxRoot[row][col] = data.Inf
		} else {
			fmt.Printf("NOT OK !!!\n")
			break
		}
		if models.Debug {
			fmt.Println("___________________________________________________")
			fmt.Printf("\nQ from object: %d\n", bitree.BT.Q)
			fmt.Printf("\nData from result:\n")
			fmt.Printf("\nTour:\n")
			for _, v := range bitree.BT.Result.Tour {
				fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
			}
			fmt.Printf("\nBack:\n")
			for _, v := range bitree.BT.Result.Back {
				fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
			}
			fmt.Println("##########################################################################")
			bitree.PrintTree(bitree.BT.RootNode)
			fmt.Println("###############################  stop branch #############################")
			fmt.Printf("\ntour with Q: %d\n", bitree.BT.Q)
			for _, v := range toursArray {
				fmt.Printf("ID:%d, W:%d, (%d,%d)\n", v.ID, v.W, v.Out, v.In)
			}
		}

	}
	rt := make(map[int]int)
	fmt.Printf("\nResult tour with Q: %d\n", bitree.BT.Q)
	for _, v := range toursArray {
		fmt.Printf("ID:%d, W:%d, %s(%d,%d)\n", v.ID, v.W, v.Sign, v.Out, v.In)
		rt[v.Out] = v.In
	}
	temp := 1
	if out != 0 {
		temp = out
	}
	fmt.Printf("\nГород отправления: %d\n", temp)
	sum := 0
	for i := 0; i < len(rt); i++ {
		fmt.Printf("(%d,%d),Cost:%d\n", temp, rt[temp], matrixOriginal[temp][rt[temp]])
		sum += matrixOriginal[temp][rt[temp]]
		temp = rt[temp]
	}
	fmt.Printf("Sum: %d\n", sum)

	if models.Debug {
		fmt.Println("___________________________________________________")
		bitree.PrintTree(bitree.BT.RootNode)
		listTour := make([]int, 0)
		fmt.Printf("Length of Tour array: %d\n", len(bitree.BT.Result.Tour))
		fmt.Printf("\nTour from state:\n")
		for _, v := range bitree.BT.Result.Tour {
			fmt.Printf("ID:%d, W:%d, (%d,%d)\n", v.ID, v.W, v.Out, v.In)
			listTour = append(listTour, v.Out)
		}

		fmt.Printf("\nSort Tour from state:\n")
		for _, v := range listTour {
			fmt.Printf("%d,", v)
		}
	}
	if models.Debug {
		//	printAllNodes()
	}
}

func printAllNodes() {
	fmt.Println("Все узлы Маршрута:")
	for _, v := range bitree.BT.Result.Tour {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		// methods.PrintMatrix(v.Mxs)
		methods.PrintMatrix(models.MxRoot)
	}
	fmt.Println("Все отложенные узлы:")
	for _, v := range bitree.BT.Result.Back {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		// methods.PrintMatrix(v.Mxs)
		methods.PrintMatrix(models.MxRoot)
	}
}

func findInBack() (*bitree.TreeNode, int, int, int) {
	fmt.Printf("Поиск в отложенных узлах: %d штук\n", len(bitree.BT.Result.Back))
	minWeight := math.MaxInt
	var n int
	for i := 1; i < len(bitree.BT.Result.Back); i++ {

		if bitree.BT.Result.Back[i].W < minWeight {
			minWeight = bitree.BT.Result.Back[i].W
			n = i
		}
	}

	if bitree.BT.Q > bitree.BT.Result.Back[n].W {
		fmt.Printf("Найдено в отложенных:  W:%d, %s(%d,%d), id: %d\n",
			bitree.BT.Result.Back[n].W,
			bitree.BT.Result.Back[n].Sign,
			bitree.BT.Result.Back[n].Out,
			bitree.BT.Result.Back[n].In,
			bitree.BT.Result.Back[n].ID)

		node := bitree.BT.Result.Back[n].Node
		w := bitree.BT.Result.Back[n].W
		row := bitree.BT.Result.Back[n].Out
		col := bitree.BT.Result.Back[n].In
		bitree.BT.Result.Back[n] = bitree.BT.Result.Back[len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Back = bitree.BT.Result.Back[:len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Tour = bitree.BT.Result.Tour[:0]
		return node, w, row, col
	}

	return nil, 0, 0, 0
}
