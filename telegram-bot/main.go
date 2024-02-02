package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gopkg.in/tucnak/telebot.v2"
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
	token_key := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token_key,                                      // Token from .env
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second}, // Poller with timeout
	})
	if err != nil {
		log.Fatal(err)
	}

	// Handle /start command
	bot.Handle("/start", func(m *telebot.Message) {
		_, err := bot.Send(m.Sender, "Hello! I'm a bot, I can help you with your tasks")
		if err != nil {
			log.Println("Error sending message:", err)
		}
	})

	// Start the bot in a goroutine
	go func() {
		bot.Start()
	}()

	// Start Fiber app (this should be the last call to ensure everything above is initialized)
	log.Fatal(app.Listen(":3000"))
}
