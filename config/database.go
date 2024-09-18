package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type dbConfig struct {
	DB_HOST 	  string
	DB_PORT   	  string
    DB_USERNAME   string
    DB_PASSWORD   string
    DB_DATABASE   string
}

func (config *dbConfig) initDBEnv() {
	config.DB_HOST = os.Getenv("DB_HOST")
    config.DB_PORT = os.Getenv("DB_PORT")
    config.DB_USERNAME = os.Getenv("DB_USERNAME")
    config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
    config.DB_DATABASE = os.Getenv("DB_DATABASE")
}

func (config *dbConfig) initConnectionString() string {
	/**
	*    Change the connection string according to your database configuration.
    *    For example, if your database is hosted on localhost, port 3306, username is "user", password is "pass", and database name is "dbname", your connection string should look like this:
    *        dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    *        For more details, refer to https://gorm.io/docs/connecting_to_the_database.html#connecting-with-a-database
    *        For more information about DSN, refer to https://golang.org/pkg/database/sql/driver/#Driver.Open
    *        For more information about gorm.Open, refer to https://gorm.io/docs/connecting_
	**/

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.DB_USERNAME,
        config.DB_PASSWORD,
        config.DB_HOST,
        config.DB_PORT,
        config.DB_DATABASE,
    );
}

func CreateConnection() *gorm.DB {
	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

	dbCfg := dbConfig{};
	dbCfg.initDBEnv();

	db, err := gorm.Open(mysql.Open(dbCfg.initConnectionString()), &gorm.Config{})

	if err!= nil {
		fmt.Println(os.Getenv("DB_HOST"))
        panic("Failed to connect to database!")
    }

	return db;
}
