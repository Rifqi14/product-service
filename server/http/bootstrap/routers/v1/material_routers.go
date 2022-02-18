package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type MaterialRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route MaterialRouters) MaterialRoute() {
	handlerMaterialCat := v1.NewMaterialCategoryHandler(route.Handler)
	handlerMaterial := v1.NewMaterialHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	materialRoutes := route.RouteGroup.Group("/material")
	categoryRoutes := materialRoutes.Group("/category")
	materialRoutes.Get("", handlerMaterial.List)
	categoryRoutes.Get("", handlerMaterialCat.List)
	materialRoutes.Get("/:id", handlerMaterial.Detail)
	categoryRoutes.Get("/:id", handlerMaterialCat.Detail)

	materialRoutes.Use(jwtMiddleware.New)
	materialRoutes.Post("", handlerMaterial.Create)
	materialRoutes.Patch("/:id", handlerMaterial.Update)
	materialRoutes.Delete("/:id", handlerMaterial.Delete)
	materialRoutes.Post("/export", handlerMaterial.Export)
	categoryRoutes.Post("", handlerMaterialCat.Create)
	categoryRoutes.Patch("/:id", handlerMaterialCat.Update)
	categoryRoutes.Delete("/:id", handlerMaterialCat.Delete)
	categoryRoutes.Post("/export", handlerMaterialCat.Export)
}
