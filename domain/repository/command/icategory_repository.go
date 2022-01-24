package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type ICategoryRepository interface {
	Create(model models.Category) (res models.Category, err error)

	Update(model models.Category) (res models.Category, err error)

	Delete(model models.Category) (err error)
}
