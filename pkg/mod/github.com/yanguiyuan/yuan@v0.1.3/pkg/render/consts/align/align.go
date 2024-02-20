package align

type Align int

const (
	Top Align = 1 << iota
	Bottom
	Left
	Right
	Center
)
