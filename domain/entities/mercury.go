package entities

import "time"

type Mercury struct {
	Planet
}

func NewMercury(id int64, position Position, retrograde bool) *Mercury {
	return &Mercury{
		Planet: NewPlanet(id, "Mercury", position),
	}
}

func (mc *Mercury) CalculatePosition(t time.Time) Position {
	pos := mc.Planet.CalculatePosition(t)
	if mc.isRetrograde {
		pos.Degree -= 0.2
	}
	mc.UpdatePosition(pos)
	return pos
}
