package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"task/app"
	"task/app/container"
	middleware2 "task/middleware"
	"task/router"

	"github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	cors "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// TODO - add global logger
// TODO - add error handling
// TODO - add test cases and start looking into TDD
// TODO - change receiver implementations to pointer for services, to keep the container light

var isDev *bool
var port string

func init() {
	isDev = flag.Bool("dev", false, "To run server in dev mode")
	flag.Parse()

	port = ""
	if *isDev {
		envPath := "./.env"
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("%s was not found", envPath)
		}
	}

	if envPort, exists := os.LookupEnv(app.EnvServerPortKey); exists {
		port = envPort
	} else {
		port = "80"
	}

}

func main() {
	apiInstance := app.InitializeAPI(context.Background())
	container.InitControllers(apiInstance)
	// different sub routers for different middlewares
	// corresponding lambda handlers to be found grouped by APIGW authorizer under handler directory
	apiSubRouter := chi.NewRouter() // To be used for RESTful APIs
	webhookSubRouter := chi.NewRouter()

	router.API.Init(apiSubRouter, apiInstance)

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware2.ApplyPanicRecovery)
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		// AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))
	mainRouter.Mount("/api", apiSubRouter)
	mainRouter.Mount("/webhook", webhookSubRouter)
	mainRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	if app.IsDebug() {
		if logErr := muxDebugLogger(mainRouter); logErr != nil {
			log.Print(logErr)
		}
	}

	startServer(mainRouter)
}

func startServer(router http.Handler) {
	fmt.Printf("\n server started on port %s. Do ctrl+c to exit... \n", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}

func muxDebugLogger(router *chi.Mux) error {
	return chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: %s has %d middlewares\n", method, route, len(middlewares))
		return nil
	})
}
