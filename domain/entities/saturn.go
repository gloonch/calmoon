package entities

type Saturn struct {
	Planet
	isMalefic bool
}

func NewSaturn(id int64, position Position) *Saturn {
	return &Saturn{
		Planet:    NewPlanet(id, "Saturn", position),
		isMalefic: true,
	}
}

func (st *Saturn) IsMalefic() bool {
	return st.isMalefic
}
