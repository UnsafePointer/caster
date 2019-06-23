package geometry

import "math"

type Ray struct {
	Start  Point
	End    Point
	Angle  Angle
	Length float64
}

func NewRay(start Point, angle Angle) *Ray {
	r := NewRay2(start, start, angle)
	return r
}

func NewRay2(start Point, end Point, angle Angle) *Ray {
	deltaX := end.X - start.X
	deltaY := end.Y - start.Y
	length := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	r := &Ray{
		Start:  start,
		End:    end,
		Angle:  angle,
		Length: length,
	}
	return r
}
