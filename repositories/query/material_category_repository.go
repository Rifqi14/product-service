package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type MaterialCategoryRepository struct {
	DB *gorm.DB
}

func NewQueryMaterialCategoryRepository(db *gorm.DB) query.IMaterialCategoryRepository {
	return &MaterialCategoryRepository{DB: db}
}

func (repo MaterialCategoryRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.MaterialCategory, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Where("LOWER(name) like ?", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo MaterialCategoryRepository) Detail(materialCatId uuid.UUID) (res *models.MaterialCategory, err error) {
	tx := repo.DB

	err = tx.Find(&res, "id = ?", materialCatId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo MaterialCategoryRepository) All() (res []models.MaterialCategory, err error) {
	tx := repo.DB

	err = tx.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
