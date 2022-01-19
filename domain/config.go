package domain

import (
	"fmt"
	"log"
	"os"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	gormdb "gitlab.com/s2.1-backend/shm-package-svc/gorm"
	"gitlab.com/s2.1-backend/shm-package-svc/jwe"
	"gitlab.com/s2.1-backend/shm-package-svc/jwt"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	redisPkg "gitlab.com/s2.1-backend/shm-package-svc/redis"
	"gitlab.com/s2.1-backend/shm-package-svc/str"
	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	JwtConfig     jwtware.Config
	Validator     *validator.Validate
	Redis         redisPkg.RedisClient
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	dbInfo := gormdb.Connection{
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USERNAME"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		DBMaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}

	res.DB, err = dbInfo.Conn()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "open-db-connection")
	}

	// JWE Credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	// JWT Config
	res.JwtConfig = jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Claims:     &jwt.CustomClaims{},
	}

	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	res.Redis = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}
	pong, err := res.Redis.Client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Redis ping status: "+pong, err)

	res.Validator = ValidatorDriver

	return res, err
}
