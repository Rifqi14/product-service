package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-package-svc/responses"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/usecase/v1"

	hinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/handlers"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
)

type ProductHandler struct {
	handlers.Handler
}

func NewProductHandler(handler handlers.Handler) hinterface.IProductHandler {
	return &ProductHandler{Handler: handler}
}

func (handler ProductHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.ProductRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}
	for _, variant := range req.Variants {
		image := handlers.ValidateImage(variant)
		if len(image) <= 0 {
			return responses.NewResponse(responses.ResponseError(nil, nil, fiber.StatusBadRequest, "each variant at least need one image", errors.New("each variant at least need one image"))).Send(ctx)
		}
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler ProductHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.FilterQueryProductRequest)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	reqFilter := new(request.FilterProductRequest)
	reqFilter.Pagination = req.Pagination
	reqFilter.ProductName = req.ProductName
	reqFilter.MinPrice = req.MinPrice
	reqFilter.MaxPrice = req.MaxPrice

	if len(req.Product) > 0 && req.Product[0] != "" {
		for _, productID := range req.Product {
			ID := uuid.MustParse(productID)
			reqFilter.Product = append(reqFilter.Product, &ID)
		}
	}
	if len(req.Brand) > 0 && req.Brand[0] != "" {
		for _, brandID := range req.Brand {
			ID := uuid.MustParse(brandID)
			reqFilter.Brand = append(reqFilter.Brand, &ID)
		}
	}
	if len(req.Color) > 0 && req.Color[0] != "" {
		for _, colorID := range req.Color {
			ID := uuid.MustParse(colorID)
			reqFilter.Color = append(reqFilter.Color, &ID)
		}
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	data, meta, err := uc.List(reqFilter)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "data product success fetched")).Send(ctx)
}

func (handler ProductHandler) Detail(ctx *fiber.Ctx) (err error) {
	productId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		fmt.Println(err)
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.UserNotFound, err)).Send(ctx)
	}

	// isPublic := strings.Contains(ctx.Path(), "/product/v1/product/public")

	// if isPublic {
	// 	return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	// }

	uc := v1.NewProductUsecase(handler.UcContract)
	data, err := uc.Detail(productId)
	if err != nil {
		fmt.Println(err)
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, nil, "product success fetched")).Send(ctx)
}

func (handler ProductHandler) Update(ctx *fiber.Ctx) (err error) {
	productId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.ProductRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	res, err := uc.Update(req, productId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.EditFailed, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler ProductHandler) Delete(ctx *fiber.Ctx) (err error) {
	productId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	err = uc.Delete(productId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "Success delete product")).Send(ctx)
}

func (handler ProductHandler) Export(ctx *fiber.Ctx) (err error) {
	fileType := ctx.Query("type")
	if !handlers.FileType(fileType).IsValid() {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	res, err := uc.Export(fileType)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, "data suuccess export")).Send(ctx)
}

func (handler ProductHandler) ChangeStatus(ctx *fiber.Ctx) (err error) {
	productId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	req := new(request.BannedProductRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewProductUsecase(handler.UcContract)
	err = uc.ChangeStatus(req, &productId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.EditFailed, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "product berhasil dinonaktifkan")).Send(ctx)
}
