package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IBrandRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Brand, count int64, err error)

	Detail(brandID uuid.UUID) (res models.Brand, err error)
}
