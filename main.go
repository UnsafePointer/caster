package main

import "github.com/Ruenzuo/caster/world"

func main() {
	screen := &world.Screen{}
	for i := int(0); i < world.Width; i++ {
		world.Render(i, screen)
	}
}
