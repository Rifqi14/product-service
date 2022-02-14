package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
)

type ProductVariant struct {
	ProductID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	ColorID   uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Size      int64     `gorm:"primaryKey;type:numeric"`
	Stock     int64
	Sku       *string
	Status    bool
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	Product   Product
	Color     Color
}

type ProductVariantImage struct {
	ProductID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	ColorID   uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Look      string    `gorm:"primaryKey"`
	ImageID   *uuid.UUID
	Image     *models.File `gorm:"foreignKey:ImageID;constraint:OnUpdate:restrict,OnDelete:restrict;"`
	CreatedAt time.Time    `gorm:"<-:create"`
	UpdatedAt time.Time
	Product   Product
	Color     Color
}
