package api

import (
	"fmt"
	"latihan_rest_api/api/handlers"
	"latihan_rest_api/api/middlewares"
	"latihan_rest_api/auth"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/fileupload"
	"latihan_rest_api/users"

	"github.com/cloudinary/cloudinary-go/v2"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig) *echo.Echo {
	var (
		e           = echo.New()
		authService = auth.New(repository)
		userService = users.New(repository)
		uploader    = &fileupload.CloudinaryUploader{Cld: cld}
		authHandler = handlers.NewAuth(authService, jwtConfig)
		userHandler = handlers.NewUsers(userService, uploader)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	logger, _ := zap.NewProduction()
	loggerConfig := middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}

	logMiddleware := middlewares.LoggerConfig{Config: loggerConfig}
	jwtMiddleware := jwtConfig.Init()

	e.Use(logMiddleware.Init())

	autthRoutes := e.Group(fmt.Sprintf("%s/auth", constant.API_V1_PREFIX))

	autthRoutes.POST("/register", authHandler.RegisterUser)
	autthRoutes.POST("/login", authHandler.LoginUser)

	userRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	userRoutes.GET("/profile", userHandler.GetProfile)
	userRoutes.PATCH("/profile/edit", userHandler.UpdateProfile)

	return e
}
