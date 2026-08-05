package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const LOAD_DOTENV_FROM_FILE = false

func main() {
	if LOAD_DOTENV_FROM_FILE {

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file. ", err)
		}

		log.Println("Loaded .env locally")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", PingHandler)
	mux.HandleFunc("/user", EmailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	log.Println("server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, CORS(mux)))
}
