package drivers

import (
	"fmt"
	"latihan_rest_api/db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

func (config *DBConfig) InitDB() *gorm.DB {
	var err error
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.Username, config.Password, config.Database, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("error when connecting to the database: %s\n", err)
	}

	fmt.Println("connected to the database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}

	log.Println("database migrate succed")
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connction: %v", err)
		return err
	}

	log.Println("connection database is closed")
	return nil
}
