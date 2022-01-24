package v1

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	ucinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/usecase"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/command"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/query"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
)

type ColorUsecase struct {
	*usecase.Contract
}

func NewColorUsecase(contract *usecase.Contract) ucinterface.IColorUsecase {
	return &ColorUsecase{Contract: contract}
}

func (uc ColorUsecase) Create(req *request.ColorRequest) (res view_models.ColorDetailVm, err error) {
	db := uc.DB
	repository := command.NewCommandColorRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Color{
		Name:      req.Name,
		RgbCode:   req.RgbCode,
		ParentID:  req.ParentID,
		CreatedBy: &userId,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	color, err := repository.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-color")
		return res, err
	}

	res = view_models.NewColorVm().BuildDetail(&color)
	tx.Commit()
	return res, nil
}

func (uc ColorUsecase) List(req *request.Pagination) (res []view_models.ColorListVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repository := query.NewQueryColorRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	colors, count, err := repository.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-color")
		return res, pagination, err
	}

	res = view_models.NewColorVm().BuildList(colors)

	pagination = uc.SetPaginationResponse(page, limit, count)

	return res, pagination, nil
}

func (uc ColorUsecase) Detail(colorID uuid.UUID) (res view_models.ColorDetailVm, err error) {
	db := uc.DB
	repository := query.NewQueryColorRepository(db)

	color, err := repository.Detail(colorID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-color")
		return res, err
	}

	res = view_models.NewColorVm().BuildDetail(&color)
	return res, nil
}

func (uc ColorUsecase) Update(req *request.ColorRequest, colorID uuid.UUID) (res view_models.ColorDetailVm, err error) {
	db := uc.DB
	repository := command.NewCommandColorRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Color{
		ID:        colorID,
		Name:      req.Name,
		ParentID:  req.ParentID,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	color, err := repository.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-color")
		return res, err
	}

	res = view_models.NewColorVm().BuildDetail(&color)
	tx.Commit()
	return res, nil
}

func (uc ColorUsecase) Delete(colorID uuid.UUID) (err error) {
	db := uc.DB
	repository := command.NewCommandColorRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Color{
		ID:        colorID,
		DeletedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return err
	}
	err = repository.Delete(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-color")
		return err
	}

	return nil
}

func (uc ColorUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
