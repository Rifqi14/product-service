package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
)

type Routers struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (routers Routers) ProductRoute() {
	apiV1 := routers.RouteGroup.Group("/v1")

	brandRoutes := BrandRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	brandRoutes.BrandRoute()

	categoryRoutes := CategoryRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	categoryRoutes.CategoryRoute()

	colorRoutes := ColorRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	colorRoutes.ColorRoute()

	genderRoutes := GenderRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	genderRoutes.GenderRoute()

	labelRoutes := LabelRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	labelRoutes.LabelRoute()

	materialRoutes := MaterialRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	materialRoutes.MaterialRoute()

	productRoutes := ProductRouters{
		RouteGroup: apiV1,
		Handler:    routers.Handler,
	}
	productRoutes.ProductRoute()
}
