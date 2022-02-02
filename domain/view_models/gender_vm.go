package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type GenderListVm struct {
	No     int64  `json:"no"`
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

type GenderVm struct {
	List   GenderListVm   `json:"list_gender"`
	Detail GenderDetailVm `json:"detail_gender"`
}

func NewGenderVm() GenderVm {
	return GenderVm{}
}

func (vm GenderVm) BuildList(genders []models.Gender) (res []GenderListVm) {
	for i, gender := range genders {
		res = append(res, GenderListVm{
			No:     int64(i + 1),
			Parent: gender.Parent.Name,
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

func (vm GenderVm) BuildParent(parent *models.Gender) *GenderParentVm {
	return &GenderParentVm{
		ID:   parent.ID.String(),
		Name: parent.Name,
	}
}
