package usecases

import (
	"calmoon/domain/entity"
	"calmoon/domain/service"
)

type SignUseCase interface {
	GetAllSigns() ([]entity.Sign, error)
}

type signUseCase struct {
	signService service.SignService
}

func NewSignUseCase(signService service.SignService) SignUseCase {
	return &signUseCase{signService: signService}
}

func (u *signUseCase) GetAllSigns() ([]entity.Sign, error) {
	return u.signService.GetAllSigns()
}
