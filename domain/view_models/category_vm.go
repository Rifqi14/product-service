package view_models

import (
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type CategoryListVm struct {
	No     int64              `json:"no"`
	Parent string             `json:"parent_category"`
	Child  string             `json:"child_category"`
	Banner view_models.FileVm `json:"banner"`
	ID     string             `json:"category_id"`
}

type CategoryDetailVm struct {
	ID                string             `json:"category_id"`
	Name              string             `json:"category"`
	Parent            CategoryParentVm   `json:"parent"`
	BannerMobile      view_models.FileVm `json:"banner_mobile"`
	BannerWebsite     view_models.FileVm `json:"banner_website"`
	HeroBannerMobile  view_models.FileVm `json:"hero_banner_mobile"`
	HeroBannerWebsite view_models.FileVm `json:"hero_banner_website"`
}

type CategoryParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type CategoryMediaVm struct {
	ID                string             `json:"category_id"`
	BannerMobile      view_models.FileVm `json:"banner_mobile"`
	BannerWebsite     view_models.FileVm `json:"banner_website"`
	HeroBannerMobile  view_models.FileVm `json:"hero_banner_mobile"`
	HeroBannerWebsite view_models.FileVm `json:"hero_banner_website"`
}

type CategoryVm struct {
	List   CategoryListVm   `json:"list_category"`
	Detail CategoryDetailVm `json:"detail_category"`
	Media  CategoryMediaVm  `json:"media_category"`
}

func NewCategoryVm() CategoryVm {
	return CategoryVm{}
}

func (vm CategoryVm) BuildList(categories []models.Category) (res []CategoryListVm) {
	for i, category := range categories {
		res = append(res, CategoryListVm{
			No:     int64(i + 1),
			ID:     category.ID.String(),
			Parent: category.Parent.Name,
			Child:  category.Name,
			Banner: view_models.NewFileVm().Build(category.MobileBanner),
		})
	}
	return res
}

func (vm CategoryVm) BuildDetail(category *models.Category) CategoryDetailVm {
	return CategoryDetailVm{
		ID:                category.ID.String(),
		Name:              category.Name,
		Parent:            vm.BuildParent(category.Parent),
		BannerMobile:      view_models.NewFileVm().Build(category.MobileBanner),
		BannerWebsite:     view_models.NewFileVm().Build(category.WebsiteBanner),
		HeroBannerMobile:  view_models.NewFileVm().Build(category.MobileHeroBanner),
		HeroBannerWebsite: view_models.NewFileVm().Build(category.WebsiteHeroBanner),
	}
}

func (vm CategoryVm) BuildParent(parent *models.Category) CategoryParentVm {
	return CategoryParentVm{
		ID:   parent.ID.String(),
		Name: parent.Name,
	}
}

func (vm CategoryVm) BuildMedia(category *models.Category) CategoryMediaVm {
	return CategoryMediaVm{
		ID:                category.ID.String(),
		BannerMobile:      view_models.NewFileVm().Build(category.MobileBanner),
		BannerWebsite:     view_models.NewFileVm().Build(category.WebsiteBanner),
		HeroBannerMobile:  view_models.NewFileVm().Build(category.MobileHeroBanner),
		HeroBannerWebsite: view_models.NewFileVm().Build(category.WebsiteHeroBanner),
	}
}
