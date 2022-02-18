package usecase

import (
	"github.com/google/uuid"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IGenderUsecase interface {
	Create(req *request.GenderRequest) (res view_models.GenderDetailVm, err error)

	List(req *request.Pagination) (res []view_models.GenderListVm, pagination view_models.PaginationVm, err error)

	Detail(genderId uuid.UUID) (res view_models.GenderDetailVm, err error)

	Update(req *request.GenderRequest, genderId uuid.UUID) (res view_models.GenderDetailVm, err error)

	Delete(genderId uuid.UUID) (err error)

	Export(fileType string) (link *fileVm.FileVm, err error)
}
