package v1

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/command"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/query"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"

	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	ucinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/usecase"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

type ProductUsecase struct {
	*usecase.Contract
}

func NewProductUsecase(contract *usecase.Contract) ucinterface.IProductUsecase {
	return &ProductUsecase{Contract: contract}
}

func (uc ProductUsecase) Create(req *request.ProductRequest) (res *view_models.ProductVm, err error) {
	db := uc.DB
	repo := command.NewCommandProductRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	var isDisplayed bool = true
	if len(req.Variants) > 0 {
		totalStock := 0
		for _, variants := range req.Variants {
			for _, detail := range variants.Details {
				totalStock = totalStock + int(detail.Stock)
			}
		}
		if totalStock <= 0 {
			isDisplayed = !isDisplayed
		}
	}

	model := models.Product{
		Name:          req.Name,
		BrandID:       req.BrandID,
		NormalPrice:   req.NormalPrice,
		StripePrice:   req.StripePrice,
		DiscountPrice: req.Discount,
		FinalPrice:    req.StripePrice - req.Discount,
		Description:   req.Description,
		Measurement:   req.Measurement,
		Length:        int(req.PackageDimension.Length),
		Width:         int(req.PackageDimension.Width),
		Height:        int(req.PackageDimension.Height),
		PoStatus:      req.PreOrder.Status,
		PoDay:         int(req.PreOrder.PreOrderDay),
		Categories:    uc.appendCategories(req),
		Labels:        uc.appendLabels(req.Labels),
		Materials:     uc.appendMaterials(req.Materials),
		Genders:       uc.appendGenders(req.Gender),
		Colors:        uc.appendColors(req.Variants),
		Variants:      uc.appendVariantDetails(req.Variants),
		Images:        uc.appendVariantImages(req.Variants),
		IsDisplayed:   &isDisplayed,
		CreatedBy:     &userId,
		UpdatedBy:     &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	product, err := repo.Create(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-create-product")
		return res, err
	}

	res, _ = uc.Detail(product.ID)
	tx.Commit()
	return res, nil
}

func (uc ProductUsecase) List(req *request.FilterProductRequest) (res []*view_models.ProductVm, pagination view_models.PaginationVm, err error) {
	db := uc.DB
	repo := query.NewQueryProductRepository(db)

	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Offset, req.Limit, req.OrderBy, req.Sort)

	products, count, err := repo.List(req.Search, orderBy, sort, req.ProductName, limit, offset, req.MinPrice, req.MaxPrice, req.Brand, req.Product, req.Color)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-list-material")
		return res, pagination, err
	}

	res = view_models.NewProductVm().BuildList(products)

	pagination = uc.SetPaginationResponse(page, limit, count)
	return res, pagination, nil
}

func (uc ProductUsecase) Detail(productId uuid.UUID) (res *view_models.ProductVm, err error) {
	db := uc.DB
	repo := query.NewQueryProductRepository(db)

	product, err := repo.Detail(productId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-product")
		return res, err
	}

	if product.Name == "" {
		return nil, errors.New(messages.DataNotFound)
	}
	res = view_models.NewProductVm().BuildDetail(product)

	return res, nil
}

func (uc ProductUsecase) Update(req *request.ProductRequest, productId uuid.UUID) (res *view_models.ProductVm, err error) {
	db := uc.DB
	repo := command.NewCommandProductRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return res, err
	}

	var isDisplayed bool = true
	if len(req.Variants) > 0 {
		totalStock := 0
		for _, variants := range req.Variants {
			for _, detail := range variants.Details {
				totalStock = totalStock + int(detail.Stock)
			}
		}
		if totalStock <= 0 {
			isDisplayed = !isDisplayed
		}
	}

	model := models.Product{
		ID:            productId,
		Name:          req.Name,
		BrandID:       req.BrandID,
		NormalPrice:   req.NormalPrice,
		StripePrice:   req.StripePrice,
		DiscountPrice: req.Discount,
		FinalPrice:    req.StripePrice - req.Discount,
		Description:   req.Description,
		Measurement:   req.Measurement,
		Length:        int(req.PackageDimension.Length),
		Width:         int(req.PackageDimension.Width),
		Height:        int(req.PackageDimension.Height),
		PoStatus:      req.PreOrder.Status,
		PoDay:         int(req.PreOrder.PreOrderDay),
		IsDisplayed:   &isDisplayed,
		CreatedBy:     &userId,
		UpdatedBy:     &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return nil, err
	}
	product, err := repo.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-product")
		return nil, err
	}
	categoriesAssoc := tx.Model(&product).Association("Categories")
	labelsAssoc := tx.Model(&product).Association("Labels")
	materialsAssoc := tx.Model(&product).Association("Materials")
	gendersAssoc := tx.Model(&product).Association("Genders")
	colorsAssoc := tx.Model(&product).Association("Colors")
	variantsAssoc := tx.Model(&product).Association("ProductVariants")
	imagesAssoc := tx.Model(&product).Association("ProductImages")
	categoriesAssoc.Clear()
	labelsAssoc.Clear()
	materialsAssoc.Clear()
	gendersAssoc.Clear()
	colorsAssoc.Clear()
	variantsAssoc.Clear()
	imagesAssoc.Clear()
	tx.Model(&product).Association("Categories").Append(uc.appendCategories(req))
	tx.Model(&product).Association("Labels").Append(uc.appendLabels(req.Labels))
	tx.Model(&product).Association("Materials").Append(uc.appendMaterials(req.Materials))
	tx.Model(&product).Association("Genders").Append(uc.appendGenders(req.Gender))
	tx.Model(&product).Association("Colors").Append(uc.appendColors(req.Variants))
	tx.Model(&product).Association("Variants").Append(uc.appendVariantDetails(req.Variants))
	tx.Model(&product).Association("Images").Append(uc.appendVariantImages(req.Variants))

	tx.Commit()
	res, _ = uc.Detail(product.ID)
	return res, nil
}

