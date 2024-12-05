package app

import (
	"math"
)

func GetLB(mx [][]int) int {
	count := len(mx)
	sum := 0

	for i := 1; i < count; i++ {
		min := math.MaxInt
		for j := 1; j < count; j++ {
			if mx[i][j] < min {
				min = mx[i][j]
			}
		}
		sum += min
	}
	return sum
}

func GetH(mx [][]int, rowName, colName int, s *Store) int {
	rootMx := GetRootMatrix()

	count := len(mx)
	var sum int

	for i := 1; i < count; i++ {
		min := math.MaxInt
		for j := 1; j < count; j++ {
			if mx[i][j] < min {
				min = mx[i][j]
			}
		}
		sum += min
	}

	nodeID := s.CurrentNodeID
	for nodeID != 0 {
		sum += rootMx[s.Tree[nodeID].Out][s.Tree[nodeID].In]
		nodeID = s.Tree[nodeID].ParentID
	}
	if nodeID == 0 {
		sum += rootMx[rowName][colName]
	}

	return sum
}
