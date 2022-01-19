package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BrandMediaSocial struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrandID   uuid.UUID
	Brand     Brand     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type      string    `gorm:"not null;size:50"`
	Link      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
