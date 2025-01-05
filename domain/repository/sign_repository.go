package repository

import (
	"calmoon/domain/entity"
	"gorm.io/gorm"
)

type signRepository struct {
	db *gorm.DB
}

func NewSignRepository(db *gorm.DB) SignRepository {
	return &signRepository{db: db}
}

func (r *signRepository) GetAll() ([]entity.Sign, error) {
	var signs []entity.Sign
	if err := r.db.Find(&signs).Error; err != nil {
		return nil, err
	}
	return signs, nil
}

func (r *signRepository) GetByID(id int64) (*entity.Sign, error) {
	var sign entity.Sign
	if err := r.db.First(&sign, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &sign, nil
}

func (r *signRepository) GetByName(name string) (*entity.Sign, error) {
	var sign entity.Sign
	if err := r.db.First(&sign, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &sign, nil
}

func (r *signRepository) Create(sign *entity.Sign) error {
	return r.db.Create(sign).Error
}

func (r *signRepository) Update(sign *entity.Sign) error {
	return r.db.Save(sign).Error
}

func (r *signRepository) Delete(id int64) error {
	return r.db.Delete(&entity.Sign{}, id).Error
}
