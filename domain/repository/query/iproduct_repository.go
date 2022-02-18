package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IProductRepository interface {
	List(search, orderby, sort, productName string, limit, offset, minPrice, maxPrice int64, brand []*uuid.UUID, product []*uuid.UUID, color []*uuid.UUID) (res []*models.Product, count int64, err error)

	Detail(productId uuid.UUID) (res *models.Product, err error)

	FindBy(column, operator string, value interface{}) (res []*models.Product, err error)

	All() (res []models.Product, err error)
}
