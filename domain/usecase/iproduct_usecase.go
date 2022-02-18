package usecase

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type IProductUsecase interface {
	Create(req *request.ProductRequest) (res *view_models.ProductVm, err error)

	FindBy(req *request.FindByRequest) (res []*view_models.ProductVm, pagination view_models.PaginationVm, err error)

	List(req *request.FilterProductRequest) (res []*view_models.ProductVm, pagination view_models.PaginationVm, err error)

	Detail(productId uuid.UUID) (res *view_models.ProductVm, err error)

	Update(req *request.ProductRequest, productId uuid.UUID) (res *view_models.ProductVm, err error)

	Delete(productId uuid.UUID) (err error)

	Export(fileType string) (err error)

	ChangeStatus(req *request.BannedProductRequest, productId *uuid.UUID) (err error)
}
