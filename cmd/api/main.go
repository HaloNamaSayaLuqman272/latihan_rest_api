package main

import (
	"fmt"
	"latihan_rest_api/api"
	"latihan_rest_api/api/middlewares"
	"latihan_rest_api/db/drivers"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/fileupload"
	"latihan_rest_api/pkg/utils"
	"log"
	"strconv"
)

func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfigurance(constant.DB_USERNAME),
		Password: utils.GetConfigurance(constant.DB_PASSWORD),
		Database: utils.GetConfigurance(constant.DB_NAME),
		Host:     utils.GetConfigurance(constant.DB_HOST),
		Port:     utils.GetConfigurance(constant.DB_PORT),
	}

	cloudinaryConfig := fileupload.CludinaryConfig{
		CloudinaryURL: utils.GetConfigurance(constant.CLOUDINARY_URL),
	}

	expireDuration, err := strconv.Atoi(utils.GetConfigurance(constant.JWT_EXPIRE_DURATION))
	if err != nil {
		log.Fatalf("error when parsing expire duration: %v\n", err)
	}

	jwtConfig := middlewares.JWTConfig{
		SecretKey:      utils.GetConfigurance(constant.JWT_SECRET_KEY),
		ExpireDuration: expireDuration,
	}

	var (
		repository = dbConfig.InitDB()
		cloudinary = cloudinaryConfig.InitCloudinary()
		e          = api.NewEcho(repository, cloudinary, jwtConfig)
	)

	drivers.MigrateDB(repository)

	appPort := fmt.Sprintf(":%s", utils.GetConfigurance(constant.PORT))
	if err := e.Start(appPort); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
