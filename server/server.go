package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/fixztter/rest-fundamentals/database"
	"github.com/fixztter/rest-fundamentals/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("Secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("Database URL is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Printf("Starting server on port%s\n", b.Config().Port)
	r, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(r)
	if err := http.ListenAndServe(b.Config().Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
