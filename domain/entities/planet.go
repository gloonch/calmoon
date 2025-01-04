package entities

import "time"

type Position struct {
	Degree   float64
	Sign     string
	HouseNum int
}

type Planet struct {
	id           int64
	name         string
	isRetrograde bool
	position     Position
}

func NewPlanet(id int64, name string, position Position) Planet {
	return Planet{
		id:       id,
		name:     name,
		position: position,
	}
}

func (p *Planet) ID() int64 {
	return p.id
}

func (p *Planet) Name() string {
	return p.name
}

func (mc *Mercury) IsRetrograde() bool {
	return mc.isRetrograde
}

func (p *Planet) Position() Position {
	return p.position
}

func (p *Planet) UpdatePosition(newPos Position) {
	p.position = newPos
}

func (p *Planet) CalculatePosition(t time.Time) Position {

	return p.position
}
