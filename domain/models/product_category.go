package models

import "github.com/google/uuid"

type ProductCategory struct {
	ProductID  *uuid.UUID
	CategoryID *uuid.UUID
	Product    *Product
	Category   *Category
}
