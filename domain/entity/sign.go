package entity

type Sign struct {
	ID               int64                  `gorm:"primary_key"`
	Name             string                 `json:"name"`
	Element          string                 `json:"element"`
	Modality         string                 `json:"modality"`
	RulingPlanetID   int64                  `json:"ruling_planet"`
	UniqueAttributes map[string]interface{} `json:"unique_attributes"`
}
