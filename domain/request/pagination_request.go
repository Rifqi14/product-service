package request

type Pagination struct {
	Search  string `form:"search" json:"search"`
	OrderBy string `form:"order_by" json:"order_by"`
	Sort    string `form:"sort" json:"sort"`
	Limit   int64  `form:"limit" json:"limit"`
	Offset  int64  `form:"offset" json:"offset"`
}

type Read struct {
	Column   string `form:"column" json:"column"`
	Operator string `form:"operator" json:"operator"`
	Value    string `form:"value" json:"value"`
}
