package service

import "calmoon/domain/entity"

type SignService interface {
	GetAllSigns() ([]entity.Sign, error)
	GetSignByID(id int64) (*entity.Sign, error)
	GetSignByName(name string) (*entity.Sign, error)
	CreateSign(sign *entity.Sign) error
	UpdateSign(sign *entity.Sign) error
	DeleteSign(id int64) error
}
