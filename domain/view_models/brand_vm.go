package view_models

import (
	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type BrandFullVm struct {
	ID              string               `json:"brand_id"`
	Name            string               `json:"name"`
	Slug            string               `json:"slug"`
	EstablishedDate string               `json:"established_date"`
	Title           string               `json:"title_brand"`
	Catchphrase     string               `json:"jargon_brand"`
	About           string               `json:"tentang_brand"`
	Status          string               `json:"status_brand"`
	BannerWeb       *view_models.FileVm  `json:"banner_web"`
	Logo            *view_models.FileVm  `json:"logo"`
	BannerMobile    *view_models.FileVm  `json:"banner_mobile"`
	Platform        []BrandMediaSocialVm `json:"platform"`
	Logs            []BrandLogVm         `json:"logs"`
}

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
	Slug            string               `json:"brand_slug"`
	Title           string               `json:"title_brand"`
	Catchphrase     string               `json:"jargon_brand"`
	Status          string               `json:"brand_status"`
	Logo            view_models.FileVm   `json:"logo"`
	WebBanner       view_models.FileVm   `json:"website_banner"`
	MobileBanner    view_models.FileVm   `json:"mobile_banner"`
	CreatedAt       string               `json:"created_at"`
	Owner           string               `json:"owner"`
	EstablishedDate string               `json:"established_date"`
	About           string               `json:"about"`
	Platform        []BrandMediaSocialVm `json:"platform"`
	Banned          []BrandLogVm         `json:"banned_logs"`
}

type BrandBannedVm struct {
	Status   string              `json:"status"`
	Reason   string              `json:"reason"`
	Document *view_models.FileVm `json:"supporting_documment"`
}

type BrandLogVm struct {
	Status   string             `json:"status"`
	Reason   string             `json:"reason"`
	Verifier string             `json:"verifier"`
	Document view_models.FileVm `json:"supporting_documment"`
}

type BrandVm struct {
	Full   []BrandFullVm `json:"list_full_brand"`
	List   []BrandListVm `json:"list_brand"`
	Detail BrandDetailVm `json:"detail_brand"`
	Logs   []BrandLogVm  `json:"log_brand"`
}

func NewBrandVm() BrandVm {
	return BrandVm{}
}

func (vm BrandVm) BuildFull(brands []models.Brand) (res []BrandFullVm) {
	for _, brand := range brands {
		var logo view_models.FileVm
		var web view_models.FileVm
		var mobile view_models.FileVm
		if brand.Logo != nil {
			logo = view_models.NewFileVm().Build(*brand.Logo)
		}
		if brand.BannerWeb != nil {
			web = view_models.NewFileVm().Build(*brand.BannerWeb)
		}
		if brand.BannerMobile != nil {
			mobile = view_models.NewFileVm().Build(*brand.BannerMobile)
		}
		res = append(res, BrandFullVm{
			ID:              brand.ID.String(),
			Name:            brand.Name,
			Slug:            brand.Slug,
			EstablishedDate: brand.EstablishedDate.Format("01-02-2006"),
			Title:           brand.Title,
			Catchphrase:     brand.Catchphrase,
			About:           brand.About,
			Status:          brand.Status,
			Logo:            &logo,
			BannerWeb:       &web,
			BannerMobile:    &mobile,
			Platform:        NewBrandMediaSocialVm().Build(brand.MediaSocials),
			Logs:            vm.BuildLog(brand.Logs),
		})
	}
	return res
}

func (vm BrandVm) BuildList(model []models.Brand) (res []BrandListVm) {
	for i, brand := range model {
		res = append(res, BrandListVm{
			No:     int64(i + 1),
			ID:     brand.ID.String(),
			Name:   brand.Name,
			Logo:   brand.LogoID.String(),
			Owner:  brand.Title,
			Status: brand.Status,
		})
	}
	return res
}

func (vm BrandVm) BuildDetail(brand *models.Brand) (res BrandDetailVm) {
	var logo view_models.FileVm
	var web view_models.FileVm
	var mobile view_models.FileVm
	if brand.Logo != nil {
		logo = view_models.NewFileVm().Build(*brand.Logo)
	}
	if brand.BannerWeb != nil {
		web = view_models.NewFileVm().Build(*brand.BannerWeb)
	}
	if brand.BannerMobile != nil {
		mobile = view_models.NewFileVm().Build(*brand.BannerMobile)
	}
	if brand != nil {
		res = BrandDetailVm{
			ID:              brand.ID.String(),
			Name:            brand.Name,
			Slug:            brand.Slug,
			Title:           brand.Title,
			Catchphrase:     brand.Catchphrase,
			Status:          brand.Status,
			Logo:            logo,
			WebBanner:       web,
			MobileBanner:    mobile,
			CreatedAt:       brand.CreatedAt.Format("01-02-2006"),
			Owner:           "",
			EstablishedDate: brand.EstablishedDate.Format("01-02-2006"),
			About:           brand.About,
			Platform:        NewBrandMediaSocialVm().Build(brand.MediaSocials),
			Banned:          vm.BuildLog(brand.Logs),
		}
	}
	return res
}

func (vm BrandVm) BuildLog(logs []models.BrandLog) (res []BrandLogVm) {
	for _, log := range logs {
		var attachment view_models.FileVm
		if log.Attachment != nil {
			attachment = view_models.NewFileVm().Build(*log.Attachment)
		}
		res = append(res, BrandLogVm{
			Status:   log.Status,
			Reason:   log.Reason,
			Verifier: *log.Verifier.Username,
			Document: attachment,
		})
	}
	return res
}

func (vm BrandVm) BuildBanned(brand *models.Brand) BrandBannedVm {
	// var reason string
	// var doc fileModel.File
	// if brand.Status == "Banned" {
	// 	reason = brand.BannedReason
	// 	doc = brand.BannedDocument
	// } else {
	// 	reason = brand.UnbannedReason
	// 	doc = brand.BannedDocument
	// }
	// return BrandBannedVm{
	// 	Status:   brand.Status,
	// 	Reason:   reason,
	// 	Document: view_models.NewFileVm().Build(doc),
	// }
	panic("Under development")
}
