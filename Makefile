gen-graphql:
	@go run github.com/99designs/gqlgen generate

run:
	@go run ./server.go