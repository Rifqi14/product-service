package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type CategoryRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route CategoryRouters) CategoryRoute() {
	handler := v1.NewCategoryHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	categoryRoutes := route.RouteGroup.Group("/category")
	categoryRoutes.Get("", handler.List)
	categoryRoutes.Get("/:id", handler.Detail)
	categoryRoutes.Use(jwtMiddleware.New)
	categoryRoutes.Post("", handler.Create)
	categoryRoutes.Patch("/:id", handler.Update)
	categoryRoutes.Delete("/:id", handler.Delete)
}
