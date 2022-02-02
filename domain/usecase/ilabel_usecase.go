package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type ILabelUsecase interface {
	Create(req *request.LabelRequest) (res view_models.LabelDetailVm, err error)

	List(req *request.Pagination) (res []view_models.LabelListVm, pagination view_models.PaginationVm, err error)

	Detail(labelId uuid.UUID) (res view_models.LabelDetailVm, err error)

	Update(req *request.LabelRequest, labelId uuid.UUID) (res view_models.LabelDetailVm, err error)

	Delete(labelId uuid.UUID) (err error)

	Export(fileType string) (err error)
}
