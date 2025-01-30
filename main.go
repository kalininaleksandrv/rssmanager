package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kalininaleksandrv/rssmanager/internal/database"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("$DB_URL must be set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbCfg := dbConfig{
		DB : database.New(conn),
	}

	http.HandleFunc("POST /user", dbCfg.handlerCreateUser)

	http.HandleFunc("GET /user/id", dbCfg.handlerGetUser)

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
