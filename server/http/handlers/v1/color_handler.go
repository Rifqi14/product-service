package v1

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-package-svc/responses"
	hinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/handlers"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/usecase/v1"
)

type ColorHandler struct {
	handlers.Handler
}

func NewColorHandler(handler handlers.Handler) hinterface.IColorHandler {
	return &ColorHandler{Handler: handler}
}

func (handler ColorHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.ColorRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "Validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler ColorHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.Pagination)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	data, meta, err := uc.List(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, "error get data", err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "Data color success fetched")).Send(ctx)
}

func (handler ColorHandler) Detail(ctx *fiber.Ctx) (err error) {
	colorId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	data, err := uc.Detail(colorId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, "Data not found", err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, nil, "Color success fetched")).Send(ctx)
}

func (handler ColorHandler) Update(ctx *fiber.Ctx) (err error) {
	colorId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.ColorRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "Error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "Validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	res, err := uc.Update(req, colorId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler ColorHandler) Delete(ctx *fiber.Ctx) (err error) {
	colorId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	err = uc.Delete(colorId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "Success delete color")).Send(ctx)
}

func (handler ColorHandler) Export(ctx *fiber.Ctx) (err error) {
	fileType := ctx.Query("type")
	if !handlers.FileType(fileType).IsValid() {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	res, err := uc.Export(fileType)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, "data success export")).Send(ctx)
}
