package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type ColorRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route ColorRouters) ColorRoute() {
	handler := v1.NewColorHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	colorRoutes := route.RouteGroup.Group("/color")
	colorRoutes.Get("", handler.List)
	colorRoutes.Get("/:id", handler.Detail)
	colorRoutes.Use(jwtMiddleware.New)
	colorRoutes.Post("", handler.Create)
	colorRoutes.Patch("/:id", handler.Update)
	colorRoutes.Delete("/:id", handler.Delete)
	colorRoutes.Get("/export", handler.Export)
}
