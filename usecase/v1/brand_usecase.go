package v1

import (
	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/datetime"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-package-svc/str"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	ucinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/usecase"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/command"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/query"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
)

type BrandUsecase struct {
	*usecase.Contract
}

func NewBrandUsecase(contract *usecase.Contract) ucinterface.IBrandUsecase {
	return &BrandUsecase{Contract: contract}
}

func (uc BrandUsecase) Create(req *request.BrandRequest) (res view_models.BrandDetailVm, err error) {
	repository := command.NewCommandBrandRepository(uc.DB)

	userID, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}
	var brandSocmend []models.BrandMediaSocial
	for _, socmed := range req.MediaSocial {
		brandSocmend = append(brandSocmend, models.BrandMediaSocial{
			Type: socmed.Type,
			Link: socmed.Link,
		})
	}
	model := models.Brand{
		Name:            req.Name,
		Slug:            str.Slug(req.Name),
		EstablishedDate: datetime.StrParseToTime(req.EstablishedDate, "2006-10-02"),
		Title:           req.Title,
		Catchphrase:     req.Catchphrase,
		About:           req.About,
		LogoID:          req.LogoID,
		BannerWebID:     req.BannerWebID,
		BannerMobileID:  req.BannerMobileID,
		CreatedBy:       &userID,
		UpdatedBy:       &userID,
		MediaSocials:    brandSocmend,
	}
	tx := uc.DB.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	brand, err := repository.Create(model, tx)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-brand")
		return res, err
	}

	res = view_models.NewBrandVm().BuildDetail(&brand)
	tx.Commit()
	return res, nil
}

func (uc BrandUsecase) List(req *request.Pagination) (res []view_models.BrandListVm, pagination view_models.PaginationVm, err error) {
	repository := query.NewQueryBrandRepository(uc.DB)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	brands, count, err := repository.List(req.Search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-brand")
		return res, pagination, err
	}

	res = view_models.NewBrandVm().BuildList(brands)

	pagination = uc.SetPaginationResponse(page, limit, count)

	return res, pagination, nil
}

func (uc BrandUsecase) Detail(brandID uuid.UUID) (res view_models.BrandDetailVm, err error) {
	repository := query.NewQueryBrandRepository(uc.DB)

	brand, err := repository.Detail(brandID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-brand")
		return res, err
	}

	res = view_models.NewBrandVm().BuildDetail(&brand)
	return res, nil
}

func (uc BrandUsecase) Update(req *request.BrandRequest, brandID uuid.UUID) (res view_models.BrandDetailVm, err error) {
	repository := command.NewCommandBrandRepository(uc.DB)

	userID, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}
	var brandSocmed []models.BrandMediaSocial
	for _, socmed := range req.MediaSocial {
		brandSocmed = append(brandSocmed, models.BrandMediaSocial{
			BrandID: brandID,
			Type:    socmed.Type,
			Link:    socmed.Link,
		})
	}
	model := models.Brand{
		ID:              brandID,
		Name:            req.Name,
		EstablishedDate: datetime.StrParseToTime(req.EstablishedDate, "2006-10-02"),
		Title:           req.Title,
		Catchphrase:     req.Catchphrase,
		About:           req.About,
		LogoID:          req.LogoID,
		BannerWebID:     req.BannerWebID,
		BannerMobileID:  req.BannerMobileID,
		UpdatedBy:       &userID,
		MediaSocials:    brandSocmed,
	}
	tx := uc.DB.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	brand, err := repository.Update(model, tx)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-brand")
		return res, err
	}

	res = view_models.NewBrandVm().BuildDetail(&brand)
	tx.Commit()
	return res, nil
}

func (uc BrandUsecase) Delete(brandID uuid.UUID) (err error) {
	repository := command.NewCommandBrandRepository(uc.DB)

	userID, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}
	model := models.Brand{
		ID:        brandID,
		DeletedBy: &userID,
	}
	tx := uc.DB.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return err
	}
	err = repository.Delete(model, tx)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-brand")
		return err
	}

	return nil
}

func (uc BrandUsecase) Export(fileType string) (err error) {
	panic("Under Maintenance")
}

func (uc BrandUsecase) Banned(req *request.BannedBrandRequest, brandID uuid.UUID) (res view_models.BrandDetailVm, err error) {
	db := uc.DB
	repo := command.NewCommandBrandRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	model := models.Brand{
		ID:     brandID,
		Status: req.Status,
	}
	if req.Status == "Banned" {
		model.BannedReason = req.Reason
		model.BannedDocumentID = req.DocID
		model.BannedBy = &userId
	} else {
		model.UnbannedReason = req.Reason
		model.UnbannedDocumentID = req.DocID
		model.UnbannedBy = &userId
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	banned, err := repo.Update(model, tx)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-banned-brand")
		return res, err
	}

	res = view_models.NewBrandVm().BuildDetail(&banned)
	tx.Commit()
	return res, nil
}
