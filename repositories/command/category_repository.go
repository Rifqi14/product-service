package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCommandCategoryRepository(db *gorm.DB) command.ICategoryRepository {
	return &CategoryRepository{DB: db}
}

func (repo CategoryRepository) Create(category models.Category) (res models.Category, err error) {
	tx := repo.DB
	err = tx.Create(&category).Error
	if err != nil {
		return res, err
	}
	return category, nil
}

func (repo CategoryRepository) Update(category models.Category) (res models.Category, err error) {
	tx := repo.DB
	err = tx.Preload("Parent.Parent.Parent").Updates(&category).Error
	if err != nil {
		return res, err
	}
	return category, nil
}

func (repo CategoryRepository) Delete(category models.Category) (err error) {
	tx := repo.DB
	err = tx.Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
