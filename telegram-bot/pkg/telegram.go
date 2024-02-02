package pkg

import (
	"log"
	"os"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

func StartBot() *telebot.Bot {
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

	return bot
}
