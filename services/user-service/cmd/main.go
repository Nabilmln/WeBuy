package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Nabilmln/WeBuy/services/user-service/internal/handler"
	// "github.com/Nabilmln/WeBuy/services/user-service/internal/model"
	"github.com/Nabilmln/WeBuy/services/user-service/internal/repository"
	"github.com/Nabilmln/WeBuy/services/user-service/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system env")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "user-service"})
	})

	app.Post("/api/auth/register", authHandler.Register)
	app.Post("/api/auth/login", authHandler.Login)

	log.Fatal(app.Listen(":8081"))
}