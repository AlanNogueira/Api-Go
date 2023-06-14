package main

import (
	"Api-Go/pkg/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	fmt.Printf("Inciando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
