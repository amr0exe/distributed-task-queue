package main

import (
	"fmt"
	"net/http"

	"tryit.me/internal/handler"
	"tryit.me/internal/router"
	"tryit.me/internal/service"
	"tryit.me/internal/store"
)

func main() {
	store := store.NewMemoryStore()

	service := service.NewTaskService(store)
	handler := handler.NewTaskHandler(service)

	router := router.RoutingSetup(handler)

	fmt.Println("Serving in port :8080")
	http.ListenAndServe(":8080", router)
}
