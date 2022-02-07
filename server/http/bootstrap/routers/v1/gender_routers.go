package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/middlewares"
)

type GenderRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route GenderRouters) GenderRoute() {
	handler := v1.NewGenderHandler(route.Handler)
	jwtMiddleware := middlewares.JwtMiddleware{Contract: route.Handler.UcContract}

	genderRoutes := route.RouteGroup.Group("/gender")
	genderRoutes.Post("", handler.List)
	genderRoutes.Get("/:id", handler.Detail)
	genderRoutes.Use(jwtMiddleware.New)
	genderRoutes.Post("", handler.Create)
	genderRoutes.Patch("/:id", handler.Update)
	genderRoutes.Delete("/:id", handler.Delete)
}
