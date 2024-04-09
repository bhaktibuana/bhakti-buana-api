package main

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func AboutMeSeeder() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(configs.DBConfig().DB_DSN, configs.DBConfig().DB_DATABASE); err != nil {
		fmt.Println(err)
		return
	}

	payload := models.About{
		NickName:  "",
		Role:      "",
		Summary:   "",
		Email:     "",
		Photo:     "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := database.About.InsertOne(context.Background(), payload); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Base about seeded successfully!")

	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
