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
	Camera   *Camera
	WorldMap *Map
}

func (r *Renderer) Render(column int, screen *Screen) {
	normalizedHeigth, textureOffset := r.castRay(column, Width)
	if textureOffset == 0 {
		textureOffset++
	}
	limitedHeigth := math.Min(normalizedHeigth, 1.0)
	columnHeigh := Heigth * limitedHeigth
	padding := (Heigth - columnHeigh) / 2
	for i := int(0); i < int(padding); i++ {
		screen.Data[column][i][0] = 0
		screen.Data[column][i][1] = 255
		screen.Data[column][i][2] = 255
	}
	for i := int(padding); i < int(padding+columnHeigh); i++ {
		screen.Data[column][i][0] = byte(255 / textureOffset)
		screen.Data[column][i][1] = byte(255 / textureOffset)
		screen.Data[column][i][2] = byte(255 / textureOffset)
	}
	for i := int(padding + columnHeigh); i < Heigth; i++ {
		screen.Data[column][i][0] = 0
		screen.Data[column][i][1] = 255
		screen.Data[column][i][2] = 255
	}
}

func (r *Renderer) castRay(column int, width int) (float64, uint) {
	relativeAngle := r.rayAngle(column, width)
	absoluteAngle := relativeAngle + r.Camera.Angle
	ray := geometry.NewRay(r.Camera.Position, absoluteAngle)

	for ray.Length <= r.WorldMap.MaxDistance {
		ray = ray.Grow()

		if r.WorldMap.HitTest(ray.End, ray.Angle) {
			var offset uint
			if ray.GrowingAxis == geometry.X {
				offset = uint(ray.End.Y*24) % 24
			} else if ray.GrowingAxis == geometry.Y {
				offset = uint(ray.End.X*24) % 24
			}
			projectedDistance := ray.Length * math.Cos(float64(relativeAngle))
			normalizedHeigth := 1.0 / projectedDistance
			return normalizedHeigth, offset
		}
	}

	return 0.0, 0
}

func (r *Renderer) rayAngle(column int, width int) geometry.Angle {
	relativePosition := (float64(column) / float64(width)) - 0.5
	virtualScreenPosition := relativePosition * relativeScreenSize
	return geometry.Angle(math.Atan(virtualScreenPosition / focalLength))
}
