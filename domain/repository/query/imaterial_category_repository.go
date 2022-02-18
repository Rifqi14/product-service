package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IMaterialCategoryRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.MaterialCategory, count int64, err error)

	Detail(materialCatId uuid.UUID) (res *models.MaterialCategory, err error)

	All() (res []models.MaterialCategory, err error)
}
