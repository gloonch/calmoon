package repository

import "calmoon/domain/entity"

type SignRepository interface {
	GetAll() ([]entity.Sign, error)
	GetByID(id int64) (*entity.Sign, error)
	GetByName(name string) (*entity.Sign, error)
	Create(sign *entity.Sign) error
	Update(sign *entity.Sign) error
	Delete(id int64) error
}

type HouseRepository interface {
}

type PlanetRepository interface {
}
