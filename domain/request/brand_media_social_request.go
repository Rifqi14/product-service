package request

type BrandMediaSocialRequest struct {
	Type string `form:"type" json:"type"`
	Link string `form:"link" json:"link"`
}
