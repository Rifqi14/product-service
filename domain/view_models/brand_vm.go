package view_models

import (
	"strings"

	"gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type BrandFullVm struct {
	ID              string                `json:"brand_id"`
	Name            string                `json:"name"`
	Slug            string                `json:"slug"`
	EstablishedDate string                `json:"established_date"`
	Title           string                `json:"title_brand"`
	Catchphrase     string                `json:"jargon_brand"`
	About           string                `json:"tentang_brand"`
	Status          string                `json:"status_brand"`
	BannerWeb       *view_models.FileVm   `json:"banner_web"`
	Logo            *view_models.FileVm   `json:"logo"`
	BannerMobile    *view_models.FileVm   `json:"banner_mobile"`
	Platform        []*BrandMediaSocialVm `json:"platform"`
	Logs            []*BrandLogVm         `json:"logs"`
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
	Banned          []*BrandLogVm        `json:"banned_logs"`
}

type BrandBannedVm struct {
	Status   string              `json:"status"`
	Reason   string              `json:"reason"`
	Document *view_models.FileVm `json:"supporting_documment"`
}

type BrandLogVm struct {
	Status   string              `json:"status"`
	Reason   string              `json:"reason"`
	Verifier string              `json:"verifier"`
	Document *view_models.FileVm `json:"supporting_documment"`
}

type BrandExportVm struct {
	Number          int    `json:"number"`
	Name            string `json:"name"`
	EstablishedDate string `json:"established_date"`
	Title           string `json:"title"`
	Catchphrase     string `json:"catchphrase"`
	About           string `json:"about"`
	Website         string `json:"website"`
	Instagram       string `json:"instagram"`
	Tiktok          string `json:"tiktok"`
	Facebook        string `json:"facebook"`
	Twitter         string `json:"twitter"`
	Email           string `json:"email"`
	Other           string `json:"other"`
}

type BrandVm struct {
	Full   []BrandFullVm   `json:"list_full_brand"`
	List   []BrandListVm   `json:"list_brand"`
	Detail BrandDetailVm   `json:"detail_brand"`
	Logs   []BrandLogVm    `json:"log_brand"`
	Export []BrandExportVm `json:"export_brand"`
}

func NewBrandVm() BrandVm {
	return BrandVm{}
}

func (vm BrandVm) BuildExport(brands []models.Brand) (res []BrandExportVm, err error) {
	if len(brands) > 0 {
		for i, brand := range brands {
			var export BrandExportVm
			var website []string
			var instagram []string
			var tiktok []string
			var facebook []string
			var twitter []string
			var email []string
			var other []string
			if len(brand.MediaSocials) > 0 {
				for _, platform := range brand.MediaSocials {
					switch platform.Type {
					case "website":
						website = append(website, platform.Link)
					case "instagram":
						instagram = append(instagram, platform.Link)
					case "tiktok":
						tiktok = append(tiktok, platform.Link)
					case "facebook":
						facebook = append(facebook, platform.Link)
					case "twitter":
						twitter = append(twitter, platform.Link)
					case "email":
						email = append(email, platform.Link)
					default:
						other = append(other, platform.Link)
					}
				}
			}
			export.Number = i + 1
			export.Name = brand.Name
			export.EstablishedDate = brand.EstablishedDate.Format("02/01/2006")
			export.Title = brand.Title
			export.Catchphrase = brand.Catchphrase
			export.About = brand.About
			export.Website = strings.Join(website, ", ")
			export.Instagram = strings.Join(instagram, ", ")
			export.Tiktok = strings.Join(tiktok, ", ")
			export.Facebook = strings.Join(facebook, ", ")
			export.Twitter = strings.Join(twitter, ", ")
			export.Email = strings.Join(email, ", ")
			export.Other = strings.Join(other, ", ")
			res = append(res, export)
		}
	}
	return res, nil
}

func (vm BrandVm) BuildFull(brands []models.Brand) (res []BrandFullVm) {
	if len(brands) > 0 {
		for _, brand := range brands {
			if brand.Name != "" {
				var logo *view_models.FileVm
				var web *view_models.FileVm
				var mobile *view_models.FileVm
				var medsos []*BrandMediaSocialVm
				if brand.Logo != nil {
					logoVm := view_models.NewFileVm().Build(*brand.Logo)
					logo = &logoVm
				}
				if brand.BannerWeb != nil {
					webVm := view_models.NewFileVm().Build(*brand.BannerWeb)
					web = &webVm
				}
				if brand.BannerMobile != nil {
					mobileVm := view_models.NewFileVm().Build(*brand.BannerMobile)
					mobile = &mobileVm
				}
				if len(brand.MediaSocials) > 0 {
					medsosVm := NewBrandMediaSocialVm().Build(brand.MediaSocials)
					for _, vm := range medsosVm {
						medsos = append(medsos, &vm)
					}
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
					Logo:            logo,
					BannerWeb:       web,
					BannerMobile:    mobile,
					Platform:        medsos,
					Logs:            vm.BuildLog(brand.Logs),
				})
			}
		}
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

func (vm BrandVm) BuildLog(logs []*models.BrandLog) (res []*BrandLogVm) {
	if len(logs) > 0 {
		for _, log := range logs {
			var attachment *view_models.FileVm
			if log.Attachment != nil {
				attachmentVm := view_models.NewFileVm().Build(*log.Attachment)
				attachment = &attachmentVm
			}
			res = append(res, &BrandLogVm{
				Status:   log.Status,
				Reason:   log.Reason,
				Verifier: *log.Verifier.Username,
				Document: attachment,
			})
		}
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
