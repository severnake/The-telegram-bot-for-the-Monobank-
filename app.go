package main

import (
	"os"
)

func main() {
	bot := New(
		os.Getenv("TELEGRAM_TOKEN"),
		os.Getenv("TELEGRAM_ADMINS"),
		os.Getenv("TELEGRAM_CHATS"),
		os.Getenv("MONO_TOKENS"),
	)

	go bot.TelegramStart()
	go bot.ProcessingStart()

	// run http server
	bot.WebhookStart()
}
