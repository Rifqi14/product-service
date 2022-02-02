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

type MaterialCategoryUsecase struct {
	*usecase.Contract
}

func NewMaterialCategoryUsecase(contract *usecase.Contract) ucinterface.IMaterialCategoryUsecase {
	return &MaterialCategoryUsecase{Contract: contract}
}

func (uc MaterialCategoryUsecase) Create(req *request.MaterialCategoryRequest) (res *view_models.MaterialCategoryDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandMaterialCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.MaterialCategory{
		Name:      req.Name,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-materialCategory")
		return res, err
	}

	tx.Commit()
	res, _ = uc.Detail(material.ID)
	return res, nil
}

func (uc MaterialCategoryUsecase) List(req *request.Pagination) (res []view_models.MaterialCategoryDetailVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryMaterialCategoryRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	categories, count, err := repo.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-materialCategory")
		return res, pagination, err
	}

	for _, category := range categories {
		catVm := view_models.NewMaterialCategoryVm().BuildDetail(&category)
		res = append(res, *catVm)
	}

	pagination = uc.SetPaginationResponse(page, limit, count)
	return res, pagination, nil
}

func (uc MaterialCategoryUsecase) Detail(materialCatId uuid.UUID) (res *view_models.MaterialCategoryDetailVm, err error) {
	db := uc.DB
	repo := query.NewQueryMaterialCategoryRepository(db)

	categories, err := repo.Detail(materialCatId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-label")
		return res, err
	}

	if categories != nil {
		res = view_models.NewMaterialCategoryVm().BuildDetail(categories)
	}

	return res, nil
}

func (uc MaterialCategoryUsecase) Update(req *request.MaterialCategoryRequest, materialCatId uuid.UUID) (res *view_models.MaterialCategoryDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandMaterialCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.MaterialCategory{
		ID:        materialCatId,
		Name:      req.Name,
		UpdatedBy: &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	materialCategory, err := repo.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-materialCategory")
		return res, err
	}

	tx.Commit()
	res, _ = uc.Detail(materialCategory.ID)
	return res, nil
}

func (uc MaterialCategoryUsecase) Delete(materialCatId uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandMaterialCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.MaterialCategory{
		ID:        materialCatId,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-materialCategory")
		return err
	}

	return nil
}

func (uc MaterialCategoryUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
