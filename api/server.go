package api

import (
	"fmt"
	"github.com/spootrick/survi/api/router"
	"github.com/spootrick/survi/config"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	listen(config.Port)
}

func listen(port int) {
	fmt.Printf("\n\nListening... [::]:%d\n", port)

	p := fmt.Sprintf(":%d", port)
	r := router.New()
	log.Fatal(http.ListenAndServe(p, r))
}
