package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	// TODO: - Update config
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PW")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//dsn := "host=34.89.238.138 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	dsn := "host=" + host + " user=" + user + " password=" + pw + " dbname=" + name + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to DB :-(")
	}

	return db
}