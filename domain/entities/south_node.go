package entities

type SouthNode struct {
	Planet
}

func NewSouthNode(id int64, position Position) *SouthNode {
	return &SouthNode{
		Planet: NewPlanet(id, "SouthNode", position),
	}
}
