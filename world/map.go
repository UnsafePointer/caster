package world

type Map struct {
	Tiles [][]bool
}

func NewMap() *Map {
	m := &Map{
		Tiles: [][]bool{
			{true, true, true, true, true, true},
			{true, false, true, true, false, true},
			{true, false, false, false, false, true},
			{true, false, false, false, false, true},
			{true, false, true, true, false, true},
			{true, true, true, true, true, true},
		},
	}
	return m
}
