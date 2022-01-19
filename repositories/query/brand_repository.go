package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
)

type BrandRepository struct {
	DB *gorm.DB
}

func NewQueryBrandRepository(db *gorm.DB) query.IBrandRepository {
	return &BrandRepository{DB: db}
}

func (repo BrandRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Brand, count int64, err error) {
	db := repo.DB
	search = strings.ToLower(search)

	err = db.Where("LOWER(brands.name) like ?", "%"+search+"%").Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo BrandRepository) Detail(brandID uuid.UUID) (res models.Brand, err error) {
	db := repo.DB

	err = db.Preload("MediaSocials").Find(&res, "id = ?", brandID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
