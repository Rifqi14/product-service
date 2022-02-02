package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type MaterialRepository struct {
	DB *gorm.DB
}

func NewCommandMaterialRepository(db *gorm.DB) command.IMaterialRepository {
	return &MaterialRepository{DB: db}
}

func (repo MaterialRepository) Create(material models.Material) (res models.Material, err error) {
	tx := repo.DB
	err = tx.Create(&material).Error
	if err != nil {
		return res, err
	}
	return material, nil
}

func (repo MaterialRepository) Update(material models.Material) (res models.Material, err error) {
	tx := repo.DB
	err = tx.Preload("Parent.Parent.Parent").Updates(&material).Error
	if err != nil {
		return res, err
	}
	return material, nil
}

func (repo MaterialRepository) Delete(material models.Material) (err error) {
	tx := repo.DB
	err = tx.Delete(&material).Error
	if err != nil {
		return err
	}
	return nil
}
