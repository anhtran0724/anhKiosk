package main

import (
	"flag"
	"github.com/valyala/fasthttp"
	"os"
	"fmt"
	"./app"
	"./system"
	"./config"
	"github.com/joho/godotenv"
	"log"
)
var (
	addr     = flag.String("addr", system.GetEnv("APP_PORT", ":8000"), "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	//Load env
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Project: ", os.Getenv("APP_NAME"))

	//connect database
	app := &app.App{}
	app.Initialize(config.GetConfig())

	//route division
	h := app.RequestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}
