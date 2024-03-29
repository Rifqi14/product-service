package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type MaterialListVm struct {
	Category string `json:"material_category"`
	Parent   string `json:"parent"`
	Child    string `json:"child"`
	ID       string `json:"material_id"`
}

type MaterialDetailVm struct {
	ID               string                   `json:"material_id"`
	Name             string                   `json:"name"`
	MaterialCategory MaterialCategoryDetailVm `json:"material_category"`
	Parent           *MaterialParentVm        `json:"material_parent"`
}

type MaterialExportVm struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Parent   string `json:"parent"`
}

type MaterialParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type MaterialVm struct {
	List   MaterialListVm     `json:"list_material"`
	Detail MaterialDetailVm   `json:"detail_material"`
	Export []MaterialExportVm `json:"export_material"`
}

func NewMaterialVm() MaterialVm {
	return MaterialVm{}
}

func (vm MaterialVm) BuildExport(materials []models.Material) (res []MaterialExportVm) {
	for _, material := range materials {
		var materialVm MaterialExportVm
		materialVm.Name = material.Name
		materialVm.Category = material.Category.Name
		if material.Parent != nil {
			materialVm.Parent = material.Parent.Name
		}
		res = append(res, materialVm)
	}
	return res
}

func (vm MaterialVm) BuildList(materials []models.Material) (res []MaterialListVm) {
	for _, material := range materials {
		var parent string
		var category string
		if material.Parent != nil {
			parent = material.Parent.Name
		}
		if material.Category.Name != "" {
			category = material.Category.Name
		}
		res = append(res, MaterialListVm{
			Category: category,
			Parent:   parent,
			Child:    material.Name,
			ID:       material.ID.String(),
		})
	}
	return res
}

func (vm MaterialVm) BuildDetail(material *models.Material) (res MaterialDetailVm) {
	if material != nil {
		res = MaterialDetailVm{
			ID:               material.ID.String(),
			Name:             material.Name,
			MaterialCategory: *NewMaterialCategoryVm().BuildDetail(&material.Category),
			Parent:           vm.BuildParent(material.Parent),
		}
	}
	return res
}

func (vm MaterialVm) BuildParent(parent *models.Material) (res *MaterialParentVm) {
	if parent != nil {
		res = &MaterialParentVm{
			ID:   parent.ID.String(),
			Name: parent.Name,
		}
	}
	return res
}
