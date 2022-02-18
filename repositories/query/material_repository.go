package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err = tx.Joins("left join materials as parent on parent.id = materials.parent_id").Where("lower(materials.name) like ? or lower(parent.name) like ?", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return res, 0, err
	}

	err = tx.Joins("left join materials as parent on parent.id = materials.parent_id").Where("lower(materials.name) like ? or lower(parent.name) like ?", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Find(&models.Material{}).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo MaterialRepository) Detail(materialId uuid.UUID) (res *models.Material, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Find(&res, "id = ?", materialId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo MaterialRepository) Parent(parentId uuid.UUID) (res []models.Material, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "parent_id = ?", parentId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo MaterialRepository) GetBy(column, operator string, value interface{}) (res []*models.Material, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Find(&res, column+" "+operator+" ?", value).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo MaterialRepository) All() (res []models.Material, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
