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

type GenderUsecase struct {
	*usecase.Contract
}

func NewGenderUsecase(contract *usecase.Contract) ucinterface.IGenderUsecase {
	return &GenderUsecase{Contract: contract}
}

func (uc GenderUsecase) Create(req *request.GenderRequest) (res view_models.GenderDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandGenderRepository(db)

	genderId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Gender{
		Name:      req.Name,
		ParentID:  req.ParentID,
		CreatedBy: &genderId,
		UpdatedBy: &genderId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	gender, err := repo.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-color")
		return res, err
	}

	res = view_models.NewGenderVm().BuildDetail(&gender)
	tx.Commit()
	return res, nil
}

func (uc GenderUsecase) List(req *request.Pagination) (res []view_models.GenderListVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryGenderRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	genders, count, err := repo.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-gender")
		return res, pagination, err
	}

	res = view_models.NewGenderVm().BuildList(genders)

	pagination = uc.SetPaginationResponse(page, limit, count)
	return res, pagination, nil
}

func (uc GenderUsecase) Detail(genderId uuid.UUID) (res view_models.GenderDetailVm, err error) {
	db := uc.DB
	repo := query.NewQueryGenderRepository(db)

	gender, err := repo.Detail(genderId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-gender")
		return res, err
	}

	res = view_models.NewGenderVm().BuildDetail(&gender)
	return res, nil
}

func (uc GenderUsecase) Update(req *request.GenderRequest, genderId uuid.UUID) (res view_models.GenderDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandGenderRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Gender{
		ID:        genderId,
		Name:      req.Name,
		ParentID:  req.ParentID,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	gender, err := repo.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-gender")
		return res, err
	}

	res = view_models.NewGenderVm().BuildDetail(&gender)
	tx.Commit()
	return res, nil
}

func (uc GenderUsecase) Delete(genderId uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandGenderRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Gender{
		ID:        genderId,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-gender")
		return err
	}

	return nil
}

func (uc GenderUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
