package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RiteshParouha/go-ecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // Used for rate limiting
	r.Use(middleware.RealIP)    // Tracks the request IP
	r.Use(middleware.Logger)    // It gives additional details in the logs
	r.Use(middleware.Recoverer) // Recover from crash

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	h := products.NewHandler(nil)

	r.Get("/products", h.ListProducts)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addrs,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server started on port", app.config.addrs)

	return srv.ListenAndServe()
}

type application struct {
	config config
	//logger
	//db driver
}

type config struct {
	addrs string //8080
	db    dbConfig
}

type dbConfig struct {
	dsn string
}
