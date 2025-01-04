package entities

type Venus struct {
	Planet
}

func NewVenus(id int64, position Position) *Venus {
	return &Venus{
		Planet: NewPlanet(id, "Venus", position),
	}
}
