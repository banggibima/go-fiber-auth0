package main

import (
	"log"

	"github.com/banggibima/go-fiber-auth0/config"
	"github.com/banggibima/go-fiber-auth0/internal/interface/http"
	fiberpkg "github.com/banggibima/go-fiber-auth0/pkg/fiber"
)

func main() {
	cfg, err := config.ConfigInit()
	if err != nil {
		log.Fatal(err)
	}

	fiber, err := fiberpkg.FiberInit()
	if err != nil {
		log.Fatal(err)
	}

	http.AppServer(&http.Server{
		Fiber:  fiber,
		Config: cfg,
	})
}
