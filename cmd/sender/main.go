package main

import (
	"log"
	"newsletter-system/internal/storage"
	"newsletter-system/pkg/newsletter"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	emails, err := storage.GetEmails()
	if err != nil {
		log.Fatalf("Failed to retrieve emails: %v", err)
	}

	if err := newsletter.SendNewsletter(emails); err != nil {
		log.Fatalf("Failed to send newsletter: %v", err)
	}

	log.Println("Newsletter sent successfully!")
}
