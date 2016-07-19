package session

type Grid struct {
	Bars			[]int16            `json:"bars"`
}

func NewGrid() *Grid {
	g := new(Grid)

	return g
}
