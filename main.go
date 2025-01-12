package main

import (
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello from RssManager backend!"})
	})

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
