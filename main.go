package main

import (
	"github.com/Ruenzuo/caster/geometry"
	"github.com/Ruenzuo/caster/world"
)

func main() {
	screen := &world.Screen{}
	camera := &world.Camera{
		Position: geometry.Point{
			X: 1.5,
			Y: 1.5,
		},
	}
	worldMap := world.NewMap()
	renderer := &world.Renderer{
		Camera:   camera,
		WorldMap: worldMap,
	}
	for i := int(0); i < world.Width; i++ {
		renderer.Render(i, screen)
	}
}
