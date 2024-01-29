package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionStringDB = ""

	// Gate where the API will run
	Gate = 0
)

//Load will inicialize the env variables

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Gate, err = strconv.Atoi(os.Getenv("API_GATE"))
	if err != nil {
		Gate = 9000
	}

	ConnectionStringDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=local",
		os.Getenv("DB_USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
