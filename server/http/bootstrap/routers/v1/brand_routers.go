package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type BrandRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route BrandRouters) BrandRoute() {
	handler := v1.NewBrandHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	brandRoutes := route.RouteGroup.Group("/brand")
	brandRoutes.Get("", handler.List)
	brandRoutes.Get("/:id", handler.Detail)
	brandRoutes.Use(jwtMiddleware.New)
	brandRoutes.Post("", handler.Create)
	brandRoutes.Post("/ban/:id", handler.Banned)
	brandRoutes.Patch("/:id", handler.Update)
	brandRoutes.Delete("/:id", handler.Delete)
}
