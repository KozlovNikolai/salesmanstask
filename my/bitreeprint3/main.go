package main

import (
	"fmt"
	"strings"
)

// Определение структуры узла дерева
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Получение высоты дерева
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.left)
	rightHeight := getHeight(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Прямое отображение дерева в матрицу
func printTreeMatrix(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	height := getHeight(root)  // Высота дерева
	width := (1 << height) - 1 // Ширина матрицы: 2^height - 1 (максимальная ширина)
	matrix := make([][]string, height)

	// Инициализация пустой матрицы
	for i := range matrix {
		matrix[i] = make([]string, width)
		for j := range matrix[i] {
			matrix[i][j] = " "
		}
	}

	// Вспомогательная функция для размещения узлов в матрице
	var fillMatrix func(node *TreeNode, row, left, right int)
	fillMatrix = func(node *TreeNode, row, left, right int) {
		if node == nil || row >= height {
			return
		}
		// Расчет позиции узла в строке
		mid := (left + right) / 2
		matrix[row][mid] = fmt.Sprintf("%03d", node.val) // Записываем значение с ведущими нулями

		// Рекурсивно размещаем левое и правое поддеревья
		fillMatrix(node.left, row+1, left, mid-1)
		fillMatrix(node.right, row+1, mid+1, right)
	}

	// Заполняем матрицу значениями из дерева
	fillMatrix(root, 0, 0, width-1)

	// Печатаем матрицу
	for _, line := range matrix {
		fmt.Println(strings.Join(line, ""))
	}
}

func main() {
	// Пример дерева
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 4,
			},
			right: &TreeNode{
				val: 5,
			},
		},
		right: &TreeNode{
			val: 3,
		},
	}

	// Вывод дерева в виде матрицы
	printTreeMatrix(root)
}
