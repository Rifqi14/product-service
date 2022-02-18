package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewCommandProductRepository(db *gorm.DB) command.IProductRepository {
	return &ProductRepository{DB: db}
}

func (repo ProductRepository) Create(product models.Product) (res *models.Product, err error) {
	tx := repo.DB
	err = tx.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo ProductRepository) Update(product models.Product) (res *models.Product, err error) {
	tx := repo.DB
	err = tx.Updates(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo ProductRepository) Delete(product models.Product) (err error) {
	tx := repo.DB
	_, err = repo.Update(product)
	if err != nil {
		return err
	}
	err = tx.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
