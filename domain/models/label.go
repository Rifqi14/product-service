package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Label struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string
	ParentID  *uuid.UUID `gorm:"type:uuid"`
	Level     int64
	Path      string
	CreatedBy *uuid.UUID `gorm:"type:uuid"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid"`
	DeletedBy *uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time  `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Parent    *Label  `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Childs    []Label `gorm:"foreignkey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
