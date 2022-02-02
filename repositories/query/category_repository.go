package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewQueryCategoryRepository(db *gorm.DB) query.ICategoryRepository {
	return &CategoryRepository{DB: db}
}

func (repo CategoryRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Category, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Preload("Parent", "LOWER(name) like ?", "%"+search+"%").Or("LOWER(categories.name) like ?", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo CategoryRepository) Detail(categoryID uuid.UUID) (res models.Category, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "id = ?", categoryID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo CategoryRepository) Parent(parentId uuid.UUID) (res []models.Category, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "parent_id = ?", parentId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
