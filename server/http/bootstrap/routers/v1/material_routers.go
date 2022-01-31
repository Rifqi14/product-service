package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type MaterialCategory struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route MaterialCategory) MaterialRoute() {
	handlerMaterialCat := v1.NewMaterialCategoryHandler(route.Handler)
	handlerMaterial := v1.NewMaterialHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	materialRoutes := route.RouteGroup.Group("/material")
	materialRoutes.Get("", handlerMaterial.List)
	materialRoutes.Get("/:id", handlerMaterial.Detail)
	materialRoutes.Use(jwtMiddleware.New)
	materialRoutes.Post("", handlerMaterial.Create)
	materialRoutes.Patch("/:id", handlerMaterial.Update)
	materialRoutes.Delete("/:id", handlerMaterial.Delete)

	categoryRoutes := materialRoutes.Group("/category")
	categoryRoutes.Get("", handlerMaterialCat.List)
	categoryRoutes.Get("/:id", handlerMaterialCat.Detail)
	categoryRoutes.Use(jwtMiddleware.New)
	categoryRoutes.Post("", handlerMaterialCat.Create)
	categoryRoutes.Patch("/:id", handlerMaterialCat.Update)
	categoryRoutes.Delete("/:id", handlerMaterialCat.Delete)
}
