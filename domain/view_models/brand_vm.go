package view_models

import (
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gorm.io/gorm"
)

type BrandListVm struct {
	No     int64  `json:"no"`
	ID     string `json:"brand_id"`
	Name   string `json:"brand_name"`
	Logo   string `json:"logo"`
	Owner  string `json:"owner"`
	Status string `json:"status"`
}

type BrandDetailVm struct {
	ID              string               `json:"brand_id"`
	Name            string               `json:"brand_name"`
	Logo            string               `json:"logo"`
	WebBanner       string               `json:"website_banner"`
	MobileBanner    string               `json:"mobile_banner"`
	CreatedAt       string               `json:"created_at"`
	Owner           string               `json:"owner"`
	EstablishedDate string               `json:"established_date"`
	About           string               `json:"about"`
	Platform        []BrandMediaSocialVm `json:"platform"`
}

type BrandVm struct {
	List   BrandListVm   `json:"list_brand"`
	Detail BrandDetailVm `json:"detail_brand"`
}

func NewBrandVm() BrandVm {
	return BrandVm{}
}

func (vm BrandVm) BuildList(model []models.Brand) (res []BrandListVm) {
	for i, brand := range model {
		res = append(res, BrandListVm{
			No:     int64(i + 1),
			ID:     brand.ID.String(),
			Name:   brand.Name,
			Logo:   brand.LogoID.String(),
			Owner:  brand.Title,
			Status: buildStatus(brand.DeletedAt),
		})
	}
	return res
}

func (vm BrandVm) BuildDetail(brand *models.Brand) BrandDetailVm {
	return BrandDetailVm{
		ID:              brand.ID.String(),
		Name:            brand.Name,
		Logo:            brand.LogoID.String(),
		WebBanner:       brand.BannerWebID.String(),
		MobileBanner:    brand.BannerMobileID.String(),
		CreatedAt:       brand.CreatedAt.Format("02-10-2006"),
		Owner:           brand.Title,
		EstablishedDate: brand.EstablishedDate.Format("02-10-2006"),
		About:           brand.About,
		Platform:        NewBrandMediaSocialVm().Build(brand.MediaSocials),
	}
}

func buildStatus(deletedAt gorm.DeletedAt) (status string) {
	if deletedAt.Valid {
		return "Non-Active"
	}
	return "Active"
}
