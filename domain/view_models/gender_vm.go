package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type GenderListVm struct {
	Parent string `json:"parent"`
	Child  string `json:"child"`
	ID     string `json:"gender_id"`
}

type GenderDetailVm struct {
	ID     string          `json:"gender_id"`
	Name   string          `json:"name"`
	Parent *GenderParentVm `json:"gender_parent"`
}

type GenderParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type GenderExportVm struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type GenderVm struct {
	List   GenderListVm   `json:"list_gender"`
	Detail GenderDetailVm `json:"detail_gender"`
}

func NewGenderVm() GenderVm {
	return GenderVm{}
}

func (vm GenderVm) BuildExport(genders []models.Gender) (res []GenderExportVm) {
	for _, gender := range genders {
		var genderVm GenderExportVm
		genderVm.Name = gender.Name
		if gender.Parent != nil {
			genderVm.Parent = gender.Parent.Name
		}
		res = append(res, genderVm)
	}
	return res
}

func (vm GenderVm) BuildList(genders []models.Gender) (res []GenderListVm) {
	for _, gender := range genders {
		var parent string
		if gender.Parent != nil {
			parent = gender.Parent.Name
		}
		res = append(res, GenderListVm{
			Parent: parent,
			Child:  gender.Name,
			ID:     gender.ID.String(),
		})
	}
	return res
}

func (vm GenderVm) BuildDetail(gender *models.Gender) GenderDetailVm {
	return GenderDetailVm{
		ID:     gender.ID.String(),
		Name:   gender.Name,
		Parent: vm.BuildParent(gender.Parent),
	}
}

func (vm GenderVm) BuildParent(parent *models.Gender) (res *GenderParentVm) {
	if parent != nil {
		res = &GenderParentVm{
			ID:   parent.ID.String(),
			Name: parent.Name,
		}
	}
	return res
}
