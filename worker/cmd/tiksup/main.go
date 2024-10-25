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
	"github.com/jsusmachaca/tiksup/api/middleware"
	"github.com/jsusmachaca/tiksup/internal/config"
	"github.com/jsusmachaca/tiksup/internal/database"
	"github.com/jsusmachaca/tiksup/internal/service"
	"github.com/jsusmachaca/tiksup/pkg/movie"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	configMap  kafka.ConfigMap
	ctx        = context.TODO()
	db         *sql.DB
	mC         movie.MongoConnection
	client     *http.Client
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("\033[31mNot .env file found. Using system variables\033[0m")
	}

	client = &http.Client{}

	configMap = config.KafkaConfig()

	var err error
	db, err = database.PGConnection()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	if err = database.PGMigrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	collection, err = database.MongoConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to mongo: %v", err)
	}
	mC = movie.MongoConnection{Collection: collection, CTX: ctx}
}

func main() {
	go service.KafkaWorker(client, &configMap, db, mC)

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
	userInfo := &handler.GetUserInfo{DB: db}
	randomMovies := &handler.GetRandomMovies{DB: db, MongoConn: mC}

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		handler.Login(w, r, db)
	})
	mux.HandleFunc("POST /api/register", func(w http.ResponseWriter, r *http.Request) {
		handler.Register(w, r, db)
	})
	mux.Handle("GET /user-info", middleware.AuthMiddleware(userInfo))
	mux.Handle("GET /random-movies", middleware.AuthMiddleware(randomMovies))
}
