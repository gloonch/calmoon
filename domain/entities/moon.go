package entities

import "time"

type Moon struct {
	Planet
}

func NewMoon(id int64, position Position, isRetrograde bool) *Moon {
	return &Moon{
		Planet: NewPlanet(id, "Moon", position),
	}
}

func (m *Moon) CalculatePosition(t time.Time) Position {
	newPos := m.Planet.CalculatePosition(t)

	if m.isRetrograde {
		newPos.Degree -= 0.5
	}

	m.UpdatePosition(newPos)
	return m.Position()
}
