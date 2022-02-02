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

type LabelUsecase struct {
	*usecase.Contract
}

func NewLabelUsecase(contract *usecase.Contract) ucinterface.ILabelUsecase {
	return &LabelUsecase{Contract: contract}
}

func (uc LabelUsecase) Create(req *request.LabelRequest) (res view_models.LabelDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandLabelRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Label{
		Name:      req.Name,
		ParentID:  req.ParentID,
		CreatedBy: &userId,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	label, err := repo.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-label")
		return res, err
	}

	res = view_models.NewLabelVm().BuildDetail(&label)
	tx.Commit()
	return res, nil
}

func (uc LabelUsecase) List(req *request.Pagination) (res []view_models.LabelListVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryLabelRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	labels, count, err := repo.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-label")
		return res, pagination, err
	}

	res = view_models.NewLabelVm().BuildList(labels)

	pagination = uc.SetPaginationResponse(page, limit, count)
	return res, pagination, nil
}

func (uc LabelUsecase) Detail(labelId uuid.UUID) (res view_models.LabelDetailVm, err error) {
	db := uc.DB
	repo := query.NewQueryLabelRepository(db)

	gender, err := repo.Detail(labelId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-label")
		return res, err
	}

	res = view_models.NewLabelVm().BuildDetail(&gender)
	return res, nil
}

func (uc LabelUsecase) Update(req *request.LabelRequest, labelId uuid.UUID) (res view_models.LabelDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandLabelRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Label{
		ID:        labelId,
		Name:      req.Name,
		ParentID:  req.ParentID,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	label, err := repo.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-label")
		return res, err
	}

	res = view_models.NewLabelVm().BuildDetail(&label)
	tx.Commit()
	return res, nil
}

func (uc LabelUsecase) Delete(labelId uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandLabelRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Label{
		ID:        labelId,
		DeletedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return err
	}
	err = repo.Delete(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-label")
		return err
	}

	return nil
}

func (uc LabelUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
