package entities

import "time"

type Sun struct {
	Planet
}

func NewSun(id int64, position Position) *Sun {
	return &Sun{
		Planet: NewPlanet(id, "Sun", position),
	}
}

func (s *Sun) CalculatePosition(t time.Time) Position {
	pos := s.Planet.CalculatePosition(t)
	pos.Degree += 0.05
	s.UpdatePosition(pos)
	return pos
}
