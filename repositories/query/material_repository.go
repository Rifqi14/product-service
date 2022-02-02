package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type MaterialRepository struct {
	DB *gorm.DB
}

func NewQueryMaterialRepository(db *gorm.DB) query.IMaterialRepository {
	return &MaterialRepository{DB: db}
}

func (repo MaterialRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Material, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Preload("Parent", "LOWER(name) like ?", "%"+search+"%").Where("LOWER(materials.name) like ?", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo MaterialRepository) Detail(materialId uuid.UUID) (res models.Material, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "id = ?", materialId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
