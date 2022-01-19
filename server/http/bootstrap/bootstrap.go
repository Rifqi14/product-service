package bootstrap

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
	"gorm.io/gorm"
)

type Bootstrap struct {
	App        *fiber.App
	DB         *gorm.DB
	UcContract usecase.Contract
	Validator  *validator.Validate
	Translator ut.Translator
}
