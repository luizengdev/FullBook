package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Representa a URL para comunicação com a API
	APIURL = ""

	// Porta onde a aplicação está rodando
	Porta = 0

	// É utilizada para autenticar o cookie
	HashKey []byte

	// É utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
