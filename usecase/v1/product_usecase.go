package v1

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/command"
	"gitlab.com/s2.1-backend/shm-product-svc/repositories/query"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"

	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/request"
	ucinterface "gitlab.com/s2.1-backend/shm-product-svc/domain/usecase"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
)

var headerExportProduct = []string{"Nama Produk", "Brand Produk", "Kategori", "Label", "Material", "Gender", "Harga Normal", "Harga Coret", "Diskon/Pot.Harga", "Warna", "Ukuran", "stok", "nomor SKU", "deskripsi produk", "detail ukuran", "panjang paket", "lebar paket", "tinggi paket", "waktu proses preorder"}

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

func (uc ProductUsecase) FindBy(req *request.FindByRequest) (res []*view_models.ProductVm, pagination view_models.PaginationVm, err error) {
	// db := uc.DB
	// repo := query.NewQueryProductRepository(db)

	// offset, limit, page, orderBy, sort := uc.SetPaginationParameter(req.Pagination.Offset, req.Pagination.Limit, req.Pagination.OrderBy, req.Pagination.Sort)

	// products, count, err := repo.FindBy()
	panic("Under development")
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
		Logs: []*models.ProductLog{
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

func (uc ProductUsecase) Export(fileType string) (link *fileVm.FileVm, err error) {
	db := uc.DB
	repo := query.NewQueryProductRepository(db)

	products, err := repo.All()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get-all-products")
		return nil, err
	}
	productsVm := view_models.NewProductVm().BuildExport(products)
	f := excelize.NewFile()
	sheet := "Product"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// Set header table
	f.SetCellValue(sheet, "A1", headerExportProduct[0])
	f.SetCellValue(sheet, "B1", headerExportProduct[1])
	f.SetCellValue(sheet, "C1", headerExportProduct[2])
	f.SetCellValue(sheet, "D1", headerExportProduct[3])
	f.SetCellValue(sheet, "E1", headerExportProduct[4])
	f.SetCellValue(sheet, "F1", headerExportProduct[5])
	f.SetCellValue(sheet, "G1", headerExportProduct[6])
	f.SetCellValue(sheet, "H1", headerExportProduct[7])
	f.SetCellValue(sheet, "I1", headerExportProduct[8])
	f.SetCellValue(sheet, "J1", headerExportProduct[9])
	f.SetCellValue(sheet, "K1", headerExportProduct[10])
	f.SetCellValue(sheet, "L1", headerExportProduct[11])
	f.SetCellValue(sheet, "M1", headerExportProduct[12])
	f.SetCellValue(sheet, "N1", headerExportProduct[13])
	f.SetCellValue(sheet, "O1", headerExportProduct[14])
	f.SetCellValue(sheet, "P1", headerExportProduct[15])
	f.SetCellValue(sheet, "Q1", headerExportProduct[16])
	f.SetCellValue(sheet, "R1", headerExportProduct[17])
	f.SetCellValue(sheet, "S1", headerExportProduct[18])

	for i, productVm := range productsVm {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), productVm.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), productVm.Brand)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), productVm.Category)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), productVm.Label)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), productVm.Material)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), productVm.Gender)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), productVm.NormalPrice)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", i+2), productVm.StripePrice)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", i+2), productVm.Discount)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", i+2), productVm.Color)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", i+2), productVm.Size)
		f.SetCellValue(sheet, fmt.Sprintf("L%d", i+2), productVm.Stock)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", i+2), productVm.SKU)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", i+2), productVm.Description)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", i+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("P%d", i+2), productVm.Length)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", i+2), productVm.Width)
		f.SetCellValue(sheet, fmt.Sprintf("R%d", i+2), productVm.Height)
		f.SetCellValue(sheet, fmt.Sprintf("S%d", i+2), productVm.PoDay)
	}
	filename := fmt.Sprintf("%d_product.xlsx", time.Now().Unix())
	if err := f.SaveAs("../../domain/files/" + filename); err != nil {
		return nil, err
	}
	link, err = uc.ExportBase(filename)
	if err != nil {
		return nil, err
	}
	err = os.Remove("../../domain/files/" + filename)
	if err != nil {
		return nil, err
	}
	return link, nil
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
