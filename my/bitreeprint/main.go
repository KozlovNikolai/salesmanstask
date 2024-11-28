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
	const maxDepth = 10
	output := make([]string, maxDepth*2) // Массив строк для вывода
	count := 0
	// Вспомогательная функция для построения дерева
	var buildTree func(node *TreeNode, depth, pos, width int)
	buildTree = func(node *TreeNode, depth, pos, width int) {
		count++
		// fmt.Printf("%d, ", count)
		fmt.Printf("i: %d, depth: %d, pos: %d, width: %d\n", count, depth, pos, width)
		if node == nil || depth >= len(output) {
			return
		}

		// Заполняем текущий уровень значением
		line := strings.Repeat(" ", width*2)
		val := fmt.Sprintf("%03d", node.val)
		line = line[:pos] + val + line[pos+len(val):]
		output[depth] = output[depth] + line

		// Рекурсивно обрабатываем левое и правое поддеревья
		nextWidth := width / 2
		buildTree(node.left, depth+2, pos-nextWidth, nextWidth)
		buildTree(node.right, depth+2, pos+nextWidth, nextWidth)
	}

	// Начинаем построение дерева
	fmt.Printf("pos: %d\n", 1<<(maxDepth-1)-1)

	buildTree(root, 0, (1<<(maxDepth-1))-1, (1<<(maxDepth-1))-1)

	// Печатаем дерево
	for _, line := range output {
		if line != "" {
			fmt.Println(line)
		}
	}
}

// Пример
func main() {
	// Построение тестового дерева
	root := &TreeNode{
		val: 999,
		left: &TreeNode{
			val: 999,
			left: &TreeNode{
				val:  999,
				left: &TreeNode{val: 999},
				right: &TreeNode{
					val:  999,
					left: &TreeNode{val: 999},
				},
			},
			right: &TreeNode{val: 999},
		},
		right: &TreeNode{
			val:  999,
			left: &TreeNode{val: 999},
			right: &TreeNode{
				val: 999,
				right: &TreeNode{
					val:   999,
					right: &TreeNode{val: 999},
				},
			},
		},
	}

	// Вывод дерева на печать
	printTree(root)
}
