package main

import (
	"fmt"
	"salesmanstask/004/bitree"
	"salesmanstask/004/iteration"
	"salesmanstask/004/methods"
	"salesmanstask/004/models"

	"salesmanstask/data"
	"time"
)

var Debug = false

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
	matrixNamed := methods.SetNaming(mx)

	if Debug {
		methods.PrintMatrix(matrixNamed)
		fmt.Println("исходная матрица    ^^^")
		fmt.Println("____________________________________________________________________________")
	}

	models.MxRoot, models.LowWeightLimit = methods.MatrixConversion(matrixNamed)
	if Debug {
		// methods.PrintMatrix(mtr)
		methods.PrintMatrix(models.MxRoot)
		fmt.Printf("H_root = %d\n", models.LowWeightLimit)
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
	bitree.BT = bitree.NewBiTree(matrixNamed, models.LowWeightLimit)

	toursArray := iteration.IterationBranch()

	rt := make(map[int]int)
	fmt.Printf("\nResult tour with Q: %d\n", bitree.BT.CurWeight)
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
		fmt.Printf("(%d,%d),Cost:%d\n", temp, rt[temp], matrixNamed[temp][rt[temp]])
		sum += matrixNamed[temp][rt[temp]]
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
