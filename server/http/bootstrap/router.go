package bootstrap

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	v1 "gitlab.com/s2.1-backend/shm-product-svc/server/http/bootstrap/routers/v1"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/handlers"
)

func (boot Bootstrap) ProductRoute() {
	handlerType := handlers.Handler{
		App:        boot.App,
		UcContract: &boot.UcContract,
		DB:         boot.DB,
		Validate:   boot.Validator,
		Translator: boot.Translator,
	}

	// Route for check health
	rootParentGroup := boot.App.Group("/product")
	rootParentGroup.Get("", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Product service is working")
	})

	v1Routers := v1.Routers{
		RouteGroup: rootParentGroup,
		Handler:    handlerType,
	}
	v1Routers.ProductRoute()
}
