package models

import (
	"time"

	"github.com/google/uuid"
	adminModel "gitlab.com/s2.1-backend/shm-auth-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
	"gorm.io/gorm"
)

type ProductLog struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Reason       string    `gorm:"type:text"`
	Status       string
	ProductID    *uuid.UUID `gorm:"type:uuid"`
	AttachmentID *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy    *uuid.UUID `gorm:"type:uuid"`
	CreatedAt    time.Time  `gorm:"<-:create"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Product      Product               `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Attachment   *models.File          `gorm:"foreignkey:AttachmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Verifier     *adminModel.AdminUser `gorm:"foreignkey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
