package world

import (
	"math"

	"github.com/Ruenzuo/caster/geometry"
)

const (
	relativeScreenSize float64 = 1.0
	focalLength        float64 = 0.75
)

type Renderer struct {
	camera *Camera
}

func (r *Renderer) Render(column int, screen *Screen) {
	r.castRay(column, Width)
}

func (r *Renderer) castRay(column int, width int) {
	relativeAngle := r.rayAngle(column, width)
	absoluteAngle := relativeAngle + r.camera.Direction
	geometry.NewRay(r.camera.Position, absoluteAngle)
}

func (r *Renderer) rayAngle(column int, width int) geometry.Angle {
	relativePosition := (float64(column) / float64(width)) - 0.5
	virtualScreenPosition := relativePosition * relativeScreenSize
	return geometry.Angle(math.Atan(virtualScreenPosition / focalLength))
}
