package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type LabelRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route LabelRouters) LabelRoute() {
	handler := v1.NewLabelHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	labelRoutes := route.RouteGroup.Group("/label")
	labelRoutes.Get("", handler.List)
	labelRoutes.Get("/:id", handler.Detail)
	labelRoutes.Use(jwtMiddleware.New)
	labelRoutes.Post("", handler.Create)
	labelRoutes.Patch("/:id", handler.Update)
	labelRoutes.Delete("/:id", handler.Delete)
	labelRoutes.Get("/export", handler.Export)
}
