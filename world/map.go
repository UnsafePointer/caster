package world

import (
	"math"

	"github.com/Ruenzuo/caster/geometry"
)

type Map struct {
	Tiles       [][]bool
	MaxDistance float64
}

func NewMap() *Map {
	m := &Map{
		Tiles: [][]bool{
			{true, true, true, true, true, true},
			{true, false, true, true, false, true},
			{true, false, false, false, false, true},
			{true, false, false, false, false, true},
			{true, false, true, true, false, true},
			{true, true, true, true, true, true},
		},
		MaxDistance: 6 * 6,
	}
	return m
}

func (m *Map) HitTest(point geometry.Point, angle geometry.Angle) bool {
	x := positionComponent(&point, angle, geometry.X)
	y := positionComponent(&point, angle, geometry.Y)
	if x < 0 || x >= 6 || y < 0 || y >= 6 {
		return false
	}
	return m.Tiles[x][y]
}

func positionComponent(point *geometry.Point, angle geometry.Angle, axis geometry.Axis) int {
	pointComponent := point.Component(axis)

	if pointComponent == math.Trunc(pointComponent) {
		direction := geometry.DirectionForAngleInAxis(angle, axis)
		switch direction {
		case geometry.Increasing:
			break
		case geometry.Decreasing:
			pointComponent -= 1.0
		}
	}

	return int(pointComponent)
}
