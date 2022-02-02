package request

import "github.com/google/uuid"

type ColorRequest struct {
	Name     string     `form:"name" json:"name" validate:"required"`
	RgbCode  string     `form:"rgb_code" json:"rgb_code" validate:"required"`
	ParentID *uuid.UUID `form:"parent_id" json:"parent_id"`
}
