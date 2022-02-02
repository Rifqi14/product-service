package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type IGenderRepository interface {
	Create(model models.Gender) (res models.Gender, err error)

	Update(model models.Gender) (res models.Gender, err error)

	Delete(model models.Gender) (err error)
}
