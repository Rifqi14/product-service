package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type ColorRepository struct {
	DB *gorm.DB
}

func NewQueryColorRepository(db *gorm.DB) query.IColorRepository {
	return &ColorRepository{DB: db}
}

func (repo ColorRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Color, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Preload("Parent", "LOWER(name) like ? OR LOWER(rgb_code) like ?", "%"+search+"%", "%"+search+"%").Or("LOWER(colors.name) like ? OR LOWER(colors.rgb_code) like ?", "%"+search+"%", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo ColorRepository) Detail(colorID uuid.UUID) (res models.Color, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "id = ?", colorID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
