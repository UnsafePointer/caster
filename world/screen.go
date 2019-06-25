package world

const (
	Width  = 1280
	Heigth = 720
	Scale  = 1
)

type Screen struct {
	Data [Width][Heigth][3]byte
}
