package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type MaterialListVm struct {
	No     int64  `json:"no"`
	Parent string `json:"parent"`
	Child  string `json:"child"`
	ID     string `json:"material_id"`
}

type MaterialDetailVm struct {
	ID               string                   `json:"material_id"`
	Name             string                   `json:"name"`
	MaterialCategory MaterialCategoryDetailVm `json:"material_category"`
	Parent           *MaterialParentVm        `json:"material_parent"`
}

type MaterialParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type MaterialVm struct {
	List   MaterialListVm   `json:"list_material"`
	Detail MaterialDetailVm `json:"detail_material"`
}

func NewMaterialVm() MaterialVm {
	return MaterialVm{}
}

func (vm MaterialVm) BuildList(materials []models.Material) (res []MaterialListVm) {
	for i, material := range materials {
		res = append(res, MaterialListVm{
			No:     int64(i + 1),
			Parent: material.Parent.Name,
			Child:  material.Name,
			ID:     material.ID.String(),
		})
	}
	return res
}

func (vm MaterialVm) BuildDetail(material *models.Material) MaterialDetailVm {
	return MaterialDetailVm{
		ID:               material.ID.String(),
		Name:             material.Name,
		MaterialCategory: NewMaterialCategoryVm().BuildDetail(&material.Category),
		Parent:           vm.BuildParent(material.Parent),
	}
}

func (vm MaterialVm) BuildParent(parent *models.Material) *MaterialParentVm {
	return &MaterialParentVm{
		ID:   parent.ID.String(),
		Name: parent.Name,
	}
}
