package view_models

import (
	adminVm "gitlab.com/s2.1-backend/shm-auth-svc/domain/view_models"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/models"
)

type ProductLogVm struct {
	ID         string          `json:"product_log_id"`
	Reason     string          `json:"reason"`
	Status     string          `json:"status"`
	ProductID  string          `json:"product_id"`
	Attachment *fileVm.FileVm  `json:"attachment"`
	Verifier   adminVm.AdminVm `json:"verifier"`
}

type LogVm struct {
	List   []ProductLogVm `json:"product_logs"`
	Detail ProductLogVm   `json:"product_log"`
}

func NewProductLogVm() LogVm {
	return LogVm{}
}

func (vm LogVm) BuildList(logs []*models.ProductLog) (res []*ProductLogVm) {
	if len(logs) > 0 {
		for _, log := range logs {
			var attachment fileVm.FileVm
			if log.Attachment != nil {
				attachment = fileVm.NewFileVm().Build(*log.Attachment)
			}
			var verifier adminVm.AdminVm
			if log.Verifier != nil {
				verifier = adminVm.NewAdminVm().Build(log.Verifier)
			}
			res = append(res, &ProductLogVm{
				ID:         log.ID.String(),
				Reason:     log.Reason,
				Status:     log.Status,
				ProductID:  log.ProductID.String(),
				Verifier:   verifier,
				Attachment: &attachment,
			})
		}
	}
	return res
}

func (vm LogVm) BuildDetail(log *models.ProductLog) (res *ProductLogVm) {
	if log != nil {
		var attachment fileVm.FileVm
		if log.Attachment != nil {
			attachment = fileVm.NewFileVm().Build(*log.Attachment)
		}
		res = &ProductLogVm{
			ID:         log.ID.String(),
			Reason:     log.Reason,
			Status:     log.Status,
			ProductID:  log.ProductID.String(),
			Verifier:   adminVm.NewAdminVm().Build(log.Verifier),
			Attachment: &attachment,
		}
	}
	return res
}
