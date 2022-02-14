package request

import "github.com/google/uuid"

type ProductRequest struct {
	Name             string                   `form:"product_name" json:"product_name" validate:"required"`
	BrandID          *uuid.UUID               `form:"brand_id" json:"brand_id" validate:"required"`
	Categories       []*uuid.UUID             `form:"categories" json:"categories"`
	Labels           []*uuid.UUID             `form:"labels" json:"labels"`
	Materials        []*uuid.UUID             `form:"materials" json:"materials"`
	Gender           []*uuid.UUID             `form:"genders" json:"genders"`
	NormalPrice      int64                    `form:"normal_price" json:"normal_price" validate:"ltfield=StripePrice,required"`
	StripePrice      int64                    `form:"stripe_price" json:"stripe_price" validate:"gtfield=NormalPrice,required"`
	Discount         int64                    `form:"discount" json:"discount" validate:"ltfield=StripePrice,required"`
	Variants         []*VariantRequest        `form:"variants" json:"variants"`
	Description      *string                  `form:"description" json:"description"`
	Measurement      *string                  `form:"measurement" json:"measurement"`
	PackageDimension *PackageDimensionRequest `json:"package_dimension" form:"package_dimension"`
	PreOrder         *PreOrderRequest         `form:"pre_order" json:"pre_order"`
}

type VariantRequest struct {
	Colors  *uuid.UUID              `form:"colors" json:"colors"`
	Images  []*VariantImagesRequest `form:"images" json:"images"`
	Details []*VariantDetailRequest `form:"details" json:"details"`
}

type VariantImagesRequest struct {
	ImageID *uuid.UUID `form:"image_id" json:"image_id"`
	Look    string     `form:"look" json:"look"`
}

type VariantDetailRequest struct {
	Size   int64  `form:"size" json:"size"`
	Stock  int64  `form:"stock" json:"stock"`
	Sku    string `form:"sku" json:"sku"`
	Status bool   `form:"status" json:"status"`
}

type PackageDimensionRequest struct {
	Length int64 `form:"length" json:"length"`
	Width  int64 `form:"width" json:"width"`
	Height int64 `form:"height" json:"height"`
}

type PreOrderRequest struct {
	Status      bool  `form:"status" json:"status"`
	PreOrderDay int64 `form:"po_day" json:"po_day" validate:"required_if=Status false,lte=30"`
}

type FilterProductRequest struct {
	Pagination
	ProductName string       `json:"product_name" form:"product_name" query:"product_name"`
	Product     []*uuid.UUID `json:"product_id" form:"product_id" query:"product_id"`
	Brand       []*uuid.UUID `json:"brand_id" form:"brand_id" query:"brand_id"`
	Color       []*uuid.UUID `json:"color_id" form:"color_id" query:"color_id"`
	MinPrice    int64        `json:"minimum_price" form:"minimum_price" query:"minimum_price"`
	MaxPrice    int64        `json:"maximum_price" form:"maximum_price" query:"maximum_price"`
}

type FilterQueryProductRequest struct {
	Pagination
	ProductName string   `json:"product_name" form:"product_name" query:"product_name"`
	Product     []string `json:"product_id" form:"product_id" query:"product_id"`
	Brand       []string `json:"brand_id" form:"brand_id" query:"brand_id"`
	Color       []string `json:"color_id" form:"color_id" query:"color_id"`
	MinPrice    int64    `json:"minimum_price" form:"minimum_price" query:"minimum_price"`
	MaxPrice    int64    `json:"maximum_price" form:"maximum_price" query:"maximum_price"`
}

type BannedProductRequest struct {
	Reason       string `json:"reason" form:"reason"`
	AttachmentID string `json:"attachment_id" form:"attachment_id"`
	Status       bool   `json:"status" form:"status"`
}
