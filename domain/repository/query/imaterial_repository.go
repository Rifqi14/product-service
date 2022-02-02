package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IMaterialRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Material, count int64, err error)

	Detail(materialId uuid.UUID) (res models.Material, err error)
}
