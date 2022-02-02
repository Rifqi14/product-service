package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LabelRepository struct {
	DB *gorm.DB
}

func NewQueryLabelRepository(db *gorm.DB) query.ILabelRepository {
	return &LabelRepository{DB: db}
}

func (repo LabelRepository) List(search, orderBy, sort string, limit, offset int64) (res []models.Label, count int64, err error) {
	tx := repo.DB
	search = strings.ToLower(search)

	err = tx.Joins("left join labels as parent on parent.id = labels.parent_id").Where("lower(labels.name) like ? or lower(parent.name) like ?", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Error
	if err != nil {
		return res, 0, err
	}

	err = tx.Joins("left join labels as parent on parent.id = labels.parent_id").Where("lower(labels.name) like ? or lower(parent.name) like ?", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Find(&models.Label{}).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo LabelRepository) Detail(labelId uuid.UUID) (res *models.Label, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "id = ?", labelId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo LabelRepository) Parent(parentId uuid.UUID) (res []models.Label, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "parent_id = ?", parentId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
