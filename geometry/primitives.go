package geometry

import "math"

type Angle float64

type Axis int

const (
	X Axis = iota
	Y Axis = iota
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Component(axis Axis) float64 {
	switch axis {
	case X:
		return p.X
	case Y:
		return p.Y
	}
	return 0.0
}

type Direction bool

const (
	Increasing Direction = true
	Decreasing Direction = false
)

func DirectionForAngleInAxis(angle Angle, axis Axis) Direction {
	switch axis {
	case X:
		if math.Cos(float64(angle)) > 0 {
			return Increasing
		}
		return Decreasing
	case Y:
		if math.Sin(float64(angle)) > 0 {
			return Increasing
		}
		return Decreasing
	}
	return Decreasing
}
