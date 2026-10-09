package main

import (
	"errors"
	"latihan_rest_api/db/drivers"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/constant"
	"latihan_rest_api/pkg/utils"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfigurance(constant.DB_USERNAME),
		Password: utils.GetConfigurance(constant.DB_PASSWORD),
		Database: utils.GetConfigurance(constant.DB_NAME),
		Host:     utils.GetConfigurance(constant.DB_HOST),
		Port:     utils.GetConfigurance(constant.DB_PORT),
	}

	repository := dbConfig.InitDB()

	provinceId, errProv := strconv.Atoi(utils.GetConfigurance(constant.ADMIN_PROVINCE_ID))
	cityId, errCity := strconv.Atoi(utils.GetConfigurance(constant.ADMIN_CITY_ID))
	districtId, errDist := strconv.Atoi(utils.GetConfigurance(constant.ADMIN_DISTRICT_ID))
	if err := errors.Join(errProv, errCity, errDist); err != nil {
		log.Fatalf("failed to read admin address configurations: %v\n", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(utils.GetConfigurance(constant.ADMIN_PASSWORD)), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to create admin password: %v\n", err)
	}

	record := models.User{
		Username:   utils.GetConfigurance(constant.ADMIN_USERNAME),
		Email:      utils.GetConfigurance(constant.ADMIN_EMAIL),
		Password:   string(password),
		Address:    utils.GetConfigurance(constant.ADMIN_ADDRESS),
		ProvinceID: uint(provinceId),
		CityID:     uint(cityId),
		DistrictID: uint(districtId),
		Role:       models.Admin,
	}

	if err := repository.Create(&record).Error; err != nil {
		log.Fatalf("failed to create admin: %v\n", err)
	}

	log.Println("admin created successfully")
}
