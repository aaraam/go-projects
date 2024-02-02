package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"telegram-bot/pkg"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Initialize Telegram bot
	pkg.StartBot()
	fmt.Println("Bot started")

	// Start Fiber app (this should be the last call to ensure everything above is initialized)
	log.Fatal(app.Listen(":3000"))
}
