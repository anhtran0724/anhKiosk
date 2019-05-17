package app

import (
	"fmt"
	"log"

	"../config"
	"../model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
	ctx    *fasthttp.RequestCtx
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(fasthttp.ListenAndServe(host, fastHTTPHandler))
}
