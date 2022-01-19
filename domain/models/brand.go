package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	ID              uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string     `gorm:"unique;not null"`
	Slug            string     `gorm:"not null"`
	EstablishedDate time.Time  `gorm:"not null;type:date"`
	Title           string     `gorm:"not null;size:70"`
	Catchphrase     string     `gorm:"not null;size:140"`
	About           string     `gorm:"not null;type:text"`
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
}
