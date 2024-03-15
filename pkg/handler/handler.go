package handler

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/xsadia/backlogs-server/graph"
)

func GraphQLHandler(db *sql.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
