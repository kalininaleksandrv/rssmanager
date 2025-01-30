package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	returnFromRoot := func(w http.ResponseWriter, r *http.Request) {
		respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello from RssManager backend!"})
	}

	http.HandleFunc("GET /hello", returnFromRoot)

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
