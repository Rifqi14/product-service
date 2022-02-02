package command

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/repository/command"
	"gorm.io/gorm"
)

type LabelRepository struct {
	DB *gorm.DB
}

func NewCommandLabelRepository(db *gorm.DB) command.ILabelRepository {
	return &LabelRepository{DB: db}
}

func (repo LabelRepository) Create(label models.Label) (res models.Label, err error) {
	tx := repo.DB
	err = tx.Create(&label).Error
	if err != nil {
		return res, err
	}
	return label, nil
}

func (repo LabelRepository) Update(label models.Label) (res models.Label, err error) {
	tx := repo.DB
	err = tx.Preload("Childs.Childs.Childs").Preload("Parent.Parent.Parent").Updates(&label).Error
	if err != nil {
		return res, err
	}
	return label, nil
}

func (repo LabelRepository) Delete(label models.Label) (err error) {
	tx := repo.DB
	err = tx.Updates(&label).Error
	if err != nil {
		return err
	}
	err = tx.Delete(&label).Error
	if err != nil {
		return err
	}
	return nil
}
