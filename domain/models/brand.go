package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
	"gorm.io/gorm"
)

type Brand struct {
	ID                  uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name                string    `gorm:"unique;not null"`
	Slug                string    `gorm:"not null"`
	EstablishedDate     time.Time `gorm:"not null;type:date"`
	Title               string    `gorm:"not null;size:70"`
	Catchphrase         string    `gorm:"not null;size:140"`
	About               string    `gorm:"not null;type:text"`
	Status              string
	BannedReason        string     `gorm:"type:text"`
	BannedDocumentID    *uuid.UUID `gorm:"type:uuid"`
	BannedBy            *uuid.UUID `gorm:"type:uuid"`
	UnbannedReason      string
	UnbannedDocummentID *uuid.UUID `gorm:"type:uuid"`
	UnbannedBy          *uuid.UUID `gorm:"type:uuid"`
	LogoID              *uuid.UUID `gorm:"type:uuid"`
	BannerWebID         *uuid.UUID `gorm:"type:uuid"`
	BannerMobileID      *uuid.UUID `gorm:"type:uuid"`
	CreatedBy           *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy           *uuid.UUID `gorm:"type:uuid"`
	DeletedBy           *uuid.UUID `gorm:"type:uuid"`
	CreatedAt           time.Time  `gorm:"<-:create"`
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
	MediaSocials        []BrandMediaSocial `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logo                models.File        `gorm:"foreignKey:LogoID"`
	BannerWeb           models.File        `gorm:"foreignKey:BannerWebID"`
	BannerMobile        models.File        `gorm:"foreignKey:BannerMobileID"`
	BannedDocument      models.File        `gorm:"foreignKey:BannedDocumentID"`
	UnbannedDocument    models.File        `gorm:"foreignKey:UnabnnedDocumentID"`
}
