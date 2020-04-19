package api

import (
	"fmt"
	"github.com/spootrick/survi/api/router"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("Listening... [::]:3000")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}
