package request

import "github.com/google/uuid"

type CategoryRequest struct {
	Name                string     `form:"name" json:"name" validate:"required"`
	ParentID            *uuid.UUID `form:"parent_id" json:"parent_id"`
	MobileBannerID      *uuid.UUID `form:"mobile_banner_id" json:"mobile_banner_id"`
	WebsiteBannerID     *uuid.UUID `form:"website_banner_id" json:"website_banner_id"`
	MobileHeroBannerID  *uuid.UUID `form:"mobile_hero_banner_id" json:"mobile_hero_banner_id"`
	WebsiteHeroBannerID *uuid.UUID `form:"website_hero_banner_id" json:"website_hero_banner_id"`
}
