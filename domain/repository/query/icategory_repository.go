package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type ICategoryRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Category, count int64, err error)

	Detail(categoryID uuid.UUID) (res models.Category, err error)

	Parent(parentId uuid.UUID) (res []models.Category, err error)
}
