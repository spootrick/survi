package api

import (
	"fmt"
	"github.com/spootrick/survi/api/router"
	"github.com/spootrick/survi/config"
	"github.com/spootrick/survi/seed"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	seed.Load() // seed dummy values to database

	listen(config.Port)
}

func listen(port int) {
	fmt.Printf("\n\nListening... [::]:%d\n", port)

	p := fmt.Sprintf(":%d", port)
	r := router.New()
	log.Fatal(http.ListenAndServe(p, r))
}
