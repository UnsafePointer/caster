package world

import (
	"fmt"
	"math"

	"github.com/Ruenzuo/caster/geometry"
)

const (
	relativeScreenSize float64 = 1.0
	focalLength        float64 = 0.75
)

type Renderer struct {
	Camera   *Camera
	WorldMap *Map
}

func (r *Renderer) Render(column int, screen *Screen) {
	r.castRay(column, Width)
}

func (r *Renderer) castRay(column int, width int) {
	relativeAngle := r.rayAngle(column, width)
	absoluteAngle := relativeAngle + r.Camera.Angle
	ray := geometry.NewRay(r.Camera.Position, absoluteAngle)

	for ray.Length <= r.WorldMap.MaxDistance {
		ray = ray.Grow()

		if r.WorldMap.HitTest(ray.End, ray.Angle) {
			length := ray.Length * math.Cos(float64(relativeAngle))
			println(fmt.Sprintf("Wall at column: %d with length: %f", column, length))
		}
	}
}

func (r *Renderer) rayAngle(column int, width int) geometry.Angle {
	relativePosition := (float64(column) / float64(width)) - 0.5
	virtualScreenPosition := relativePosition * relativeScreenSize
	return geometry.Angle(math.Atan(virtualScreenPosition / focalLength))
}
