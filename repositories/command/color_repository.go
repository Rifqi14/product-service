package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type ColorRepository struct {
	DB *gorm.DB
}

func NewCommandColorRepository(db *gorm.DB) command.IColorRepository {
	return &ColorRepository{DB: db}
}

func (repo ColorRepository) Create(color models.Color) (res models.Color, err error) {
	tx := repo.DB
	err = tx.Create(&color).Error
	if err != nil {
		return res, err
	}
	return color, nil
}

func (repo ColorRepository) Update(color models.Color) (res models.Color, err error) {
	tx := repo.DB
	err = tx.Preload("Parent.Parent.Parent").Updates(&color).Error
	if err != nil {
		return res, err
	}
	return color, nil
}

func (repo ColorRepository) Delete(color models.Color) (err error) {
	tx := repo.DB
	err = tx.Updates(&color).Error
	if err != nil {
		return err
	}
	err = tx.Delete(&color).Error
	if err != nil {
		return err
	}
	return nil
}
