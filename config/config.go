package config

import (
	"log"
	"os"
	"strconv"
)

var (
	Port = 3406
)

func Load() {
	loadPortFromEnv()
}

func loadPortFromEnv() {
	if port, exists := os.LookupEnv("API_PORT"); exists {
		p, err := strconv.Atoi(port)
		if err != nil {
			log.Printf("'API_PORT' variable could not parsed: (%s) Setting port to '%d'.", err, Port)
		} else {
			Port = p
		}
	} else {
		log.Printf("'API_PORT' variable is not set on environment. Setting port to '%d'.", Port)
	}
}
