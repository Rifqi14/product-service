package view_models

import (
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type CategoryListVm struct {
	ID             string             `json:"category_id"`
	ParentID       string             `json:"parent_id"`
	Parent         string             `json:"parent_category"`
	Child          string             `json:"child_category"`
	Banner         view_models.FileVm `json:"banner"`
	CategoryBanner CategoryBannerVm   `json:"category_banner"`
}

type CategoryBannerVm struct {
	MobileBanner      view_models.FileVm `json:"mobile_banner"`
	WebsiteBanner     view_models.FileVm `json:"website_banner"`
	MobileHeroBanner  view_models.FileVm `json:"mobile_hero_banner"`
	WebsiteHeroBanner view_models.FileVm `json:"website_hero_banner"`
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

type CategoryExportVm struct {
	Name   string
	Parent string
}

type CategoryVm struct {
	List   CategoryListVm   `json:"list_category"`
	Detail CategoryDetailVm `json:"detail_category"`
	Media  CategoryMediaVm  `json:"media_category"`
}

func NewCategoryVm() *CategoryVm {
	return &CategoryVm{}
}

func (vm CategoryVm) BuildExport(categories []models.Category) (res []CategoryExportVm) {
	for _, category := range categories {
		var categoryVm CategoryExportVm
		categoryVm.Name = category.Name
		if category.Parent != nil {
			categoryVm.Parent = category.Parent.Name
		}
		res = append(res, categoryVm)
	}
	return res
}

func (vm CategoryVm) BuildList(categories []models.Category) (res []CategoryListVm) {
	for _, category := range categories {
		var banner view_models.FileVm
		var categoryBanner CategoryBannerVm
		var parent string
		var parentId string
		if category.ParentID != nil {
			parentId = category.ParentID.String()
		}
		if category.MobileBanner != nil {
			banner = view_models.NewFileVm().Build(*category.MobileBanner)
		}
		if category.Parent != nil {
			parent = category.Parent.Name
		}
		if category.MobileBanner != nil {
			categoryBanner.MobileBanner = view_models.NewFileVm().Build(*category.MobileBanner)
		}
		if category.WebsiteBanner != nil {
			categoryBanner.WebsiteBanner = view_models.NewFileVm().Build(*category.WebsiteBanner)
		}
		if category.MobileHeroBanner != nil {
			categoryBanner.MobileHeroBanner = view_models.NewFileVm().Build(*category.MobileHeroBanner)
		}
		if category.WebsiteHeroBanner != nil {
			categoryBanner.WebsiteHeroBanner = view_models.NewFileVm().Build(*category.WebsiteHeroBanner)
		}
		res = append(res, CategoryListVm{
			ID:             category.ID.String(),
			ParentID:       parentId,
			Parent:         parent,
			Child:          category.Name,
			Banner:         banner,
			CategoryBanner: categoryBanner,
		})
	}
	return res
}

func (vm CategoryVm) BuildDetail(category *models.Category) CategoryDetailVm {
	var bannerMobile view_models.FileVm
	var bannerWebsite view_models.FileVm
	var heroMobile view_models.FileVm
	var heroWebsite view_models.FileVm
	var parent CategoryParentVm
	if category.MobileBanner != nil {
		bannerMobile = view_models.NewFileVm().Build(*category.MobileBanner)
	}
	if category.WebsiteBanner != nil {
		bannerWebsite = view_models.NewFileVm().Build(*category.WebsiteBanner)
	}
	if category.MobileHeroBanner != nil {
		heroMobile = view_models.NewFileVm().Build(*category.MobileHeroBanner)
	}
	if category.WebsiteHeroBanner != nil {
		heroWebsite = view_models.NewFileVm().Build(*category.WebsiteHeroBanner)
	}
	if category.Parent != nil {
		parent = vm.BuildParent(*category.Parent)
	}
	return CategoryDetailVm{
		ID:                category.ID.String(),
		Name:              category.Name,
		Parent:            parent,
		BannerMobile:      bannerMobile,
		BannerWebsite:     bannerWebsite,
		HeroBannerMobile:  heroMobile,
		HeroBannerWebsite: heroWebsite,
	}
}

func (vm CategoryVm) BuildParent(parent models.Category) CategoryParentVm {
	return CategoryParentVm{
		ID:   parent.ID.String(),
		Name: parent.Name,
	}
}

func (vm CategoryVm) BuildMedia(category *models.Category) CategoryMediaVm {
	var bannerMobile view_models.FileVm
	var bannerWebsite view_models.FileVm
	var heroMobile view_models.FileVm
	var heroWebsite view_models.FileVm
	if category.MobileBanner != nil {
		bannerMobile = view_models.NewFileVm().Build(*category.MobileBanner)
	}
	if category.WebsiteBanner != nil {
		bannerWebsite = view_models.NewFileVm().Build(*category.WebsiteBanner)
	}
	if category.MobileHeroBanner != nil {
		heroMobile = view_models.NewFileVm().Build(*category.MobileHeroBanner)
	}
	if category.WebsiteHeroBanner != nil {
		heroWebsite = view_models.NewFileVm().Build(*category.WebsiteHeroBanner)
	}
	return CategoryMediaVm{
		ID:                category.ID.String(),
		BannerMobile:      bannerMobile,
		BannerWebsite:     bannerWebsite,
		HeroBannerMobile:  heroMobile,
		HeroBannerWebsite: heroWebsite,
	}
}
