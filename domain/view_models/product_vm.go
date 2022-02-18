package view_models

import (
	"strconv"
	"strings"
	"time"

	adminVm "gitlab.com/s2.1-backend/shm-auth-svc/domain/view_models"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type ProductVm struct {
	ID            string                   `json:"product_id"`
	Code          int64                    `json:"product_code"`
	Name          string                   `json:"product_name"`
	Brand         BrandDetailVm            `json:"brand"`
	NormalPrice   int64                    `json:"normal_price"`
	StripePrice   int64                    `json:"stripe_price"`
	DiscountPrice int64                    `json:"discount_price"`
	FinalPrice    int64                    `json:"final_price"`
	Description   *string                  `json:"description"`
	Measurement   *string                  `json:"measurement"`
	Dimension     *PackageDimensionVm      `json:"package_dimension"`
	PreOrder      *PreOrderVm              `json:"pre_order"`
	IsDisplayed   bool                     `json:"is_displayed"`
	Logs          []*ProductLogVm          `json:"logs"`
	Categories    []*CategoryDetailVm      `json:"categories"`
	Labels        []*LabelDetailVm         `json:"labels"`
	Materials     []*MaterialDetailVm      `json:"materials"`
	Genders       []*GenderDetailVm        `json:"genders"`
	Colors        []*ColorDetailVm         `json:"colors"`
	Variants      []*ProductVariantVm      `json:"variant"`
	Images        []*ProductVariantImageVm `json:"images"`
	CreatedBy     *adminVm.AdminVm         `json:"created_by"`
	UpdatedBy     *adminVm.AdminVm         `json:"updated_by"`
	DeletedBy     *adminVm.AdminVm         `json:"deleted_by"`
}

type ProductVariantVm struct {
	Colors    *ColorDetailVm `json:"colors"`
	Size      int64          `json:"size"`
	Stock     int64          `json:"stock"`
	Sku       string         `json:"sku"`
	Status    bool           `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
}

type ProductVariantImageVm struct {
	Colors *ColorDetailVm `json:"colors"`
	Look   string         `json:"look"`
	Image  *fileVm.FileVm `json:"image"`
}

type PackageDimensionVm struct {
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type PreOrderVm struct {
	PoStatus bool `json:"po_status"`
	PoDay    int  `json:"po_day"`
}

type ProductExportVm struct {
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	Label       string `json:"label"`
	Material    string `json:"material"`
	Gender      string `json:"gender"`
	NormalPrice string `json:"normal_price"`
	StripePrice string `json:"stripe_price"`
	Discount    string `json:"discount"`
	Color       string `json:"color"`
	Size        string `json:"size"`
	Stock       string `json:"stock"`
	SKU         string `json:"sku"`
	Description string `json:"description"`
	Length      string `json:"length"`
	Width       string `json:"width"`
	Height      string `json:"height"`
	PoDay       string `json:"po_day"`
}

type BuildProductVm struct {
	List   []*ProductVm      `json:"list_product"`
	Detail *ProductVm        `json:"detail_product"`
	Export []ProductExportVm `json:"export_product"`
}

func NewProductVm() *BuildProductVm {
	return &BuildProductVm{}
}

func (vm BuildProductVm) BuildExport(products []models.Product) (res []ProductExportVm) {
	var categories []string
	var labels []string
	var materials []string
	var genders []string
	for _, product := range products {
		for _, variant := range product.Variants {
			var productVm ProductExportVm
			productVm.Name = product.Name
			if product.Brand != nil {
				productVm.Brand = product.Brand.Name
			}
			for _, category := range product.Categories {
				categories = append(categories, category.Name)
			}
			for _, label := range product.Labels {
				labels = append(labels, label.Name)
			}
			for _, material := range product.Materials {
				materials = append(materials, material.Name)
			}
			for _, gender := range product.Genders {
				genders = append(genders, gender.Name)
			}
			productVm.Category = strings.Join(categories, ", ")
			productVm.Label = strings.Join(labels, ", ")
			productVm.Material = strings.Join(materials, ", ")
			productVm.Gender = strings.Join(genders, ", ")
			productVm.NormalPrice = strconv.Itoa(int(product.NormalPrice))
			productVm.StripePrice = strconv.Itoa(int(product.StripePrice))
			productVm.Discount = strconv.Itoa(int(product.DiscountPrice))
			productVm.Color = variant.Color.Name
			productVm.Size = strconv.Itoa(int(variant.Size))
			productVm.Stock = strconv.Itoa(int(variant.Stock))
			productVm.SKU = *variant.Sku
			if product.Description != nil {
				productVm.Description = *product.Description
			}
			productVm.Length = strconv.Itoa(int(product.Length))
			productVm.Width = strconv.Itoa(int(product.Width))
			productVm.Height = strconv.Itoa(int(product.Height))
			productVm.PoDay = strconv.Itoa(int(product.PoDay))
			res = append(res, productVm)
		}
	}
	return res
}

func (vm BuildProductVm) BuildList(products []*models.Product) (res []*ProductVm) {
	if len(products) > 0 {
		for _, product := range products {
			var brand BrandDetailVm
			var deleted adminVm.AdminVm
			if product.Brand != nil {
				brand = NewBrandVm().BuildDetail(product.Brand)
			}
			created := adminVm.NewAdminVm().Build(product.Created)
			updated := adminVm.NewAdminVm().Build(product.Updated)
			if product.Deleted != nil {
				deleted = adminVm.NewAdminVm().Build(product.Deleted)
			}
			res = append(res, &ProductVm{
				ID:            product.ID.String(),
				Code:          product.Code,
				Name:          product.Name,
				Brand:         brand,
				NormalPrice:   product.NormalPrice,
				StripePrice:   product.StripePrice,
				DiscountPrice: product.DiscountPrice,
				FinalPrice:    product.FinalPrice,
				Description:   product.Description,
				Measurement:   product.Measurement,
				Dimension: &PackageDimensionVm{
					Length: product.Length,
					Width:  product.Width,
					Height: product.Height,
				},
				PreOrder: &PreOrderVm{
					PoStatus: product.PoStatus,
					PoDay:    product.PoDay,
				},
				IsDisplayed: *product.IsDisplayed,
				Logs:        NewProductLogVm().BuildList(product.Logs),
				Categories:  vm.buildProductCategories(product),
				Labels:      vm.buildProductLabels(product),
				Materials:   vm.buildProductMaterials(product),
				Genders:     vm.buildProductGenders(product),
				Colors:      vm.buildProductColors(product),
				Variants:    vm.buildVariant(product),
				Images:      vm.buildImages(product),
				CreatedBy:   &created,
				UpdatedBy:   &updated,
				DeletedBy:   &deleted,
			})
		}
	}
	return res
}

func (vm BuildProductVm) BuildDetail(product *models.Product) (res *ProductVm) {
	if product != nil {
		var brand BrandDetailVm
		var deleted adminVm.AdminVm
		if product.Brand != nil {
			brand = NewBrandVm().BuildDetail(product.Brand)
		}
		created := adminVm.NewAdminVm().Build(product.Created)
		updated := adminVm.NewAdminVm().Build(product.Updated)
		if product.Deleted != nil {
			deleted = adminVm.NewAdminVm().Build(product.Deleted)
		}
		res = &ProductVm{
			ID:            product.ID.String(),
			Code:          product.Code,
			Name:          product.Name,
			Brand:         brand,
			NormalPrice:   product.NormalPrice,
			StripePrice:   product.StripePrice,
			DiscountPrice: product.DiscountPrice,
			FinalPrice:    product.FinalPrice,
			Description:   product.Description,
			Measurement:   product.Measurement,
			Dimension: &PackageDimensionVm{
				Length: product.Length,
				Width:  product.Width,
				Height: product.Height,
			},
			PreOrder: &PreOrderVm{
				PoStatus: product.PoStatus,
				PoDay:    product.PoDay,
			},
			IsDisplayed: *product.IsDisplayed,
			Logs:        NewProductLogVm().BuildList(product.Logs),
			Categories:  vm.buildProductCategories(product),
			Labels:      vm.buildProductLabels(product),
			Materials:   vm.buildProductMaterials(product),
			Genders:     vm.buildProductGenders(product),
			Colors:      vm.buildProductColors(product),
			Variants:    vm.buildVariant(product),
			Images:      vm.buildImages(product),
			CreatedBy:   &created,
			UpdatedBy:   &updated,
			DeletedBy:   &deleted,
		}
	}
	return res
}

func (vm BuildProductVm) buildVariant(product *models.Product) (res []*ProductVariantVm) {
	if len(product.Variants) > 0 {
		for _, variant := range product.Variants {
			color := NewColorVm().BuildDetail(&variant.Color)
			res = append(res, &ProductVariantVm{
				Colors:    &color,
				Size:      variant.Size,
				Stock:     variant.Stock,
				Sku:       *variant.Sku,
				Status:    variant.Status,
				CreatedAt: variant.CreatedAt,
			})
		}
	}
	return res
}

func (vm BuildProductVm) buildImages(product *models.Product) (res []*ProductVariantImageVm) {
	if len(product.Images) > 0 {
		for _, image := range product.Images {
			var imageVm fileVm.FileVm
			if image.Image != nil {
				imageVm = fileVm.FileVm{
					ID:   image.ImageID.String(),
					Name: image.Image.Name,
					Ext:  image.Image.Extension,
					Path: image.Image.Path,
				}
			}
			color := NewColorVm().BuildDetail(&image.Color)
			res = append(res, &ProductVariantImageVm{
				Colors: &color,
				Look:   image.Look,
				Image:  &imageVm,
			})
		}
	}
	return res
}

func (vm BuildProductVm) buildProductCategories(product *models.Product) (res []*CategoryDetailVm) {
	if len(product.Categories) > 0 {
		for _, category := range product.Categories {
			categoryVm := NewCategoryVm().BuildDetail(category)
			res = append(res, &categoryVm)
		}
	}
	return res
}

func (vm BuildProductVm) buildProductLabels(product *models.Product) (res []*LabelDetailVm) {
	if len(product.Labels) > 0 {
		for _, label := range product.Labels {
			labelVm := NewLabelVm().BuildDetail(label)
			res = append(res, &labelVm)
		}
	}
	return res
}

func (vm BuildProductVm) buildProductMaterials(product *models.Product) (res []*MaterialDetailVm) {
	if len(product.Materials) > 0 {
		for _, material := range product.Materials {
			materialVm := NewMaterialVm().BuildDetail(material)
			res = append(res, &materialVm)
		}
	}
	return res
}

func (vm BuildProductVm) buildProductGenders(product *models.Product) (res []*GenderDetailVm) {
	if len(product.Genders) > 0 {
		for _, gender := range product.Genders {
			genderVm := NewGenderVm().BuildDetail(gender)
			res = append(res, &genderVm)
		}
	}
	return res
}

func (vm BuildProductVm) buildProductColors(product *models.Product) (res []*ColorDetailVm) {
	if len(product.Colors) > 0 {
		for _, color := range product.Colors {
			colorVm := NewColorVm().BuildDetail(color)
			res = append(res, &colorVm)
		}
	}
	return res
}
