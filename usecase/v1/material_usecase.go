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

type MaterialUsecase struct {
	*usecase.Contract
}

func NewMaterialusecase(contract *usecase.Contract) ucinterface.IMaterialUsecase {
	return &MaterialUsecase{Contract: contract}
}

func (uc MaterialUsecase) Create(req *request.MaterialRequest) (res view_models.MaterialDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandMaterialRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Material{
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
	material, err := repo.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-material")
		return res, err
	}

	res = view_models.NewMaterialVm().BuildDetail(&material)
	tx.Commit()
	return res, nil
}

func (uc MaterialUsecase) List(req *request.Pagination) (res []view_models.MaterialListVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryMaterialRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	materials, count, err := repo.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-material")
		return res, pagination, err
	}

	res = view_models.NewMaterialVm().BuildList(materials)

	pagination = uc.SetPaginationResponse(page, limit, count)
	return res, pagination, nil
}

func (uc MaterialUsecase) Detail(materialId uuid.UUID) (res view_models.MaterialDetailVm, err error) {
	db := uc.DB
	repo := query.NewQueryMaterialRepository(db)

	gender, err := repo.Detail(materialId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-material")
		return res, err
	}

	res = view_models.NewMaterialVm().BuildDetail(&gender)
	return res, nil
}

func (uc MaterialUsecase) Update(req *request.MaterialRequest, materialId uuid.UUID) (res view_models.MaterialDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandMaterialRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Material{
		ID:        materialId,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-material")
		return res, err
	}

	res = view_models.NewMaterialVm().BuildDetail(&label)
	tx.Commit()
	return res, nil
}

func (uc MaterialUsecase) Delete(materialId uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandMaterialRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Material{
		ID:        materialId,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-material")
		return err
	}

	return nil
}

func (uc MaterialUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
