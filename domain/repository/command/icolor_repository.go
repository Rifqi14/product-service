package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type IColorRepository interface {
	Create(model models.Color) (res models.Color, err error)

	Update(model models.Color) (res models.Color, err error)

	Delete(model models.Color) (err error)
}
