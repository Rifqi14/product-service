package query

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type IGenderRepository interface {
	List(search, orderBy, sort string, limit, offset int64) (res []models.Gender, count int64, err error)

	Detail(genderId uuid.UUID) (res models.Gender, err error)

	Parent(parentId uuid.UUID) (res []models.Gender, err error)
}
