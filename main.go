package main

import (
	"fmt"
	"goapi/src/config"
	"goapi/src/router"
	"log"
	"net/http"
)

func main() {
	config.Load()

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Gate), r))
}
