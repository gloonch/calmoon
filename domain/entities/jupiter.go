package entities

type Jupiter struct {
	Planet
}

func NewJupiter(id int64, position Position) *Jupiter {
	return &Jupiter{
		Planet: NewPlanet(id, "Jupiter", position),
	}
}
