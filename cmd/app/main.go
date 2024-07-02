package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/srivalli123125/Todo-Space/internal/app/api"
	"github.com/srivalli123125/Todo-Space/internal/app/repository"
	"github.com/srivalli123125/Todo-Space/internal/app/routes"
	"github.com/srivalli123125/Todo-Space/internal/app/scylladb"
)

func main() {
	session, err := scylladb.SetupScyllaDB()
	if err != nil {
		log.Fatalf("Error setting up ScyllaDB: %v", err)
	}
	defer session.Close()

	repo := repository.NewTodoRepository(session)

	router := gin.Default()

	api := api.NewTodoAPI(repo)

	routes.SetupRoutes(router, api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
