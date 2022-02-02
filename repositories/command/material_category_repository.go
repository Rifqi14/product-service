package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type MaterialCategoryRepository struct {
	DB *gorm.DB
}

func NewCommandMaterialCategoryRepository(db *gorm.DB) command.IMaterialCategoryRepository {
	return &MaterialCategoryRepository{DB: db}
}

func (repo MaterialCategoryRepository) Create(category models.MaterialCategory) (res models.MaterialCategory, err error) {
	tx := repo.DB
	err = tx.Create(&category).Error
	if err != nil {
		return res, err
	}
	return category, nil
}

func (repo MaterialCategoryRepository) Update(category models.MaterialCategory) (res models.MaterialCategory, err error) {
	tx := repo.DB
	err = tx.Updates(&category).Error
	if err != nil {
		return res, err
	}
	return category, nil
}

func (repo MaterialCategoryRepository) Delete(category models.MaterialCategory) (err error) {
	tx := repo.DB
	err = tx.Updates(&category).Error
	if err != nil {
		return err
	}
	err = tx.Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
