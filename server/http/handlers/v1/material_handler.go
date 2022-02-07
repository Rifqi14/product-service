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

type MaterialHandler struct {
	handlers.Handler
}

func NewMaterialHandler(handler handlers.Handler) hinterface.IMaterialHandler {
	return &MaterialHandler{Handler: handler}
}

func (handler MaterialHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.MaterialRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewMaterialusecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler MaterialHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.Pagination)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewMaterialusecase(handler.UcContract)
	data, meta, err := uc.List(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "data material success fetched")).Send(ctx)
}

func (handler MaterialHandler) Detail(ctx *fiber.Ctx) (err error) {
	materialId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.UserNotFound, err)).Send(ctx)
	}

	uc := v1.NewMaterialusecase(handler.UcContract)
	data, err := uc.Detail(materialId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, nil, "material success fetched")).Send(ctx)
}

func (handler MaterialHandler) Update(ctx *fiber.Ctx) (err error) {
	material, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.MaterialRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewMaterialusecase(handler.UcContract)
	res, err := uc.Update(req, material)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.EditFailed, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler MaterialHandler) Delete(ctx *fiber.Ctx) (err error) {
	materialId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewMaterialusecase(handler.UcContract)
	err = uc.Delete(materialId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "Success delete material")).Send(ctx)
}

func (handler MaterialHandler) Export(ctx *fiber.Ctx) (err error) {
	panic("Under development")
}
