package command

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type IProductRepository interface {
	Create(model models.Product) (res *models.Product, err error)

	Update(model models.Product) (res *models.Product, err error)

	Delete(model models.Product) (err error)
}
