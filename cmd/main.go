// @title           Go Gin GORM Backend API
// @version         1.0
// @description     A RESTful API for managing topics and topic details using Go, Gin, and GORM
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"fmt"
	"go-gin-gorm-backend/config"
	"log"

	"go-gin-gorm-backend/handler"
	"go-gin-gorm-backend/repository"
	"go-gin-gorm-backend/router"
	"go-gin-gorm-backend/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	fmt.Println("Server starting...")
	dbConfig := config.LoadDBConfig()
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Auto-migrate models
	if err := config.MigrateDB(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Seed initial topics
	config.SeedTopics(db)

	// Seed initial topic details
	config.SeedTopicDetails(db)

	// Seed admin user
	config.SeedAdminUser(db)

	// Initialize JWT
	config.InitJWT()

	// Initialize repositories
	topicRepo := repository.NewTopicRepository(db)
	topicDetailRepo := repository.NewTopicDetailRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	topicService := service.NewTopicService(topicRepo)
	topicDetailService := service.NewTopicDetailService(topicDetailRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	topicHandler := handler.NewTopicHandler(topicService)
	topicDetailHandler := handler.NewTopicDetailHandler(topicDetailService)
	authHandler := handler.NewAuthHandler(userService)

	// Setup router
	r := router.SetupRouter(topicHandler, topicDetailHandler, authHandler)

	// Start server
	r.Run()
}
