package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ConnString is the database connection string
	ConnString = ""

	// APIPort is the port where API will listen to
	APIPort = 0
)

// LoadEnv will load the environment variables
func LoadEnv() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	APIPort, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		APIPort = 9000
	}

	ConnString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
}
