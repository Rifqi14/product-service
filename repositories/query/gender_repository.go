package query

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err = tx.Joins("LEFT JOIN genders as parent on parent.id = genders.parent_id and parent.deleted_at is null").Where("lower(genders.name) like ? or lower(parent.name) like ?", "%"+search+"%", "%"+search+"%").Preload(clause.Associations).Order(orderBy + " " + sort).Limit(int(limit)).Offset(int(offset)).Find(&res).Limit(-1).Count(&count).Error
	if err != nil {
		return res, count, err
	}
	return res, count, nil
}

func (repo GenderRepository) Detail(genderId uuid.UUID) (res models.Gender, err error) {
	tx := repo.DB

	err = tx.Preload(clause.Associations).Find(&res, "id = ?", genderId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo GenderRepository) Parent(parentId uuid.UUID) (res []models.Gender, err error) {
	tx := repo.DB

	err = tx.Preload("Parent").Find(&res, "parent_id = ?", parentId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (repo GenderRepository) All() (res []models.Gender, err error) {
	tx := repo.DB

	err = tx.Preload("Parent.Parent.Parent.Parent.Parent").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
