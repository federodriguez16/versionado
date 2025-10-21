package database

import (
	"fmt"
	"os"

	"github.com/federodriguez16/versionado/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

func GetConnection() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
	}

	url := os.Getenv("VERSIONADO_DB_URL")
	user := os.Getenv("VERSIONADO_DB_USER")
	pw := os.Getenv("VERSIONADO_DB_PASSWORD")
	name := os.Getenv("VERSIONADO_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pw, url, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	if !(db.Migrator().HasTable("clients")) {
		db.AutoMigrate(&models.Client{})
	}

	if !(db.Migrator().HasTable("versions")) {
		db.AutoMigrate(&models.Versions{})
	}

	return db, nil
}
