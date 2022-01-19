package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-package-svc/functioncaller"
	jwtPkg "gitlab.com/s2.1-backend/shm-package-svc/jwt"
	"gitlab.com/s2.1-backend/shm-package-svc/logruslogger"
	"gitlab.com/s2.1-backend/shm-package-svc/messages"
	"gitlab.com/s2.1-backend/shm-package-svc/responses"
	"gitlab.com/s2.1-backend/shm-product-svc/usecase"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func (jwtMiddlewar JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}

	// Check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := []byte(jwtMiddlewar.JwtCredential.TokenSecret)
		return secret, nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Jwe roll back encrypted ID
	jweRes, err := jwtMiddlewar.JweCredential.Rollback(claims.Payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return responses.NewResponse(responses.ResponseError(nil, nil, http.StatusUnauthorized, "unauthorized", errors.New("unauthorized"))).Send(ctx)
	}

	// Set id to usecase contract
	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	roleID := fmt.Sprintf("%v", jweRes["role_id"])
	jwtMiddlewar.Contract.UserID = claims.Id
	jwtMiddlewar.Contract.RoleID = roleID

	return ctx.Next()
}
