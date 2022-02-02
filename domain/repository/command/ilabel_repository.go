package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type ILabelRepository interface {
	Create(model models.Label) (res models.Label, err error)

	Update(model models.Label) (res models.Label, err error)

	Delete(model models.Label) (err error)
}
