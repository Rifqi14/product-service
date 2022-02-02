package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type GenderRepository struct {
	DB *gorm.DB
}

func NewQueryGenderRepository(db *gorm.DB) query.IGenderRepository {
	return &GenderRepository{DB: db}
}

func (repo GenderRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Gender, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Preload("Parent", "LOWER(name) like ?", "%"+search+"%").Or("LOWER(genders.name) like ?", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo GenderRepository) Detail(genderId uuid.UUID) (res models.Gender, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "id = ?", genderId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
