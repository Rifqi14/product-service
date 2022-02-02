package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IColorRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Color, count int64, err error)

	Detail(colorID uuid.UUID) (res models.Color, err error)

	Parent(parentId uuid.UUID) (res []models.Color, err error)
}
