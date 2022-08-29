package main

import (
	"context"
	"fmt"
	"log"

	"freg/ent"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	entOptions := []ent.Option{ent.Debug()}

	client, err := ent.Open("postgres", "postgresql://localhost:5432/testent?user=postgres&password=12345&sslmode=disable", entOptions...)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	fmt.Println("connected to db")
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx, schema.WithDropColumn(true), schema.WithDropIndex(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(NewSchema(client))
	{
		e.POST("/gql", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/docs", func(c echo.Context) error {
			playground.Handler("GraphQL", "/gql").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}
	srv.Use(entgql.Transactioner{TxOpener: client})
	e.Logger.Fatal(e.Start(":8080"))
}
