package config

import (
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/repository"
	"log"

	"gorm.io/gorm"
)

// SeedAdminUser creates an initial admin user if it doesn't exist
func SeedAdminUser(db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	// Check if admin user already exists
	if userRepo.CheckUsernameExists("admin") {
		log.Println("Admin user already exists, skipping...")
		return
	}

	// Create admin user
	hashedPassword, err := HashPassword("admin123")
	if err != nil {
		log.Printf("Error hashing admin password: %v", err)
		return
	}

	adminUser := &model.User{
		Username: "admin",
		Email:    "admin@example.com",
		Password: hashedPassword,
		FullName: "System Administrator",
		Role:     "admin",
		IsActive: true,
	}

	if err := userRepo.CreateUser(adminUser); err != nil {
		log.Printf("Error creating admin user: %v", err)
		return
	}

	log.Println("Admin user created successfully")
	log.Println("Username: admin")
	log.Println("Password: admin123")
	log.Println("Please change the password after first login!")
}
