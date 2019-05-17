package main

import (
	"flag"
	"./app"
	"./system"
	"github.com/valyala/fasthttp"
	"log"
)

//func main() {
//	envErr := godotenv.Load()
//	if envErr != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	appName := os.Getenv("APP_NAME")
//	fmt.Println("Project: ", appName)
//
//	//connect database
//	config := config.GetConfig()
//	app := &app.App{}
//	app.Initialize(config)
//
//	app.Run(system.GetEnv("APP_PORT", ":8000"))
//}
var (
	addr     = flag.String("addr", system.GetEnv("APP_PORT", ":8000"), "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := app.RequestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

