package main

import (
	"fmt"
	"log"

	"github.com/Redarcher9/public-library/config"
	"github.com/Redarcher9/public-library/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var envConfig *config.Config

func main() {
	// Load config from environment
	envConfig = config.Init()
	fmt.Printf("Loaded config: %+v\n", envConfig)

	// Initialize database
	dbInstance := setUpDatabase()
	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from GORM DB: %v", err)
	}

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Create Gin router
	r := gin.Default()

	// Healthcheck route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Setup Routes and Swagger URLs
	routes.SetupRoutes(r, dbInstance)

	// Start Gin server
	port := fmt.Sprintf(":%s", envConfig.APIPort)
	log.Printf("Starting server on %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// setUpDatabase initializes and returns a GORM DB instance
func setUpDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		envConfig.DbHost,
		envConfig.DbUsername,
		envConfig.DbPassword,
		envConfig.DbName,
		envConfig.DbPort,
		envConfig.DbSSLMode,
	)

	fmt.Println("Connecting to Postgres with DSN:", dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
