package main

import (
	"Api-Go/pkg/configuration"
	"Api-Go/pkg/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configuration.Load()
	configuration.DbConnect()
	r := router.Generate()

	fmt.Printf("Inciando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
