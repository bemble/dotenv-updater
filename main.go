package main

import (
	"dotenv-updater/api"
	"dotenv-updater/core/config"
	"net/http"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	logLevel := log.InfoLevel
	if config.IsDebug() {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	if config.IsDebug() {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	basePath := config.GetBasePath()

	r.Route(basePath, api.Router)

	log.WithField("category", "general").Info("Server started: http://0.0.0.0:5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}
