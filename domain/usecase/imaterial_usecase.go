package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IMaterialUsecase interface {
	Create(req *request.MaterialRequest) (res view_models.MaterialDetailVm, err error)

	List(req *request.Pagination) (res []view_models.MaterialListVm, pagination view_models.PaginationVm, err error)

	Detail(materialId uuid.UUID) (res view_models.MaterialDetailVm, err error)

	Update(req *request.MaterialRequest, materialId uuid.UUID) (res view_models.MaterialDetailVm, err error)

	Delete(materialId uuid.UUID) (err error)

	Export(fileType string) (err error)
}
