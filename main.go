package main

import (
	"fmt"
	"image"
	"image/png"
	"math"

	"github.com/Ruenzuo/caster/geometry"
	"github.com/Ruenzuo/caster/graphics"
	"github.com/Ruenzuo/caster/world"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	renderer *world.Renderer
	screen   *world.Screen
	camera   *world.Camera
)

func update(image *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		camera.Angle -= 0.05
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		camera.Angle += 0.05
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		camera.Position.Add(geometry.Point{X: 0.02 * math.Cos(float64(camera.Angle)), Y: 0.02 * math.Sin(float64(camera.Angle))})
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		camera.Position.Add(geometry.Point{X: -0.02 * math.Cos(float64(camera.Angle)), Y: -0.02 * math.Sin(float64(camera.Angle))})
	}
	for i := int(0); i < world.Width; i++ {
		renderer.Render(i, screen)
	}
	p := make([]byte, world.Width*world.Heigth*4)
	for y := 0; y < world.Heigth; y++ {
		for x := 0; x < world.Width; x++ {
			position := y*world.Width + x
			p[(position*4 + 0)] = screen.Data[x][y][0]
			p[(position*4 + 1)] = screen.Data[x][y][1]
			p[(position*4 + 2)] = screen.Data[x][y][2]
			p[(position*4 + 3)] = 255
		}
	}
	image.ReplacePixels(p)
	ebitenutil.DebugPrint(image, fmt.Sprintf("FPS: %f", ebiten.CurrentFPS()))
	return nil
}

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	screen = &world.Screen{}
	camera = &world.Camera{
		Position: geometry.Point{
			X: 1.5,
			Y: 1.5,
		},
	}
	worldMap := world.NewMap()
	wallTexture := graphics.NewTexture("texture.png")
	renderer = &world.Renderer{
		Camera:      camera,
		WorldMap:    worldMap,
		WallTexture: wallTexture,
	}
	if err := ebiten.Run(update, world.Width, world.Heigth, world.Scale, "caster"); err != nil {
		panic(err)
	}
}
