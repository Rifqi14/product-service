package view_models

import "gitlab.com/s2.1-backend/shm-product-svc/domain/models"

type ColorListVm struct {
	No     int64  `json:"no"`
	Parent string `json:"parent"`
	Child  string `json:"child"`
	Hex    string `json:"hex"`
	ID     string `json:"color_id"`
}

type ColorDetailVm struct {
	ID     string         `json:"color_id"`
	Name   string         `json:"name"`
	Hex    string         `json:"hex"`
	Parent *ColorParentVm `json:"color_parent"`
}

type ColorParentVm struct {
	ID   string `json:"parent_id"`
	Name string `json:"name"`
}

type ColorVm struct {
	List   ColorListVm   `json:"list_color"`
	Detail ColorDetailVm `json:"detail_color"`
}

func NewColorVm() ColorVm {
	return ColorVm{}
}

func (vm ColorVm) BuildList(colors []models.Color) (res []ColorListVm) {
	for i, color := range colors {
		var parent string
		if color.Parent != nil {
			parent = color.Parent.Name
		}
		res = append(res, ColorListVm{
			No:     int64(i + 1),
			Parent: parent,
			Child:  color.Name,
			Hex:    color.RgbCode,
			ID:     color.ID.String(),
		})
	}
	return res
}

func (vm ColorVm) BuildDetail(color *models.Color) ColorDetailVm {
	return ColorDetailVm{
		ID:     color.ID.String(),
		Name:   color.Name,
		Hex:    color.RgbCode,
		Parent: vm.BuildParent(color.Parent),
	}
}

func (vm ColorVm) BuildParent(parent *models.Color) *ColorParentVm {
	var res *ColorParentVm
	if parent != nil {
		res = &ColorParentVm{
			ID:   parent.ID.String(),
			Name: parent.Name,
		}
	}
	return res
}
