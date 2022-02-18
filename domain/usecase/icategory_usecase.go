package usecase

import (
	"github.com/google/uuid"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type ICategoryUsecase interface {
	Create(req *request.CategoryRequest) (res view_models.CategoryDetailVm, err error)

	List(req *request.Pagination) (res []view_models.CategoryListVm, pagination view_models.PaginationVm, err error)

	Detail(categoryId uuid.UUID) (res view_models.CategoryDetailVm, err error)

	Update(req *request.CategoryRequest, categoryId uuid.UUID) (res view_models.CategoryDetailVm, err error)

	Delete(categoryID uuid.UUID) (err error)

	Export(fileType string) (link *fileVm.FileVm, err error)
}
