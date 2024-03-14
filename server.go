package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/xsadia/backlogs-server/graph"
)

const defaultPort = "8080"
const (
	defaultMaxIdleConns = 5
	defaultMaxOpenConns = 10
	defaultMaxLifetime  = 5
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	connectionString :=
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SLL_MODE"),
		)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	maxIdleConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = defaultMaxIdleConns
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = defaultMaxOpenConns
	}

	maxLifetime, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_LIFETIME"))
	if err != nil {
		maxLifetime = defaultMaxLifetime
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
