package entities

type Pluto struct {
	Planet
}

func NewPluto(id int64, position Position) *Pluto {
	return &Pluto{
		Planet: NewPlanet(id, "Pluto", position),
	}
}
