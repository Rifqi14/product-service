package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IMaterialCategoryUsecase interface {
	Create(req *request.MaterialCategoryRequest) (res *view_models.MaterialCategoryDetailVm, err error)

	List(req *request.Pagination) (res []view_models.MaterialCategoryDetailVm, pagination view_models.PaginationVm, err error)

	Detail(materialCatId uuid.UUID) (res *view_models.MaterialCategoryDetailVm, err error)

	Update(req *request.MaterialCategoryRequest, materialCatId uuid.UUID) (res *view_models.MaterialCategoryDetailVm, err error)

	Delete(materialCatId uuid.UUID) (err error)

	Export(fileType string) (err error)
}
