package app

var Debug bool

var rootMatrix [][]int

func SetRootMatrix(mx [][]int) {
	rootMatrix = CloneMx(mx)
}

func GetRootMatrix() [][]int {
	return CloneMx(rootMatrix)
}

type Node struct {
	ID       int
	Name     string
	Out      int
	In       int
	ParentID int
	MX       [][]int
	W        int
}

type Leaves struct {
	NodeIDs     map[int]struct{}
	MinWeightID int
	MinWeight   int
}

type Store struct {
	Tree          map[int]*Node
	Leaves        Leaves
	CurrentNodeID int
	NextID        int
	IsSolved      bool
	Start         int
}

func NewStore(mx [][]int, lb, start int) *Store {
	return &Store{
		Tree: map[int]*Node{0: {
			ID:       0,
			Name:     "Root",
			Out:      0,
			In:       0,
			ParentID: 0,
			MX:       mx,
			W:        lb,
		}},
		Leaves: Leaves{
			NodeIDs:     map[int]struct{}{0: {}},
			MinWeightID: 0,
			MinWeight:   lb,
		},
		CurrentNodeID: 0,
		NextID:        1,
		Start:         start,
	}
}

func (s *Store) InsertLeaf() {
	s.Leaves.NodeIDs[s.NextID-1] = struct{}{}
}
func (s *Store) RemoveLeaf() {
	delete(s.Leaves.NodeIDs, s.CurrentNodeID)
	s.FindMinWeightLeaf()
}

func (s *Store) AddParentsWeight(nodeID int) int {
	if nodeID == 0 {
		return 0
	}
	node := s.Tree[nodeID]
	return node.W + s.AddParentsWeight(node.ParentID)
}

func (s *Store) FindMinWeightLeaf() {
	minWeightNodeIdx := 0
	minWeight := Inf
	s.funcHelp(&minWeight, &minWeightNodeIdx)
	s.Leaves.MinWeightID = minWeightNodeIdx
	s.Leaves.MinWeight = minWeight
	s.CurrentNodeID = minWeightNodeIdx
}

func (s *Store) funcHelp(minWeight, minWeightNodeIdx *int) {
	for i := range s.Leaves.NodeIDs {
		w := s.Tree[i].W
		if w < *minWeight {
			*minWeight = w
			*minWeightNodeIdx = i
		}
	}
}

func (s *Store) AddNode(mx [][]int, out, in, w int) int {
	mxcopy := CloneMx(mx)
	s.Tree[s.NextID] = &Node{
		ID:       s.NextID,
		Name:     "",
		Out:      out,
		In:       in,
		ParentID: s.CurrentNodeID,
		MX:       mxcopy,
		W:        w,
	}
	node := s.Tree[s.NextID]
	s.NextID++
	return node.ID
}

func CloneMx(mx [][]int) [][]int {
	lenRows := len(mx)
	lenCols := len(mx[0])
	mxClone := make([][]int, lenRows)
	for i := range mxClone {
		mxClone[i] = make([]int, lenCols)
	}
	for i := 0; i < lenRows; i++ {
		for j := 0; j < lenCols; j++ {
			mxClone[i][j] = mx[i][j]
		}
	}
	return mxClone
}
