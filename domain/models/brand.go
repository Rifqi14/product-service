package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gorm.io/gorm"
)

type Brand struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string    `gorm:"unique;not null"`
	Slug            string    `gorm:"not null"`
	EstablishedDate time.Time `gorm:"not null;type:date"`
	Title           string    `gorm:"not null;size:70"`
	Catchphrase     string    `gorm:"not null;size:140"`
	About           string    `gorm:"not null;type:text"`
	Status          string
	LogoID          *uuid.UUID `gorm:"type:uuid"`
	BannerWebID     *uuid.UUID `gorm:"type:uuid"`
	BannerMobileID  *uuid.UUID `gorm:"type:uuid"`
	CreatedBy       *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy       *uuid.UUID `gorm:"type:uuid"`
	DeletedBy       *uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time  `gorm:"<-:create"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	MediaSocials    []BrandMediaSocial `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logs            []*BrandLog        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logo            *models.File       `gorm:"foreignKey:LogoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BannerWeb       *models.File       `gorm:"foreignKey:BannerWebID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BannerMobile    *models.File       `gorm:"foreignKey:BannerMobileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Products        []*Product         `gorm:"foreignKey:BrandID;references:ID;"`
}

func (b *Brand) BeforeDelete(tx *gorm.DB) error {
	tx.First(&b, b.ID)
	if b.Name == "" {
		return errors.New(messages.DataNotFound)
	}
	var countSeller int64
	countProducts := tx.Model(&b).Association("Products").Count()
	tx.Table("seller_brands").Where("brand_id = ?", b.ID).Count(&countSeller)
	if countSeller > 0 || countProducts > 0 {
		return errors.New("data in used")
	}
	return nil
}
