package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type ILabelRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Label, count int64, err error)

	Detail(labelId uuid.UUID) (res models.Label, err error)
}
