package v1

import (
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
	var dataPath []string
	path, err := uc.createPath(&label.ID, dataPath)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-label")
		return res, err
	}
	label.Path = strings.Join(path, " / ")
	label.Level = int64(len(path))
	err = tx.Save(&label).Error
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-label")
		return res, err
	}

	tx.Commit()
	res, _ = uc.Detail(label.ID)
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

	label, err := repo.Detail(labelId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-detail-label")
		return res, err
	}

	if label != nil {
		res = view_models.NewLabelVm().BuildDetail(label)
	}

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

	var dataPath []string
	path, err := uc.createPath(&label.ID, dataPath)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-label")
		return res, err
	}
	label.Path = strings.Join(path, " / ")
	label.Level = int64(len(path))
	err = tx.Updates(&label).Error
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-label")
		return res, err
	}
	err = uc.updatePath(label.ID)
	if err != nil {
		tx.Rollback()
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-updatePath-transaction")
		return res, err
	}

	tx.Commit()
	res, _ = uc.Detail(label.ID)
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

func (uc LabelUsecase) Export(fileType string) (link *fileVm.FileVm, err error) {
	db := uc.DB
	repo := query.NewQueryLabelRepository(db)

	labels, err := repo.All()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "get-all-labels")
		return nil, err
	}
	labelsVm := view_models.NewLabelVm().BuildExport(labels)
	f := excelize.NewFile()
	sheet := "Label"
	f.SetSheetName(f.GetSheetName(0), sheet)

	// Set header table
	f.SetCellValue(sheet, "A1", "Nama Label")
	f.SetCellValue(sheet, "B1", "Parent Label")

	for i, labelVm := range labelsVm {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), labelVm.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), labelVm.Parent)
	}
	filename := fmt.Sprintf("%d_label.xlsx", time.Now().Unix())
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

func (uc LabelUsecase) createPath(labelId *uuid.UUID, path []string) (paths []string, err error) {
	repo := query.NewQueryLabelRepository(uc.DB)

	label, err := repo.Detail(*labelId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-detail-label")
		return nil, err
	}
	path = append([]string{label.Name}, path...)
	if label.Parent != nil {
		path, err := uc.createPath(label.ParentID, path)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-createPath-label")
			return nil, err
		}
		return path, nil
	}
	return path, nil
}

func (uc LabelUsecase) updatePath(labelId uuid.UUID) error {
	db := uc.DB
	queryRepo := query.NewQueryLabelRepository(db)
	commandRepo := command.NewCommandLabelRepository(db)

	labels, err := queryRepo.Parent(labelId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-updatePath-label")
		return err
	}
	tx := db.Begin()
	if err := tx.Error; err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "gorm-start-transaction")
		return err
	}
	for _, label := range labels {
		var parentPath []string
		path, _ := uc.createPath(&label.ID, parentPath)
		userId := uuid.MustParse(uc.UserID)
		model := models.Label{
			ID:        label.ID,
			Path:      strings.Join(path, " / "),
			Level:     int64(len(path)),
			UpdatedBy: &userId,
		}
		_, err := commandRepo.Update(model)
		if err != nil {
			tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-label")
			return err
		}
		err = uc.updatePath(label.ID)
		if err != nil {
			tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-updatePath-label")
			return err
		}
	}
	tx.Commit()
	return nil
}
