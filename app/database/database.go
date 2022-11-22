package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Conn *gorm.DB

var err error

const (
	db_env_host     = "DB_HOST"
	db_env_port     = "DB_PORT"
	db_env_username = "DB_USERNAME"
	db_env_password = "DB_PASSWORD"
	db_env_database = "DB_DATABASE"
)

func Init() {
	godotenv.Load(".env")
	database_host := os.Getenv(db_env_host)
	if database_host == "" {
		log.Panicf("%s environment value has not set.\n", db_env_host)
	}
	database_port := os.Getenv(db_env_port)
	if database_port == "" {
		log.Panicf("%s environment value has not set.\n", db_env_port)
	}
	database_username := os.Getenv(db_env_username)
	if database_username == "" {
		log.Panicf("%s environment value has not set.\n", db_env_username)
	}
	database_password := os.Getenv(db_env_password)
	database_name := os.Getenv(db_env_database)
	if database_name == "" {
		log.Panicf("%s environment value has not set.\n", db_env_database)
	}
	connect_string := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		database_username,
		database_password,
		database_host,
		database_port,
		database_name,
	)
	Conn, err = gorm.Open("mysql", connect_string)

	if err != nil {
		fmt.Println("Db not conned")
	} else {
		fmt.Println("DB is Connected")
	}
}
