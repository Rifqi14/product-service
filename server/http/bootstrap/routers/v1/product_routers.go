package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type ProductRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ProductRouters) ProductRoute() {
	handler := v1.NewProductHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	productRoutes := route.RouteGroup.Group("/product")
	productRoutes.Get("/public", handler.List)
	productRoutes.Get("/public/:id", handler.Detail)

	productRoutes.Get("", handler.List)
	productRoutes.Get("/:id", handler.Detail)

	productRoutes.Use(jwtMiddleware.New)
	productRoutes.Post("", handler.Create)
	productRoutes.Post("/status/:id", handler.ChangeStatus)
	productRoutes.Patch("/:id", handler.Update)
	productRoutes.Delete("/:id", handler.Delete)
	productRoutes.Post("/export", handler.Export)
}
