package main

import (
	"log"

	"icecream-backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	log.Println("Database connected and tables ready")

	app := fiber.New()

	app.Get("/api/health", func(c *fiber.Ctx) error {
		if err := sqlDB.Ping(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(
				fiber.Map{"status": "database unavailable"},
			)
		}

		return c.JSON(fiber.Map{
			"status":   "ok",
			"database": "connected",
		})
	})

	if err := app.Listen("127.0.0.1:3000"); err != nil {
		log.Print(err)
	}
}
