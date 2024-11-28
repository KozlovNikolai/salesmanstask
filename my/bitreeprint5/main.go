package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BinTree struct {
	Value int
	Left  *BinTree
	Right *BinTree
}

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

// dump0 - простой вариант печати дерева
func dump0(node *BinTree, prefix string, root bool, last bool) {
	if root {
		fmt.Println(prefix + fmt.Sprintf("%d", node.Value))
	} else {
		if last {
			fmt.Println(prefix + chUDiaHor + fmt.Sprintf("%d", node.Value))
		} else {
			fmt.Println(prefix + chVerHor + fmt.Sprintf("%d", node.Value))
		}
	}

	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}

	children := []*BinTree{node.Left, node.Right}
	for i, child := range children {
		if child != nil {
			var tempvalue string
			if root {
				tempvalue = ""
			} else {
				if last {
					tempvalue = "  "
				} else {
					tempvalue = chVerSpa
				}
			}
			// root ? "" : (last ? "  " : chVerSpa)
			dump0(child, prefix+tempvalue, false, i == len(children)-1)
		}
	}
}

// Вспомогательная функция для повторения строки несколько раз
func RepStr(s string, cnt int) string {
	if cnt < 0 {
		panic(fmt.Sprintf("RepStr: Некорректное значение %d!", cnt))
	}
	return strings.Repeat(s, cnt)
}

// Функция dump1 - реализация визуализации дерева
func dump1(node *BinTree) {
	// Вспомогательная рекурсивная функция
	var Rec func(*BinTree, bool) ([]string, int, int)
	Rec = func(node *BinTree, left bool) ([]string, int, int) {
		if node == nil {
			return []string{}, 0, 0
		}
		sval := strconv.Itoa(node.Value)
		resl, cl, lss := Rec(node.Left, true)
		resr, cr, rss := Rec(node.Right, false)

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

// dump2 - печать с линиями в стиле ASCII
func dump2(node *BinTree, rpref, cpref, lpref string) {
	if node == nil {
		return
	}
	if node.Right != nil {
		dump2(node.Right, rpref+"  ", rpref+chDDiaHor, rpref+chVerSpa)
	}
	fmt.Println(cpref + strconv.Itoa(node.Value))
	if node.Left != nil {
		dump2(node.Left, lpref+chVerSpa, lpref+chUDiaHor, lpref+"  ")
	}
}

// dump3 - простой отступ для визуализации дерева
func dump3(node *BinTree, space int) {
	if node == nil {
		return
	}

	const count = 2
	space += count
	dump3(node.Right, space)

	for i := count; i < space; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Value)

	dump3(node.Left, space)
}

// dump4 - сложная визуализация с Unicode символами
func dump4(node *BinTree, high bool, lpref, cpref, rpref []string, root, left bool, lines *[][]string) {
	if node == nil {
		return
	}

	var VSCat = func(a, b []string) []string {
		return append(a, b...)
	}

	if root {
		*lines = [][]string{}
	}
	if node.Left != nil {
		dump4(node.Left, high, VSCat(lpref, []string{" ", " "}), VSCat(lpref, []string{chDDia, chVer}), VSCat(lpref, []string{chHor, " "}), false, true, lines)
	}

	sval := strconv.Itoa(node.Value)
	sm := 0
	if left || sval == "" {
		sm = len(sval) / 2
	} else {
		sm = (len(sval)+1)/2 - 1
	}

	for i := 0; i < len(sval); i++ {
		if i < sm {
			*lines = append(*lines, VSCat(lpref, []string{string(sval[i])}))
		} else if i == sm {
			*lines = append(*lines, VSCat(cpref, []string{string(sval[i])}))
		} else {
			*lines = append(*lines, VSCat(rpref, []string{string(sval[i])}))
		}
	}

	if node.Right != nil {
		dump4(node.Right, high, VSCat(rpref, []string{chHor, " "}), VSCat(rpref, []string{chRDia, chVer}), VSCat(rpref, []string{" ", " "}), false, false, lines)
	}

	if root {
		for _, line := range *lines {
			fmt.Println(strings.Join(line, ""))
		}
	}
}

func main() {
	// Пример построения бинарного дерева
	tree := &BinTree{
		Value: 10,
		Left: &BinTree{
			Value: 5,
			Left: &BinTree{
				Value: 1,
				Right: &BinTree{Value: 2},
			},
			Right: &BinTree{
				Value: 6,
				Right: &BinTree{
					Value: 8,
					Left:  &BinTree{Value: 7},
				},
			},
		},
		Right: &BinTree{
			Value: 19,
			Left:  &BinTree{Value: 17},
			Right: &BinTree{
				Value: 21,
				Left:  &BinTree{Value: 20},
				Right: &BinTree{Value: 250},
			},
		},
	}

	fmt.Println("===dump0===")
	dump0(tree, "", true, true)
	fmt.Println("===dump1===")
	dump1(tree)
	fmt.Println("===dump2===")
	dump2(tree, "", "", "")
	fmt.Println("===dump3===")
	dump3(tree, 0)
	fmt.Println("===dump4===")
	lines := [][]string{}
	dump4(tree, true, []string{}, []string{}, []string{}, true, true, &lines)
}
