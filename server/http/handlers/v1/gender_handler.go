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

type GenderHandler struct {
	handlers.Handler
}

func NewGenderHandler(handler handlers.Handler) hinterface.IGenderHandler {
	return &GenderHandler{Handler: handler}
}

func (handler GenderHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.GenderRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewGenderUsecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler GenderHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.Pagination)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewColorUsecase(handler.UcContract)
	data, meta, err := uc.List(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "data gender success fetched")).Send(ctx)
}

func (handler GenderHandler) Detail(ctx *fiber.Ctx) (err error) {
	genderId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.UserNotFound, err)).Send(ctx)
	}

	uc := v1.NewGenderUsecase(handler.UcContract)
	data, err := uc.Detail(genderId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, nil, "gender success fetched")).Send(ctx)
}

func (handler GenderHandler) Update(ctx *fiber.Ctx) (err error) {
	genderId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.GenderRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewGenderUsecase(handler.UcContract)
	res, err := uc.Update(req, genderId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.EditFailed, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler GenderHandler) Delete(ctx *fiber.Ctx) (err error) {
	genderId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewGenderUsecase(handler.UcContract)
	err = uc.Delete(genderId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "Success delete gender")).Send(ctx)
}

func (handler GenderHandler) Export(ctx *fiber.Ctx) (err error) {
	panic("Under development")
}
