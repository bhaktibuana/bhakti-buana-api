package main

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func SuperAdminSeeder() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(configs.DBConfig().DB_DSN, configs.DBConfig().DB_DATABASE); err != nil {
		fmt.Println(err)
		return
	}

	superAdminUser := models.Users{
		Name:        "Bhakti Mega Buana",
		Email:       "bhaktibuana19@gmail.com",
		Password:    helpers.HashPassword(os.Getenv("SUPER_ADMIN_PASSWORD")),
		AccountType: models.USER_ACCOUNT_TYPE_ADMIN,
		Status:      models.USER_STATUS_VERIFIED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := database.Users.InsertOne(context.Background(), superAdminUser); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Super admin user seeded successfully!")

	defer func() {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
