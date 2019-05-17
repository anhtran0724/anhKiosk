package main

import (
	"./app"
	"./config"
	"./system"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	fmt.Println("Project: ", appName)

	//connect database
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)

	app.Run(system.GetEnv("APP_PORT", ":8000"))
}
