package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"tryit.me/internal/app"
	"tryit.me/internal/handler"
	"tryit.me/internal/repository"
	"tryit.me/internal/router"
	"tryit.me/internal/service"
)

func main() {
	_ = godotenv.Load(".env")
	connStr, ok := os.LookupEnv("DBSTRING")
	if !ok {
		log.Fatalf("DB STRING not found.")
	}

	a, err := app.New(app.Config{DBString: connStr})
	if err != nil {
		log.Fatalf("failed to initialize the app: %v", err)
		return
	}

	repo := repository.NewTaskRepository(a.DB())

	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	router := router.RoutingSetup(handler)

	fmt.Println("Serving on port :8080")
	http.ListenAndServe(":8080", router)
}
