package view_models

type PaginationVm struct {
	Pagination DetailPaginationVm `json:"pagination"`
}

type DetailPaginationVm struct {
	CurrentPage int64 `json:"current_page"`
	LastPage    int64 `json:"last_page"`
	Total       int64 `json:"total"`
	PerPage     int64 `json:"per_page"`
}

func NewPaginationVm() PaginationVm {
	return PaginationVm{}
}

func (vm PaginationVm) Build(detailPagination DetailPaginationVm) PaginationVm {
	return PaginationVm{Pagination: detailPagination}
}
