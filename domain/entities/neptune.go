package entities

type Neptune struct {
	Planet
}

func NewNeptune(id int64, position Position) *Neptune {
	return &Neptune{
		Planet: NewPlanet(id, "Neptune", position),
	}
}
