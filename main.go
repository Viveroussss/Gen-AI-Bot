package main

import (
	"log"
	"Gen-AI-Bot/bot"
)

func main() {
	bot, err := bot.NewBot()
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	log.Println("Bot started...")
	if err := bot.Start(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
} 