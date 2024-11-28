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

	fmt.Println("===dump1===")
	dump1(tree)
	fmt.Println("===dump2===")
	dump2(tree, "", "", "")

}
