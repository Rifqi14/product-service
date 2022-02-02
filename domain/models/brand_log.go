package models

import (
	"time"

	"github.com/google/uuid"
	adminModels "gitlab.com/s2.1-backend/shm-auth-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
	"gorm.io/gorm"
)

type BrandLog struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrandID      uuid.UUID `gorm:"type:uuid"`
	Reason       string    `gorm:"type:text"`
	Status       string
	AttachmentID *uuid.UUID `gorm:"type:uuid"`
	VerifierID   uuid.UUID  `gorm:"type:uuid"`
	CreatedAt    time.Time  `gorm:"<-:create"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Brand        Brand                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attachment   *models.File          `gorm:"foreignKey:AttachmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Verifier     adminModels.AdminUser `gorm:"foreignKey:VerifierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
