package service

import (
	"calmoon/domain/entity"
	"calmoon/domain/repository"
	"errors"
)

type signService struct {
	repo repository.SignRepository
}

func NewSignService(repo repository.SignRepository) SignService {
	return &signService{repo: repo}
}

func (s *signService) GetAllSigns() ([]entity.Sign, error) {
	return s.repo.GetAll()
}

func (s *signService) GetSignByID(id int64) (*entity.Sign, error) {
	return s.repo.GetByID(id)
}

func (s *signService) GetSignByName(name string) (*entity.Sign, error) {
	return s.repo.GetByName(name)
}

func (s *signService) CreateSign(sign *entity.Sign) error {
	// validation here
	if sign.Name == "" {
		return errors.New("name cannot be empty")
	}
	return s.repo.Create(sign)
}

func (s *signService) UpdateSign(sign *entity.Sign) error {
	return s.repo.Update(sign)
}

func (s *signService) DeleteSign(id int64) error {
	return s.repo.Delete(id)
}
