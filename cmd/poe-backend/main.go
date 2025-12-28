package main

import (
	"context"
	"log"

	"github.com/dev-universe/poe-backend/internal/app"
	httpserver "github.com/dev-universe/poe-backend/internal/http"
)

func main() {
	ctx := context.Background()

	cfg := app.LoadConfig()
	built, err := app.Build(ctx, cfg)
	if err != nil {
		log.Fatalf("build app: %v", err)
	}

	h := httpserver.NewProofsHandler(built.Service, built.Cfg.MaxUpload)
	r := httpserver.NewRouter(h)

	log.Printf("listening on %s", built.Cfg.ListenAddr)
	if err := r.Run(built.Cfg.ListenAddr); err != nil {
		log.Fatalf("server run: %v", err)
	}
}
