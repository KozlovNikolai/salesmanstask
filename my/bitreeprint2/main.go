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

// Вывод дерева на печать
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	// Максимальная глубина дерева
	const maxDepth = 7
	width := (1 << maxDepth) - 1 // Максимальная ширина строки
	output := make([]string, maxDepth*2)

	// Вспомогательная функция для построения дерева
	var buildTree func(node *TreeNode, depth, left, right int)
	buildTree = func(node *TreeNode, depth, left, right int) {
		if node == nil || depth >= len(output) {
			return
		}

		// Расчет позиции значения в строке
		mid := (left + right) / 2
		val := fmt.Sprintf("%02d", node.val)

		// Если строка не создана, создаем её с достаточной шириной
		if len(output[depth]) == 0 {
			output[depth] = strings.Repeat(" ", width)
		}

		// Заполняем текущий уровень значением, проверяем границы
		line := []rune(output[depth])
		copy(line[mid:], []rune(val))
		output[depth] = string(line)

		// Рекурсивно обрабатываем левое и правое поддеревья
		buildTree(node.left, depth+2, left, mid-1)
		buildTree(node.right, depth+2, mid+1, right)
	}

	// Начинаем построение дерева
	buildTree(root, 0, 0, width-1)

	// Печатаем дерево
	for _, line := range output {
		if strings.TrimSpace(line) != "" {
			fmt.Println(line)
		}
	}
}

// Пример
func main() {
	// Построение тестового дерева
	value := 10
	root := &TreeNode{
		val: value + 1,
		left: &TreeNode{
			val: value + 2,
			left: &TreeNode{
				val:  value + 3,
				left: &TreeNode{val: value + 4},
				right: &TreeNode{
					val:  value + 5,
					left: &TreeNode{val: value + 6},
				},
			},
			right: &TreeNode{val: value + 7},
		},
		right: &TreeNode{
			val:  value + 8,
			left: &TreeNode{val: value + 9},
			right: &TreeNode{
				val: value + 10,
				right: &TreeNode{
					val:   value + 11,
					right: &TreeNode{val: value + 12},
				},
			},
		},
	}

	// Вывод дерева на печать
	printTree(root)
}
