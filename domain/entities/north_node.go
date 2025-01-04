package entities

type NorthNode struct {
	Planet
}

func NewNorthNode(id int64, position Position) *NorthNode {
	return &NorthNode{
		Planet: NewPlanet(id, "NorthNode", position),
	}
}
