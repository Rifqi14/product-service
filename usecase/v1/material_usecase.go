package v1

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
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
		Name:               req.Name,
		ParentID:           req.ParentID,
		MaterialCategoryID: &req.MaterialCategoryID,
		CreatedBy:          &userId,
		UpdatedBy:          &userId,
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
	var dataPath []string
	path, err := uc.createPath(&material.ID, dataPath)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-material")
		return res, err
	}
	material.Path = strings.Join(path, " / ")
	material.Level = int64(len(path))
	err = tx.Save(&material).Error
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-material")
		return res, err
	}

	tx.Commit()
	res, _ = uc.Detail(material.ID)
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

	material, err := repo.Detail(materialId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-material")
		return res, err
	}

	if material != nil {
		res = view_models.NewMaterialVm().BuildDetail(material)
	}

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
		ID:                 materialId,
		Name:               req.Name,
		ParentID:           req.ParentID,
		MaterialCategoryID: &req.MaterialCategoryID,
		UpdatedBy:          &userId,
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return res, err
	}
	material, err := repo.Update(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-update-material")
		return res, err
	}

	// var dataPath []string
	// path, err := uc.createPath(&material.ID, dataPath)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-material")
	// 	return res, err
	// }
	// material.Path = strings.Join(path, " / ")
	// material.Level = int64(len(path))
	// err = tx.Updates(&material).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-material")
	// 	return res, err
	// }
	// err = uc.updatePath(material.ID)
	// if err != nil {
	// 	tx.Rollback()
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-updatePath-transaction")
	// 	return res, err
	// }

	tx.Commit()
	res, _ = uc.Detail(material.ID)
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
	count := tx.Model(model).Association("Products").Count()
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, "data in used", functioncaller.PrintFuncName(), "data-in-used")
		return errors.New("data in used")
	}
	err = repo.Delete(model)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-delete-material")
		return err
	}

	return nil
}

func (uc MaterialUsecase) Export(fileType string) (link *fileVm.FileVm, err error) {
	db := uc.DB
	repo := query.NewQueryMaterialRepository(db)

	materials, err := repo.All()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get-all-materials")
		return nil, err
	}
	materialsVm := view_models.NewMaterialVm().BuildExport(materials)
	f := excelize.NewFile()
	sheet := "Material"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// Set header table
	f.SetCellValue(sheet, "A1", "Kategori Material")
	f.SetCellValue(sheet, "B1", "Nama Material")
	f.SetCellValue(sheet, "C1", "Parent Material")

	for i, materialVm := range materialsVm {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), materialVm.Category)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), materialVm.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), materialVm.Parent)
	}
	filename := fmt.Sprintf("%d_material.xlsx", time.Now().Unix())
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

func (uc MaterialUsecase) createPath(materialId *uuid.UUID, path []string) (paths []string, err error) {
	repo := query.NewQueryMaterialRepository(uc.DB)

	material, err := repo.Detail(*materialId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-detail-material")
		return nil, err
	}
	path = append([]string{material.Name}, path...)
	if material.Parent != nil {
		path, err := uc.createPath(material.ParentID, path)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-material")
			return nil, err
		}
		return path, nil
	}
	return path, nil
}

func (uc MaterialUsecase) updatePath(materialId uuid.UUID) error {
	db := uc.DB
	queryRepo := query.NewQueryMaterialRepository(db)
	commandRepo := command.NewCommandMaterialRepository(db)

	materials, err := queryRepo.Parent(materialId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-updatePath-material")
		return err
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return err
	}
	for _, material := range materials {
		var parentPath []string
		path, _ := uc.createPath(&material.ID, parentPath)
		userId := uuid.MustParse(uc.UserID)
		model := models.Material{
			ID:        material.ID,
			Path:      strings.Join(path, " / "),
			Level:     int64(len(path)),
			UpdatedBy: &userId,
		}
		_, err := commandRepo.Update(model)
		if err != nil {
			tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-material")
			return err
		}
		err = uc.updatePath(material.ID)
		if err != nil {
			tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-material")
			return err
		}
	}
	tx.Commit()
	return nil
}
