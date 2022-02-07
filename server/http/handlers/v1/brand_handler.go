package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-package-svc/responses"
	handlers2 "gitlab.com/s2.1-backend/shm-product-svc/domain/handlers"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/usecase/v1"
)

type BrandHandler struct {
	handlers.Handler
}

func NewBrandHandler(handler handlers.Handler) handlers2.IBrandHandler {
	return &BrandHandler{Handler: handler}
}

func (handler BrandHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.BrandRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "Validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewBrandUsecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStored, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler BrandHandler) Update(ctx *fiber.Ctx) (err error) {
	brandID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.BrandRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "Validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewBrandUsecase(handler.UcContract)
	res, err := uc.Update(req, brandID)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler BrandHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.Pagination)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	// Service processing
	uc := v1.NewBrandUsecase(handler.UcContract)
	data, meta, err := uc.List(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, "error get data", err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "data module success fetched")).Send(ctx)
}

func (handler BrandHandler) Detail(ctx *fiber.Ctx) (err error) {
	brandID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewBrandUsecase(handler.UcContract)
	data, err := uc.Detail(brandID)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, "Data not found", err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, nil, "Brand success fetched")).Send(ctx)
}

func (handler BrandHandler) Delete(ctx *fiber.Ctx) (err error) {
	brandID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewBrandUsecase(handler.UcContract)
	err = uc.Delete(brandID)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "success delete brand")).Send(ctx)
}

func (handler BrandHandler) Export(ctx *fiber.Ctx) (err error) {
	panic("Under Maintenance")
}

func (handler BrandHandler) Banned(ctx *fiber.Ctx) (err error) {
	brandId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	req := new(request.BannedBrandRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "error validate payload", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewBrandUsecase(handler.UcContract)
	res, err := uc.Banned(req, brandId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}
