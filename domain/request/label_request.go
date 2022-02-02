package request

import "github.com/google/uuid"

type LabelRequest struct {
	Name     string     `form:"name" json:"name" validate:"required"`
	ParentID *uuid.UUID `form:"parent_id" json:"parent_id"`
}
