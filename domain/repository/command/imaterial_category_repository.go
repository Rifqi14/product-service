package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type IMaterialCategoryRepository interface {
	Create(model models.MaterialCategory) (res models.MaterialCategory, err error)

	Update(model models.MaterialCategory) (res models.MaterialCategory, err error)

	Delete(model models.MaterialCategory) (err error)
}
