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

type LabelHandler struct {
	handlers.Handler
}

func NewLabelHandler(handler handlers.Handler) hinterface.ILabelHandler {
	return &LabelHandler{Handler: handler}
}

func (handler LabelHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(request.LabelRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	res, err := uc.Create(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataStoredError, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler LabelHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(request.Pagination)
	if err := ctx.QueryParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}
	if err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validasi data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	data, meta, err := uc.List(req)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(data, meta, "data label success fetched")).Send(ctx)
}

func (handler LabelHandler) Detail(ctx *fiber.Ctx) (err error) {
	labelId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.UserNotFound, err)).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	data, err := uc.Detail(labelId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.DataNotFound, err)).Send(ctx)
	}

	var res interface{}
	if data.Name != "" {
		res = data
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, "label success fetched")).Send(ctx)
}

func (handler LabelHandler) Update(ctx *fiber.Ctx) (err error) {
	label, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing id to uuid", err)).Send(ctx)
	}
	req := new(request.LabelRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, "error parsing body request", err)).Send(ctx)
	}
	if err := handler.Validate.Struct(req); err != nil {
		return responses.NewResponse(responses.ResponseErrorValidation(nil, nil, http.StatusBadRequest, "validation data error", err.(validator.ValidationErrors))).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	res, err := uc.Update(req, label)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.EditFailed, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, messages.DataStored)).Send(ctx)
}

func (handler LabelHandler) Delete(ctx *fiber.Ctx) (err error) {
	labelId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	err = uc.Delete(labelId)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnprocessableEntity, messages.FailedLoadPayload, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(nil, nil, "Success delete label")).Send(ctx)
}

func (handler LabelHandler) Export(ctx *fiber.Ctx) (err error) {
	fileType := ctx.Query("type")
	if !handlers.FileType(fileType).IsValid() {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.FailedLoadPayload, err)).Send(ctx)
	}

	uc := v1.NewLabelUsecase(handler.UcContract)
	res, err := uc.Export(fileType)
	if err != nil {
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusBadRequest, messages.DataNotFound, err)).Send(ctx)
	}

	return responses.NewResponse(responses.ResponseSuccess(res, nil, "data suuccess export")).Send(ctx)
}
