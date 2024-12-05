package app

import "log"

func SetNaming(mx [][]int) {
	lenRows := len(mx)
	lenCols := len(mx[0])

	mmx := make([][]int, lenRows+1)
	for i := range mmx {
		mmx[i] = make([]int, lenCols+1)
		// заполняем заголовки столбцов:
		if i == 0 {
			for j := range mmx[i] {
				mmx[i][j] = j
			}
		} else {
			mmx[i][0] = i
			for j := 1; j < len(mmx[i]); j++ {
				mmx[i][j] = mx[i-1][j-1]
			}
		}

	}

	SetRootMatrix(mmx)
}

func SetStart(mx [][]int, start int) {
	if start < 0 || start > len(mx) {
		log.Fatalf("start point: %d out of range: 0...%d", start, len(mx))
	}
	if start != 0 {
		for i := range mx {
			if mx[i][start-1] != Inf {
				mx[i][start-1] = 0
			}
		}
	}
}
