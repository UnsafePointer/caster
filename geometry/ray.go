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

func (r *Ray) Grow() *Ray {
	rayOnNextXLine := r.growToNextXLine()
	rayOnNextYLine := r.growToNextYLine()

	if rayOnNextXLine.Length < rayOnNextYLine.Length {
		return rayOnNextXLine
	}
	return rayOnNextYLine
}

func (r *Ray) growToNextXLine() *Ray {
	deltaX := r.distanceToNextGridLine(X)
	deltaY := math.Tan(float64(r.Angle)) * deltaX
	return r.growWithDelta(deltaX, deltaY)
}

func (r *Ray) growToNextYLine() *Ray {
	deltaY := r.distanceToNextGridLine(Y)
	deltaX := deltaY / math.Tan(float64(r.Angle))
	return r.growWithDelta(deltaX, deltaY)
}

func (r *Ray) distanceToNextGridLine(axis Axis) float64 {
	position := r.End.Component(axis)
	direction := DirectionForAngleInAxis(r.Angle, axis)
	switch direction {
	case Increasing:
		return math.Floor(position) + 1.0 - position
	case Decreasing:
		return math.Ceil(position) - 1.0 - position
	}
	return 0.0
}

func (r *Ray) growWithDelta(deltaX float64, deltaY float64) *Ray {
	e := Point{
		X: r.End.X + deltaX,
		Y: r.End.Y + deltaY,
	}
	return NewRay2(r.Start, e, r.Angle)
}
