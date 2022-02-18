package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type MaterialCategoryDetailVm struct {
	ID   string `json:"material_category_id"`
	Name string `json:"name"`
}

type MaterialCategoryExportVm struct {
	Name string `json:"name"`
}

type MaterialCategoryVm struct {
	Detail MaterialCategoryDetailVm `json:"detail_material_category"`
	Export MaterialCategoryExportVm `json:"export_material_category"`
}

func NewMaterialCategoryVm() MaterialCategoryVm {
	return MaterialCategoryVm{}
}

func (vm MaterialCategoryVm) BuildExport(categories []models.MaterialCategory) (res []MaterialCategoryExportVm) {
	for _, category := range categories {
		var categoryVm MaterialCategoryExportVm
		categoryVm.Name = category.Name
		res = append(res, categoryVm)
	}
	return res
}

func (vm MaterialCategoryVm) BuildDetail(materialCat *models.MaterialCategory) (res *MaterialCategoryDetailVm) {
	if materialCat != nil {
		res = &MaterialCategoryDetailVm{
			ID:   materialCat.ID.String(),
			Name: materialCat.Name,
		}
	}
	return res
}
