package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-product-svc/domain"
	"gitlab.com/s2.1-backend/shm-product-svc/migrations"
	"gitlab.com/s2.1-backend/shm-product-svc/server/http/bootstrap"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
)

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
	validatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	config, err := domain.LoadConfig()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-config")
	}

	dbConn := config.DB
	dbClose, _ := config.DB.DB()
	migrations.Migrate(config.DB)
	defer dbClose.Close()

	app := fiber.New()

	ValidatorInit()

	ucContract := usecase.Contract{
		App:           app,
		DB:            dbConn,
		JweCredential: config.JweCredential,
		JwtCredential: config.JwtCredential,
		Validate:      validatorDriver,
		Translator:    translator,
	}

	// Bootstrap init
	boot := bootstrap.Bootstrap{
		App:        app,
		DB:         dbConn,
		UcContract: ucContract,
		Validator:  validatorDriver,
		Translator: translator,
	}

	boot.App.Use(recover.New())
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New())
	boot.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot.ProductRoute()
	fmt.Sprintln(boot.App)

	log.Fatal(boot.App.Listen(os.Getenv("APP_HOST")))
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	validatorDriver = validator.New()

	err := enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(validatorDriver, transID)
	if err != nil {
		fmt.Println(err)
	}
	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