func (uc ProductUsecase) Delete(productId uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandProductRepository(db)

	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "parse-userId-toUuid")
		return err
	}

	model := models.Product{
		ID:        productId,
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

	tx.Commit()
	return nil
}

func (uc ProductUsecase) ChangeStatus(req *request.BannedProductRequest, productId *uuid.UUID) (err error) {
	db := uc.DB
	repo := command.NewCommandProductRepository(db)
	userId, err := uuid.Parse(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "error-parse-toUuid")
		return err
	}
	product, err := uc.Detail(*productId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-check-product")
		return err
	}
	if product == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataNotFound, functioncaller.PrintFuncName(), "data-product-notFound")
		return errors.New(messages.DataNotFound)
	}
	if !product.IsDisplayed {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataNotFound, functioncaller.PrintFuncName(), "data-already-inActive")
		return errors.New("product already inactive")
	}

	status := "Active"
	if !req.Status {
		status = "Inactive"
	}
	attachmentId := uuid.MustParse(req.AttachmentID)

	modelProduct := models.Product{
		ID:          *productId,
		IsDisplayed: &req.Status,
		UpdatedBy:   &userId,
		Logs: []models.ProductLog{
			{
				ProductID:    productId,
				Reason:       req.Reason,
				Status:       status,
				AttachmentID: &attachmentId,
				UpdatedBy:    &userId,
			},
		},
	}
	tx := db.Begin()
	_, err = repo.Update(modelProduct)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "update-data-error")
		return err
	}
	tx.Commit()
	return nil
}

func (uc ProductUsecase) Export(fileType string) (err error) {
	panic("Under development")
}

func (uc ProductUsecase) appendCategories(req *request.ProductRequest) (res []*models.Category) {
	for _, category := range req.Categories {
		res = append(res, &models.Category{
			ID: *category,
		})
	}
	return res
}

func (uc ProductUsecase) appendLabels(labels []*uuid.UUID) (res []*models.Label) {
	if len(labels) > 0 {
		for _, label := range labels {
			res = append(res, &models.Label{
				ID: *label,
			})
		}
	}
	return res
}

func (uc ProductUsecase) appendMaterials(materials []*uuid.UUID) (res []*models.Material) {
	if len(materials) > 0 {
		for _, material := range materials {
			res = append(res, &models.Material{
				ID: *material,
			})
		}
	}
	return res
}

func (uc ProductUsecase) appendGenders(genders []*uuid.UUID) (res []*models.Gender) {
	if len(genders) > 0 {
		for _, gender := range genders {
			res = append(res, &models.Gender{
				ID: gender,
			})
		}
	}
	return res
}

func (uc ProductUsecase) appendColors(colors []*request.VariantRequest) (res []*models.Color) {
	if len(colors) > 0 {
		for _, color := range colors {
			res = append(res, &models.Color{
				ID: *color.Colors,
			})
		}
	}
	return res
}

func (uc ProductUsecase) appendVariantDetails(variants []*request.VariantRequest) (res []*models.ProductVariant) {
	if len(variants) > 0 {
		for _, variant := range variants {
			for _, detail := range variant.Details {
				res = append(res, &models.ProductVariant{
					ColorID: *variant.Colors,
					Size:    detail.Size,
					Stock:   detail.Stock,
					Sku:     &detail.Sku,
					Status:  detail.Status,
				})
			}
		}
	}
	return res
}

func (uc ProductUsecase) appendVariantImages(variants []*request.VariantRequest) (res []*models.ProductVariantImage) {
	if len(variants) > 0 {
		for _, variant := range variants {
			for _, image := range variant.Images {
				res = append(res, &models.ProductVariantImage{
					ColorID: *variant.Colors,
					Look:    image.Look,
					ImageID: image.ImageID,
				})
			}
		}
	}
	return res
}
