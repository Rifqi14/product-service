package request

type MaterialCategoryRequest struct {
	Name string `form:"name" json:"name" validate:"required"`
}
