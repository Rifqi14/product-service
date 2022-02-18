package request

type Pagination struct {
	Search  string `form:"search" json:"search" query:"search"`
	OrderBy string `form:"order_by" json:"order_by" query:"order_by"`
	Sort    string `form:"sort" json:"sort" query:"sort"`
	Limit   int64  `form:"limit" json:"limit" query:"limit"`
	Offset  int64  `form:"offset" json:"offset" query:"offset"`
}

type Read struct {
	Column   string `form:"column" json:"column"`
	Operator string `form:"operator" json:"operator"`
	Value    string `form:"value" json:"value"`
}
