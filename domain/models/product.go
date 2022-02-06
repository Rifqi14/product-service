package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-auth-svc/domain/models"
	"gorm.io/gorm"
)

type Product struct {
	ID              uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();constraint:OnDelete:CASCADE;"`
	Code            int64     `gorm:"autoIncrement"`
	Name            string
	BrandID         *uuid.UUID `gorm:"type:uuid"`
	NormalPrice     int64      `gorm:"type:decimal"`
	StripePrice     int64      `gorm:"type:decimal"`
	DiscountPrice   int64      `gorm:"type:decimal"`
	FinalPrice      int64      `gorm:"type:decimal"`
	Description     *string    `gorm:"type:text"`
	Measurement     *string    `gorm:"type:text"`
	Length          int
	Width           int
	Height          int
	PoStatus        bool `gorm:"default:false"`
	PoDay           int
	IsDisplayed     bool       `gorm:"default:true"`
	CreatedBy       *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy       *uuid.UUID `gorm:"type:uuid"`
	DeletedBy       *uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time  `gorm:"<-:create"`
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Brand           *Brand                 `gorm:"foreignKey:BrandID;constraint:OnUpdate:Cascade,OnDelete:Cascade;"`
	Logs            []ProductLog           `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categories      []*Category            `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE;"`
	Labels          []*Label               `gorm:"many2many:product_labels;constraint:OnDelete:CASCADE;"`
	Materials       []*Material            `gorm:"many2many:product_materials;constraint:OnDelete:CASCADE;"`
	Genders         []*Gender              `gorm:"many2many:product_genders;constraint:OnDelete:CASCADE;"`
	Colors          []*Color               `gorm:"many2many:product_colors;constraint:OnDelete:CASCADE;"`
	ProductVariants []*Color               `gorm:"many2many:product_variants;constraint:OnDelete:CASCADE;"`
	ProductImages   []*Color               `gorm:"many2many:product_variant_images;constraint:OnDelete:CASCADE;"`
	Variants        []*ProductVariant      `gorm:"constraint:OnDelete:CASCADE;"`
	Images          []*ProductVariantImage `gorm:"constraint:OnDelete:CASCADE;"`
	Created         *models.AdminUser      `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:Set null,OnDelete:Set null;"`
	Updated         *models.AdminUser      `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:Set null,OnDelete:Set null;"`
	Deleted         *models.AdminUser      `gorm:"foreignKey:DeletedBy;constraint:OnUpdate:Set null,OnDelete:Set null;"`
}
