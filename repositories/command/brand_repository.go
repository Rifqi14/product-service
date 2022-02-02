package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (BrandRepository) Banned(banned models.BrandLog, tx *gorm.DB) (res models.BrandLog, err error) {
	err = tx.Create(&banned).Error
	if err != nil {
		return res, err
	}
	return banned, nil
}

func (BrandRepository) Update(brand models.Brand, tx *gorm.DB) (res models.Brand, err error) {
	err = tx.Preload(clause.Associations).Updates(&brand).Error
	if err != nil {
		return res, err
	}
	if len(brand.MediaSocials) > 0 {
		err = tx.Model(&brand).Association("MediaSocials").Replace(brand.MediaSocials)
		if err != nil {
			return res, err
		}
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
