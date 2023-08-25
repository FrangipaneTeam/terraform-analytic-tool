package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FrangipaneTeam/terraform-analytic-tool/clients"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Settings struct {
	// Address is the address of the api server.
	// Default: localhost
	Address string `yaml:"address"`

	// Port is the port of the api server.
	// Default: 8000
	Port string `yaml:"port"`

	*clients.RedisClient
	*clients.InfluxDBClient
}

type api struct {
	Settings
	*clients.RedisClient
	*clients.InfluxDBClient
}

func New(aSettings Settings) {
	a := new(api)
	a.RedisClient = aSettings.RedisClient
	a.InfluxDBClient = aSettings.InfluxDBClient

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(1 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/send", a.apiSendHandler)
	})

	if aSettings.Address == "" {
		aSettings.Address = "localhost"
	}

	if aSettings.Port == "" {
		aSettings.Port = "8000"
	}

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", aSettings.Address, aSettings.Port), r); err != nil { //nolint:gosec
		panic(err)
	}
}

func (a *api) apiSendHandler(w http.ResponseWriter, r *http.Request) {
	data := &AnalyticRequest{}
	if err := render.Bind(r, data); err != nil {
		if err := render.Render(w, r, ErrInvalidRequest(err)); err != nil {
			log.Default().Println(err)
		}
		return
	}

	if err := a.RedisClient.Publish(context.Background(), "PubSubCavTerraformAnalytics", data).Err(); err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
