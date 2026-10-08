package middlewares

import (
	"context"
	"errors"
	"latihan_rest_api/db/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type JWTCustomsClaims struct {
	ID   int         `json:"id"`
	Role models.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey      string
	ExpireDuration int
}

type ContextKey string

const userContextKey = ContextKey("user")

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(JWTCustomsClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

func (jwtConfig *JWTConfig) GenerateToken(userID int, role models.Role) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(int64(jwtConfig.ExpireDuration))))

	claims := &JWTCustomsClaims{
		ID: userID,
		// klaim ini milik siapa
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
			// kadaluwarsa token dimulai
		},
		Role: role,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token mulai dibentuk
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(ctx context.Context) (*JWTCustomsClaims, error) {
	user, ok := ctx.Value(userContextKey).(*jwt.Token)
	if !ok || user == nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := user.Claims.(*JWTCustomsClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetUserID(ctx context.Context) (int, error) {
	claims, err := GetUser(ctx)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	return claims.ID, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		ctx := context.WithValue(c.Request().Context(), userContextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))

		userData, err := GetUser(ctx)

		if userData == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		c.Set("userData", userData)
		// "c.Set(...)" perintah menyimpan data "userData"
		return next(c)
	}
}

func VerifyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := GetUser(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if user.Role != models.Admin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access denied",
			})
		}

		return next(c)
	}

}
