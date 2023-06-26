package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SecretKey []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
