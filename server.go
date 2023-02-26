package main

import (
	"github.com/adaralex/trinity/graph"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adaralex/trinity/graph/db"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("starting database", "time", time.Now().UnixMilli())

	databaseConnector := &db.SecurityDatabase{}
	err := databaseConnector.Open()
	if err != nil {
		sugar.Errorw("error loading database", "time", time.Now().UnixMilli(), "error", err.Error())
	}

	_ = databaseConnector.CreateProject("Flexwitch")
	_ = databaseConnector.CreateScanner("Nuclei")
	project, err := databaseConnector.FindProject("Flexwitch")
	if err != nil {
		sugar.Errorw("error getting project", "time", time.Now().UnixMilli(), "error", err.Error())
	}
	_ = databaseConnector.AddScannerAnalysis(project, "Nuclei", "* * * * * *", 12)
	_ = databaseConnector.UpdateScanner("Nuclei", db.DatabaseScanner{
		Install: "go get nuclei",
		Run:     "go run nuclei",
		Report:  "cat report.json",
		Type:    "URL",
	})
	_ = databaseConnector.CreateUser("pano")
	_ = databaseConnector.AddUserRole(project, "pano", "admin")
	project, err = databaseConnector.FindProject("Flexwitch")
	if err != nil {
		sugar.Errorw("error getting project", "time", time.Now().UnixMilli(), "error", err.Error())
	}

	sugar.Infow("starting server", "time", time.Now().UnixMilli())

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{
				DB:     databaseConnector,
				Logger: sugar,
			}}))

	// TODO Implement rate-limiting to protect server
	// https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
	//
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
