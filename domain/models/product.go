package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name          string
	BrandID       *uuid.UUID `gorm:"type:uuid"`
	NormalPrice   *int64     `gorm:"type:decimal"`
	StripePrice   *int64     `gorm:"type:decimal"`
	DiscountPrice *int64     `gorm:"type:decimal"`
	FinalPrice    *int64     `gorm:"type:decimal"`
	Description   *string    `gorm:"type:text"`
	Measurement   *string    `gorm:"type:text"`
	Length        *int
	Width         *int
	Height        *int
	PoStatus      bool `gorm:"default:false"`
	PoDay         *int
	IsDisplayed   bool       `gorm:"default:true"`
	CreatedBy     *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy     *uuid.UUID `gorm:"type:uuid"`
	DeletedBy     *uuid.UUID `gorm:"type:uuid"`
	CreatedAt     time.Time  `gorm:"<-:create"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Logs          []ProductLog `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
