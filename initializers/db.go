package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/wpcodevo/go-postgres-jwt-auth-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB initializes the database connection using the provided environment configuration.
// It constructs the DSN (Data Source Name) from the environment variables and opens a connection to the PostgreSQL database.
// If the connection is successful, it sets up the database logger, creates the "uuid-ossp" extension if it doesn't exist,
// and performs automatic migration for the User model.
// If any error occurs during the process, it logs the error and exits the application.
func ConnectDB(env *Env) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.DBHost, env.DBUserName, env.DBUserPassword, env.DBName, env.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	fmt.Println("🚀 Connected Successfully to the Database")
}
