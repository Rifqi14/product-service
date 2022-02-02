package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err = tx.Joins("LEFT JOIN colors as parent ON parent.id = colors.parent_id").Where("LOWER(colors.name) like ? OR LOWER(colors.rgb_code) like ? OR LOWER(parent.name) like ? OR LOWER(parent.rgb_code) like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return res, count, err
	}
	err = tx.Joins("LEFT JOIN colors as parent ON parent.id = colors.parent_id").Where("LOWER(colors.name) like ? OR LOWER(colors.rgb_code) like ? OR LOWER(parent.name) like ? OR LOWER(parent.rgb_code) like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Find(&models.Color{}).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo ColorRepository) Detail(colorID uuid.UUID) (res models.Color, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Find(&res, "id = ?", colorID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo ColorRepository) Parent(parentId uuid.UUID) (res []models.Color, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "parent_id = ?", parentId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
