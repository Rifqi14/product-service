package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type BrandMediaSocialVm struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

func NewBrandMediaSocialVm() BrandMediaSocialVm {
	return BrandMediaSocialVm{}
}

func (vm BrandMediaSocialVm) Build(model []models.BrandMediaSocial) (res []BrandMediaSocialVm) {
	for _, brand := range model {
		res = append(res, BrandMediaSocialVm{
			Type: brand.Type,
			Link: brand.Link,
		})
	}
	return res
}
