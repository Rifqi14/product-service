package request

import (
	"github.com/google/uuid"
)

type BrandRequest struct {
	Name            string                    `form:"name" json:"name" validate:"required"`
	LogoID          *uuid.UUID                `form:"logo_id" json:"logo_id"`
	BannerWebID     *uuid.UUID                `form:"banner_web_id" json:"banner_web_id"`
	BannerMobileID  *uuid.UUID                `form:"banner_mobile_id" json:"banner_mobile_id"`
	EstablishedDate string                    `form:"established_date" json:"established_date"`
	Title           string                    `form:"title" json:"title"`
	Catchphrase     string                    `form:"catchphrase" json:"catchphrase"`
	About           string                    `form:"about" json:"about"`
	MediaSocial     []BrandMediaSocialRequest `form:"media_social" json:"media_social"`
}
