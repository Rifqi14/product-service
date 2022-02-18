package handlers

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-package-svc/jwe"
	"gitlab.com/s2.1-backend/shm-package-svc/jwt"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
	"gorm.io/gorm"
)

type Handler struct {
	App           *fiber.App
	UcContract    *usecase.Contract
	DB            *gorm.DB
	Validate      *validator.Validate
	Translator    ut.Translator
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
}

type (
	FileType     string
	PlatformType string
)

var (
	Pdf FileType = "pdf"
	Csv FileType = "csv"
)

var (
	Website   PlatformType = "website"
	Email     PlatformType = "email"
	Instagram PlatformType = "instagram"
	Facebook  PlatformType = "facebook"
	Tiktok    PlatformType = "tiktok"
	Twitter   PlatformType = "twitter"
	Other     PlatformType = "other"
)

func (ft FileType) IsValid() bool {
	switch ft {
	case Pdf, Csv:
		return true
	}
	return false
}

func (pt PlatformType) IsValid() bool {
	switch pt {
	case Website, Email, Instagram, Facebook, Tiktok, Twitter, Other:
		return true
	}
	return false
}
