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

type CategoryUsecase struct {
	*usecase.Contract
}

func NewCategoryUsecase(contract *usecase.Contract) ucinterface.ICategoryUsecase {
	return &CategoryUsecase{Contract: contract}
}

func (uc CategoryUsecase) Create(req *request.CategoryRequest) (res view_models.CategoryDetailVm, err error) {
	db := uc.DB
	repository := command.NewCommandCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}
	model := models.Category{
		Name:                req.Name,
		ParentID:            req.ParentID,
		MobileBannerID:      req.MobileBannerID,
		WebsiteBannerID:     req.WebsiteBannerID,
		MobileHeroBannerID:  req.MobileHeroBannerID,
		WebsiteHeroBannerID: req.WebsiteHeroBannerID,
		CreatedBy:           &userId,
		UpdatedBy:           &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	category, err := repository.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-category")
		return res, err
	}

	res = view_models.NewCategoryVm().BuildDetail(&category)
	tx.Commit()
	return res, nil
}

func (uc CategoryUsecase) List(req *request.Pagination) (res []view_models.CategoryListVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repository := query.NewQueryCategoryRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	categories, count, err := repository.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-category")
		return res, pagination, err
	}

	res = view_models.NewCategoryVm().BuildList(categories)

	pagination = uc.SetPaginationResponse(page, limit, count)

	return res, pagination, nil
}

func (uc CategoryUsecase) Detail(categoryID uuid.UUID) (res view_models.CategoryDetailVm, err error) {
	db := uc.DB
	repository := query.NewQueryCategoryRepository(db)

	category, err := repository.Detail(categoryID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-brand")
		return res, err
	}

	res = view_models.NewCategoryVm().BuildDetail(&category)
	return res, nil
}

func (uc CategoryUsecase) Update(req *request.CategoryRequest, categoryID uuid.UUID) (res view_models.CategoryDetailVm, err error) {
	db := uc.DB
	repository := command.NewCommandCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Category{
		ID:                  categoryID,
		Name:                req.Name,
		ParentID:            req.ParentID,
		MobileBannerID:      req.MobileBannerID,
		WebsiteBannerID:     req.WebsiteBannerID,
		MobileHeroBannerID:  req.MobileHeroBannerID,
		WebsiteHeroBannerID: req.WebsiteHeroBannerID,
		UpdatedBy:           &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	category, err := repository.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-category")
		return res, err
	}

	res = view_models.NewCategoryVm().BuildDetail(&category)
	tx.Commit()
	return res, nil
}

func (uc CategoryUsecase) Delete(categoryID uuid.UUID) (err error) {
	db := uc.DB
	repository := command.NewCommandCategoryRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Category{
		ID:        categoryID,
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
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-category")
		return err
	}

	return nil
}

func (uc CategoryUsecase) Export(fileType string) (err error) {
	panic("Under development")
}
