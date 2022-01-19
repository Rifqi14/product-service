package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gorm.io/gorm"
)

type IBrandRepository interface {
	Create(model models.Brand, tx *gorm.DB) (res models.Brand, err error)

	Update(model models.Brand, tx *gorm.DB) (res models.Brand, err error)

	Delete(model models.Brand, tx *gorm.DB) (err error)
}
