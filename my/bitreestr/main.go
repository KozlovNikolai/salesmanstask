package main

import (
	"errors"
	"fmt"
	"strings"
)

// type BinTree struct {
// 	Value int
// 	Left  *BinTree
// 	Right *BinTree
// }

// Unicode символы для визуализации
var (
	chHor     = "─"
	chVer     = "│"
	chDDia    = "┌"
	chRDia    = "┐"
	chUDia    = "└"
	chVerHor  = "├─"
	chUDiaHor = "└─"
	chDDiaHor = "┌─"
	chVerSpa  = "│ "
)

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

type Results struct {
	Tour []Node
	Back []Node
}
type Node struct {
	id  int
	w   int
	in  int
	out int
}

func main() {
	mem := make(map[int]*TreeNode)
	count := 0
	res := Results{}
	var currentNode *TreeNode

	// создаем Root и сохраняем в мапе
	nd := Node{id: count, w: 31}
	t := &TreeNode{val: fmt.Sprintf("w%d:Root", nd.w)}
	mem[count] = t
	count++
	currentNode = t

	//#############################################################################################################################
	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 37, out: 2, in: 5}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 37, out: 2, in: 5}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++
	//###################################################################################################################

	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 47, out: 1, in: 2}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 37, out: 1, in: 2}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	//#############################################################################################################################
	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 45, out: 4, in: 1}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 44, out: 4, in: 1}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	//#############################################################################################################################

	// создаем один позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 44, out: 5, in: 3}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	//#############################################################################################################################

	// создаем последний узел, сохраняем в мапе и в туре
	nd = Node{id: count, w: 44, out: 3, in: 4}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	count++

	fmt.Println("Tour:")
	for _, v := range res.Tour {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("Back:")
	for _, v := range res.Back {
		fmt.Printf("%+v\n", v)
	}
	PrintTree(t)

	//#############################################################################################################################
	//#############################################################################################################################
	//#############################################################################################################################

	// текущее решение имеет вес - 44
	// в Back списке нашли узел номер: 2 с меньшим весом: 37

	// сохраняем первое найденное решение и чистим массив результатов
	res1 := res
	_ = res1
	res = Results{}

	// назначаем рутом новый узел
	currentNode, ok := mem[1]
	_ = ok
	// продолжаем вычисление и заполняем дерево дальше
	//###################################################################################################################

	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 41, out: 3, in: 5}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 39, out: 3, in: 5}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	//###################################################################################################################

	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 47, out: 4, in: 1}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 42, out: 4, in: 1}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++
	//###################################################################################################################

	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 46, out: 5, in: 4}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 42, out: 5, in: 4}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++
	//#############################################################################################################################

	// создаем один позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 42, out: 1, in: 2}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	//#############################################################################################################################

	// создаем последний узел, сохраняем в мапе и в туре
	nd = Node{id: count, w: 42, out: 2, in: 3}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	count++

	fmt.Println("Tour:")
	for _, v := range res.Tour {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("Back:")
	for _, v := range res.Back {
		fmt.Printf("%+v\n", v)
	}
	PrintTree(t)
	//#############################################################################################################################
	//#############################################################################################################################
	//#############################################################################################################################

	// текущее решение имеет вес - 42
	// в Back списке нашли узел номер: 10 с меньшим весом: 41

	// сохраняем первое найденное решение и чистим массив результатов
	res2 := res
	_ = res2
	res = Results{}

	// назначаем рутом новый узел
	currentNode, ok = mem[9]
	_ = ok
	// продолжаем вычисление и заполняем дерево дальше
	//###################################################################################################################
	// создаем негативный узел и сохраняем в мапе и в отложенном списке
	nd = Node{id: count, w: 48, out: 3, in: 1}
	currentNode.InsertLeft(fmt.Sprintf("w%d:-%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.left
	res.Back = append(res.Back, nd)
	count++

	// создаем позитивный узел, сохраняем в мапе и в туре, назначаем его текущим узлом
	nd = Node{id: count, w: 44, out: 3, in: 1}
	currentNode.InsertRight(fmt.Sprintf("w%d:%d.%d", nd.w, nd.out, nd.in))
	mem[count] = currentNode.right
	res.Tour = append(res.Tour, nd)
	currentNode = currentNode.right
	count++

	fmt.Println("Tour:")
	for _, v := range res.Tour {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("Back:")
	for _, v := range res.Back {
		fmt.Printf("%+v\n", v)
	}
	PrintTree(t)
	_ = currentNode
}

// PrintInorder prints the elements in order
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Printf("%s,", t.val)
	t.right.PrintInorder()
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) InsertLeft(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	// if t.val == value {
	// 	return errors.New("this node value already exists")
	// }

	if t.left == nil {
		t.left = &TreeNode{val: value}
		return nil
	}
	return fmt.Errorf("левый лист узла: \"%s\" занят", t.val)
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) InsertRight(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	// if t.val == value {
	// 	return errors.New("this node value already exists")
	// }

	if t.right == nil {
		t.right = &TreeNode{val: value}
		return nil
	}
	return fmt.Errorf("правый лист узла: \"%s\" занят", t.val)
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) Insert(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.val == value {
		return errors.New("this node value already exists")
	}

	if value < t.val {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if value > t.val {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

// Find finds the treenode for the given node val
func (t *TreeNode) Find(value string) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.val:
		return *t, true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

// Delete removes the Item with value from the tree
func (t *TreeNode) Delete(value string) {
	t.remove(value)
}

func (t *TreeNode) remove(value string) *TreeNode {
	if t == nil {
		return nil
	}
	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}
	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}
	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

// FindMax finds the max element in the given BST
func (t *TreeNode) FindMax() string {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

// FindMin finds the min element in the given BST
func (t *TreeNode) FindMin() string {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

// Вспомогательная функция для повторения строки несколько раз
func RepStr(s string, cnt int) string {
	if cnt < 0 {
		panic(fmt.Sprintf("RepStr: Некорректное значение %d!", cnt))
	}
	return strings.Repeat(s, cnt)
}

// Функция PrintTree - реализация визуализации дерева
func PrintTree(node *TreeNode) {
	// Вспомогательная рекурсивная функция
	var Rec func(*TreeNode, bool) ([]string, int, int)
	Rec = func(node *TreeNode, left bool) ([]string, int, int) {
		if node == nil {
			return []string{}, 0, 0
		}
		// sval := strconv.Itoa(node.val)
		sval := node.val
		resl, cl, lss := Rec(node.left, true)
		resr, cr, rss := Rec(node.right, false)

		vl := resl
		vr := resr

		lv := len(sval)
		ls := 0
		if len(vl) > 0 {
			ls = lss
		}
		rs := 0
		if len(vr) > 0 {
			rs = rss
		}

		lis := 0
		if ls == 0 {
			lis = lv / 2
		} else {
			lis = max(lv/2+1-(ls-cl), 0)
		}
		ris := 0
		if rs == 0 {
			ris = (lv + 1) / 2
		} else {
			ris = max((lv+1)/2-cr, 0)
			if lis == 0 {
				ris = max(ris, 1)
			}
		}

		dashls := 0
		if ls != 0 {
			dashls = ls - cl - 1 + lis - lv/2
		}
		dashrs := 0
		if rs != 0 {
			dashrs = cr + ris - (lv+1)/2
		}

		// Формируем первую строку
		line := ""
		if ls != 0 {
			line += RepStr(" ", cl) + chDDia + RepStr(chHor, dashls)
		}
		line += sval
		if rs != 0 {
			line += RepStr(chHor, dashrs) + chRDia + RepStr(" ", rs-cr-1)
		}

		lines := []string{line}

		// Формируем остальные строки
		for i := 0; i < max(len(vl), len(vr)); i++ {
			sl := RepStr(" ", ls)
			sr := RepStr(" ", rs)
			if i < len(vl) {
				sl = vl[i]
			}
			if i < len(vr) {
				sr = vr[i]
			}
			sl += RepStr(" ", lis)
			sr = RepStr(" ", ris) + sr
			lines = append(lines, sl+sr)
		}
		var adjustedWidth int
		if left || ls+lis == 0 || lv%2 == 1 {
			adjustedWidth = ls + lis
		} else {
			adjustedWidth = ls + lis - 1
		}

		return lines, adjustedWidth, ls + lis + ris + rs
		// return lines, (left || ls+lis == 0 || lv%2 == 1) ? ls+lis : ls+lis-1, ls+lis+ris+rs
	}

	// Запуск визуализации дерева
	lines, _, _ := Rec(node, true)
	for _, line := range lines {
		fmt.Println(line)
	}
}
