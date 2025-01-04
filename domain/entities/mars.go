package entities

import (
	"time"
)

type Mars struct {
	Planet
}

func NewMars(id int64, position Position) *Mars {
	return &Mars{
		Planet: NewPlanet(id, "Mars", position),
	}
}

func (m *Mars) CalculatePosition(t time.Time) Position {
	newPos := m.Planet.CalculatePosition(t)
	newPos.Degree += 0.1
	m.UpdatePosition(newPos)
	return newPos
}
