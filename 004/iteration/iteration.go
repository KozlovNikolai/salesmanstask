package iteration

import (
	"fmt"
	"log"
	"math"
	"salesmanstask/004/bitree"
	"salesmanstask/004/methods"
	"salesmanstask/004/models"
	"salesmanstask/data"
)

func IterationBranch() []bitree.Node {
	var toursArray []bitree.Node
	prevFoundWeight := math.MaxInt
	weight := 0
	// начинаем итерации создания ветвей:
	for {
		// начинаем итерации создания узлов:
		ok := IterationNode(models.MxRoot, bitree.BT.RootNode)
		if ok {
			fmt.Printf("Current Weight: %d\n", bitree.BT.CurWeight)
			fmt.Printf("Previous found Weight: %d\n", prevFoundWeight)
			if bitree.BT.CurWeight < prevFoundWeight {
				prevFoundWeight = bitree.BT.CurWeight
				toursArray = toursArray[:0]
				toursArray = append(toursArray, bitree.BT.Result.Tour...)
			}
			var row, col int
			// ищем в отложенных узлах узел с минимальным весом
			bitree.BT.CurrentNode, weight, row, col = findInBack()
			fmt.Printf("findBack - current Node: %v, weight: %d, row: %d, col: %d\n", bitree.BT.CurrentNode, weight, row, col)
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
			models.LowWeightLimit = weight
			models.MxRoot[row][col] = data.Inf
		} else {
			fmt.Printf("NOT OK !!!\n")
			break
		}
	}
	return toursArray
}

func IterationNode(matrix [][]int, node *bitree.TreeNode) bool {
	// создаем узлы ветви:
	for {
		mx := Step(matrix)
		if bitree.BT.CurWeight < bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W {
			fmt.Printf("\nBreak, вес лучшего маршрута:%d - меньше веса создаваемого\n маршрута: %d, дальше идти нет смысла.\n", bitree.BT.CurWeight, bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W)
			return false
		}
		if len(mx) == 2 {
			fmt.Printf("\nBreak, размер матрицы достиг: [%dx%d]\n", len(mx), len(mx[0]))
			EndingBranch(mx)
			// сохраняем найденный лучший вес и выходим
			bitree.BT.CurWeight = models.LowWeightLimit
			return true
		}
		matrix = bitree.CloneMx(mx)
	}
}

func Step(mc [][]int) [][]int {
	// ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextNode := methods.FindCellWithMaxMin(mc)

	// удаляем найденную ячейку с ее строкой и столбцом:
	reductionMatrix := methods.RemoveCellFromMatrixByIndex(mc, nextNode.RowName, nextNode.ColName)

	// помечаем ячейки для предотвращения внутренних циклов
	markInfinityCells(reductionMatrix, nextNode.ColName, nextNode.RowName)

	// получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	conversionMatrix, currentLowWeightLimit := methods.MatrixConversion(reductionMatrix)

	// определяем два новых узла и выбираем из них следующий
	var setCurrentRightNode bool
	if models.LowWeightLimit+nextNode.MaxSum >= models.LowWeightLimit+currentLowWeightLimit {
		setCurrentRightNode = true
	}
	bitree.BT.CreateLeftNode(models.LowWeightLimit+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bitree.BT.CreateRightNode(models.LowWeightLimit+currentLowWeightLimit, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		models.LowWeightLimit = models.LowWeightLimit + currentLowWeightLimit
	} else {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Left
		models.LowWeightLimit = models.LowWeightLimit + nextNode.MaxSum
	}
	return conversionMatrix
}

func markInfinityCells(mx3 [][]int, colName, rowName int) {
	rowInfIdx, colInfIdx, ok := methods.IdxByName(mx3, colName, rowName)
	if ok {
		mx3[rowInfIdx][colInfIdx] = data.Inf
	} else {
		log.Println("Второй: не могу получить индексы из имени !!!")
	}
}

func EndingBranch(mx [][]int) {
	rowIdx, colIdx, ok := methods.IdxByName(models.MxRoot, mx[1][0], mx[0][1])
	if !ok {
		log.Println("Ending branch: не могу получить индексы из имени !!!")
	}
	bitree.BT.CreateRightNode(models.LowWeightLimit, rowIdx, colIdx, true)
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

	if bitree.BT.CurWeight > bitree.BT.Result.Back[n].W {
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
