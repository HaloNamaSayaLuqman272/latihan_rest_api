package api

import (
	"fmt"
	"latihan_rest_api/api/handlers"
	"latihan_rest_api/api/middlewares"
	"latihan_rest_api/auth"
	"latihan_rest_api/categories"
	"latihan_rest_api/pkg/ai"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/fileupload"
	"latihan_rest_api/pkg/rabbitmq"
	"latihan_rest_api/products"
	"latihan_rest_api/users"

	"github.com/cloudinary/cloudinary-go/v2"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig, mq rabbitmq.Publisher) *echo.Echo {
	var (
		e                 = echo.New()
		authService       = auth.New(repository, mq)
		userService       = users.New(repository)
		categoryService   = categories.New(repository)
		productService    = products.New(repository)
		aiService         = ai.InitService()
		uploader          = &fileupload.CloudinaryUploader{Cld: cld}
		authHandler       = handlers.NewAuth(authService, jwtConfig)
		usersHandler      = handlers.NewUsers(userService, uploader)
		categoriesHandler = handlers.NewCategories(categoryService)
		productsHandler   = handlers.NewProducts(productService, uploader, aiService)
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

	userRoutes.GET("/profile", usersHandler.GetProfile)
	userRoutes.PATCH("/profile/edit", usersHandler.UpdateProfile)

	categoryRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	categoryRoutes.POST("/categories", categoriesHandler.CreateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.GET("/categories/:id", categoriesHandler.GetCategoryByID)
	categoryRoutes.GET("/categories", categoriesHandler.GetAllCategories)
	categoryRoutes.PUT("/categories/:id", categoriesHandler.UpdateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.DELETE("/categories/:id", categoriesHandler.DeleteCategory, middlewares.VerifyAdmin)

	productRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	productRoutes.POST("/products", productsHandler.CreateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.ProductRequest{}))
	productRoutes.GET("/products/:id", productsHandler.GetProductByID)
	productRoutes.GET("/products/category/:id", productsHandler.GetProductsByCategory)
	productRoutes.GET("/products", productsHandler.GetAllProducts)
	productRoutes.POST("/products/recommendation", productsHandler.GetProductsByRecommendationAI, middlewares.ValidateBody(&ai.ProductRecommendationRequest{}))
	productRoutes.PUT("/products/:id", productsHandler.UpdateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.ProductRequest{}))
	productRoutes.DELETE("/products/:id", productsHandler.DeleteProduct, middlewares.VerifyAdmin)

	return e
}
