package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IBrandUsecase interface {
	Create(req *request.BrandRequest) (res view_models.BrandDetailVm, err error)

	List(req *request.Pagination) (res []view_models.BrandListVm, pagination view_models.PaginationVm, err error)

	Detail(brandId uuid.UUID) (res view_models.BrandDetailVm, err error)

	Update(req *request.BrandRequest, brandID uuid.UUID) (res view_models.BrandDetailVm, err error)

	Delete(ID uuid.UUID) (err error)

	Export(fileType string) (err error)

	Banned(req *request.BannedBrandRequest, brandID uuid.UUID) (res view_models.BrandDetailVm, err error)
}
