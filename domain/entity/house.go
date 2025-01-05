package entity

type House struct {
	ID               int64                  `gorm:"primary_key"`
	Number           int                    `json:"number"`
	AssociatedSignID int64                  `json:"sign_id"`
	RulingPlanetID   int64                  `json:"planet_id"`
	OccupyingPlanets []int64                `json:"occupying_planets"`
	UniqueAttributes map[string]interface{} `json:"unique_attributes"`
}
