package handlers

import "github.com/gofiber/fiber/v2"

type IProductHandler interface {
	Create(ctx *fiber.Ctx) (err error)

	List(ctx *fiber.Ctx) (err error)

	Detail(ctx *fiber.Ctx) (err error)

	Update(ctx *fiber.Ctx) (err error)

	Delete(ctx *fiber.Ctx) (err error)

	Export(ctx *fiber.Ctx) (err error)

	ChangeStatus(ctx *fiber.Ctx) (err error)
}
