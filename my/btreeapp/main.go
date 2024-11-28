package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {

	t := &TreeNode{val: 18111111113}

	t.Insert(11111111113)
	t.Insert(12111111113)
	t.Insert(13111111113)
	t.Insert(14111111113)
	t.Insert(15111111113)
	t.Insert(16111111113)
	t.Insert(17111111113)
	t.Insert(19111111113)
	t.Insert(20111111113)
	t.Insert(21111111113)
	t.Insert(22111111113)
	t.Insert(23111111113)

	dump1(t)

	mx := &TreeNode{val: 180111111113}

	mx.Insert(130111111113)
	mx.Insert(120111111113)
	mx.Insert(110111111113)
	mx.Insert(140111111113)
	mx.Insert(150111111113)
	mx.Insert(160111111113)
	mx.Insert(170111111113)
	mx.Insert(190111111113)
	mx.Insert(200111111113)
	mx.Insert(210111111113)
	mx.Insert(220111111113)
	mx.Insert(230111111113)

	dump1(mx)

}

// PrintInorder prints the elements in order
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Printf("%d,", t.val)
	t.right.PrintInorder()
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.val == value {
		return errors.New("this node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

// Find finds the treenode for the given node val
func (t *TreeNode) Find(value int) (TreeNode, bool) {
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
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(value int) *TreeNode {
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
func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

// FindMin finds the min element in the given BST
func (t *TreeNode) FindMin() int {
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

// Функция dump1 - реализация визуализации дерева
func dump1(node *TreeNode) {
	// Вспомогательная рекурсивная функция
	var Rec func(*TreeNode, bool) ([]string, int, int)
	Rec = func(node *TreeNode, left bool) ([]string, int, int) {
		if node == nil {
			return []string{}, 0, 0
		}
		sval := strconv.Itoa(node.val)
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
