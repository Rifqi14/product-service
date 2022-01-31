package handlers

import "github.com/gofiber/fiber/v2"

type IMaterialCategoryHandler interface {
	Create(ctx *fiber.Ctx) (err error)

	List(ctx *fiber.Ctx) (err error)

	Detail(ctx *fiber.Ctx) (err error)

	Update(ctx *fiber.Ctx) (err error)

	Delete(ctx *fiber.Ctx) (err error)
}
