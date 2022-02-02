package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type IMaterialRepository interface {
	Create(model models.Material) (res models.Material, err error)

	Update(model models.Material) (res models.Material, err error)

	Delete(model models.Material) (err error)
}
