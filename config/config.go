package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Port = 3406

	DBDriver = ""
	DBUrl    = ""
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Loading '.env' file has been failed:", err)
	}

	loadPortFromEnv()
	loadDBDetailsFromEnv()
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
func loadDBDetailsFromEnv() {
	DBDriver = os.Getenv("DB_DRIVER")
	DBUrl = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	// TODO: Do not forget to enable SSL mode
}
