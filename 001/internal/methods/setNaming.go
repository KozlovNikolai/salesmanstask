package methods

import "salesmanstask/001/internal/models"

func SetNaming(mx [][]int) *models.NamesOfIndexes {
	lenRows := len(mx)
	lenCols := len(mx[0])
	names := models.NamesOfIndexes{
		NamesOfRows: make([]int, lenRows),
		NamesOfCols: make([]int, lenCols),
	}
	for i := range names.NamesOfRows {
		names.NamesOfRows[i] = i + 1
	}
	for j := range names.NamesOfCols {
		names.NamesOfCols[j] = j + 1
	}
	return &names
}
