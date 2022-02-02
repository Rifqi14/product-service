package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type GenderRepository struct {
	DB *gorm.DB
}

func NewCommandGenderRepository(db *gorm.DB) command.IGenderRepository {
	return &GenderRepository{DB: db}
}

func (repo GenderRepository) Create(gender models.Gender) (res models.Gender, err error) {
	tx := repo.DB
	err = tx.Create(&gender).Error
	if err != nil {
		return res, err
	}
	return gender, nil
}

func (repo GenderRepository) Update(gender models.Gender) (res models.Gender, err error) {
	tx := repo.DB
	err = tx.Preload("Childs.Childs.Childs").Preload("Parent.Parent.Parent").Updates(&gender).Error
	if err != nil {
		return res, err
	}
	return gender, nil
}

func (repo GenderRepository) Delete(gender models.Gender) (err error) {
	tx := repo.DB
	err = tx.Delete(&gender).Error
	if err != nil {
		return err
	}
	return nil
}
