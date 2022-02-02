package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type LabelListVm struct {
	Parent string `json:"parent"`
	Child  string `json:"child"`
	ID     string `json:"label_id"`
}

type LabelDetailVm struct {
	ID     string         `json:"label_id"`
	Name   string         `json:"name"`
	Parent *LabelParentVm `json:"label_parent"`
}

type LabelParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type LabelVm struct {
	List   LabelListVm   `json:"list_label"`
	Detail LabelDetailVm `json:"detail_label"`
}

func NewLabelVm() LabelVm {
	return LabelVm{}
}

func (vm LabelVm) BuildList(labels []models.Label) (res []LabelListVm) {
	for _, label := range labels {
		var parent string
		if label.Parent != nil {
			parent = label.Parent.Name
		}
		res = append(res, LabelListVm{
			Parent: parent,
			Child:  label.Name,
			ID:     label.ID.String(),
		})
	}
	return res
}

func (vm LabelVm) BuildDetail(label *models.Label) LabelDetailVm {
	return LabelDetailVm{
		ID:     label.ID.String(),
		Name:   label.Name,
		Parent: vm.BuildParent(label.Parent),
	}
}

func (vm LabelVm) BuildParent(parent *models.Label) (res *LabelParentVm) {
	if parent != nil {
		res = &LabelParentVm{
			ID:   parent.ID.String(),
			Name: parent.Name,
		}
	}
	return res
}
