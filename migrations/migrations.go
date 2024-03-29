package migrations

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Brand{}, &models.BrandLog{}, &models.BrandMediaSocial{}, &models.Category{}, &models.Color{}, &models.Gender{}, &models.Label{}, &models.MaterialCategory{}, &models.Material{}, &models.Product{}, &models.ProductLog{}, &models.ProductVariant{}, &models.ProductVariantImage{})
}
