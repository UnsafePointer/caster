package world

const (
	Width  = 640
	Heigth = 480
)

type Screen struct {
	Data [Width][Heigth][3]byte
}
