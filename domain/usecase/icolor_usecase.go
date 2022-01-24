package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IColorUsecase interface {
	Create(req *request.ColorRequest) (res view_models.ColorDetailVm, err error)

	List(req *request.Pagination) (res []view_models.ColorListVm, pagination view_models.PaginationVm, err error)

	Detail(colorId uuid.UUID) (res view_models.ColorDetailVm, err error)

	Update(req *request.ColorRequest, colorId uuid.UUID) (res view_models.ColorDetailVm, err error)

	Delete(colorId uuid.UUID) (err error)

	Export(fileType string) (err error)
}
