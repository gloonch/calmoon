package entity

import "time"

type PlanetRelation int

const (
	Friend PlanetRelation = iota
	Enemy
	Neutral
)

type Position struct {
	Degree    float64   `json:"degree"`
	Sign      string    `json:"sign"`
	HouseNum  int       `json:"house_num"`
	Timestamp time.Time `json:"timestamp"`
}

type Planet struct {
	ID                 int64                    `gorm:"primary_key"`
	Name               string                   `json:"name"`
	IsRetrograde       bool                     `json:"is_retrograde"`
	Element            string                   `json:"element"`
	Modality           string                   `json:"modality"`
	Position           Position                 `json:"position"`
	ExaltationSignID   int64                    `json:"exaltation_sign_id"`
	DebilitationSignID int64                    `json:"debilitation_sign_id"`
	RulingSigns        []int64                  `json:"ruling_sign_id"`
	Relations          map[int64]PlanetRelation `json:"relations"`
	UniqueAttributes   map[string]interface{}   `json:"unique_attributes"`
}

func (p *Planet) UpdatePosition(newPosition Position) {
	p.Position = newPosition
}

func (p *Planet) IsFriendWith(other Planet) bool {
	relation, exists := p.Relations[other.ID]
	return exists && relation == Friend
}
