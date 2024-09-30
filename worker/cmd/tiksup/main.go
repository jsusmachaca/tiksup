package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/jsusmachaca/tiksup/api/handler"
	"github.com/jsusmachaca/tiksup/internal/config"
	"github.com/jsusmachaca/tiksup/internal/database"
	kafkaService "github.com/jsusmachaca/tiksup/pkg/eventstream/service"
)

// var collection *mongo.Collection
var (
	configMap kafka.ConfigMap
	ctx       = context.TODO()
	db        *sql.DB
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("\033[31mNot .env file found. Using system variables\033[0m")
	}

	configMap = config.KafkaConfig()

	var err error
	db, err = database.PGConnection()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	if err = database.PGMigrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	_, err = database.MongoConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to mongo: %v", err)
	}
}

func main() {
	go kafkaService.KafkaWorker(&configMap, db)

	mux := http.NewServeMux()
	route(mux, db)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: mux,
	}
	fmt.Printf("Server listen on http://localhost%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error to initialize server %v", err)
	}
}

func route(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		handler.Login(w, r, db)
	})
	mux.HandleFunc("POST /api/register", func(w http.ResponseWriter, r *http.Request) {
		handler.Register(w, r, db)
	})
}
