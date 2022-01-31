package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/models"
	"gorm.io/gorm"
)

type Category struct {
	ID                  uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name                string     `gorm:"not null"`
	ParentID            *uuid.UUID `gorm:"type:uuid"`
	Level               int64
	Path                string
	MobileBannerID      *uuid.UUID `gorm:"type:uuid"`
	WebsiteBannerID     *uuid.UUID `gorm:"type:uuid"`
	MobileHeroBannerID  *uuid.UUID `gorm:"type:uuid"`
	WebsiteHeroBannerID *uuid.UUID `gorm:"type:uuid"`
	CreatedBy           *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy           *uuid.UUID `gorm:"type:uuid"`
	DeletedBy           *uuid.UUID `gorm:"type:uuid"`
	CreatedAt           time.Time  `gorm:"<-:create"`
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
	Parent              *Category   `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Childs              []Category  `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MobileBanner        models.File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WebsiteBanner       models.File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MobileHeroBanner    models.File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WebsiteHeroBanner   models.File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
