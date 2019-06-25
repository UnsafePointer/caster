package graphics

import (
	"fmt"
	"image"
	"os"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Texture struct {
	pixels [][]Pixel
	Width  int
	Height int
}

func (t *Texture) GetPixelAt(x int, y int) *Pixel {
	x = x % t.Width
	y = y % t.Height
	p := &t.pixels[x][y]
	return p
}

func NewTexture(filePath string) *Texture {
	f, err := os.Open("./texture.png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer f.Close()
	img, _, err := image.Decode(f)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	p, err := getPixels(width, height, img)

	t := &Texture{
		pixels: p,
		Width:  width,
		Height: height,
	}
	return t
}

func getPixels(height int, width int, img image.Image) ([][]Pixel, error) {
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}
