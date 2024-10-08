package config

import (
	"github.com/joho/godotenv"
	"log"
	"main/models"
	"os"
	"strconv"
)

// TokenConfig func to read .env files and store their data in our structs above
func TokenConfig() (*models.AccessConfig, *models.RefreshConfig) {
	//this is a built-in function from the external godotenv library which we had to download using 'go get' in the command line
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//we're accessing the config which sets the lifetime of the access token and storing it in a variable
	accessMin, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatal("Error parsing ACCESS_TOKEN_LIFETIME")
	}

	refreshHour, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatal("Error parsing REFRESH_TOKEN_LIFETIME")
	}

	//we're returning the AccessConfig and RefreshConfig structs after filling its fields with the values located in our
	//.env file
	return &models.AccessConfig{
			Port:                os.Getenv("PORT"),
			AccessTokenSecret:   os.Getenv("ACCESS_TOKEN_SECRET"),
			AccessTokenLifetime: accessMin,
		}, &models.RefreshConfig{
			Port:                 os.Getenv("PORT"),
			RefreshTokenSecret:   os.Getenv("REFRESH_TOKEN_SECRET"),
			RefreshTokenLifetime: refreshHour,
		}
}
