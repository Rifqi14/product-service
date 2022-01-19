package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type BrandRepository struct {
	DB *gorm.DB
}

func NewCommandBrandRepository(db *gorm.DB) command.IBrandRepository {
	return &BrandRepository{DB: db}
}

func (BrandRepository) Create(brand models.Brand, tx *gorm.DB) (res models.Brand, err error) {
	err = tx.Create(&brand).Error
	if err != nil {
		return res, err
	}
	return brand, nil
}

func (BrandRepository) Update(brand models.Brand, tx *gorm.DB) (res models.Brand, err error) {
	err = tx.Updates(&brand).Error
	if err != nil {
		return res, err
	}
	err = tx.Model(&brand).Association("MediaSocials").Replace(brand.MediaSocials)
	if err != nil {
		return res, err
	}
	return brand, nil
}

func (BrandRepository) Delete(brand models.Brand, tx *gorm.DB) (err error) {
	err = tx.Delete(&brand).Error
	if err != nil {
		return err
	}
	return nil
}
