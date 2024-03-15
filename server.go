package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/xsadia/backlogs-server/pkg/config"
	"github.com/xsadia/backlogs-server/pkg/handler"
)

const defaultPort = "8080"

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	failOnError(err)
	defer db.Close()

	config.ConfigDB(db)
	failOnError(db.Ping())

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	failOnError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	failOnError(err)

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/graphql", handler.GraphQLHandler(db))
	r.GET("/", handler.PlaygroundHandler())

	log.Printf("connect to http://localhost:%s for GraphQL playground", port)
	log.Fatal(r.Run())
}
