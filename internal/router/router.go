package router

import (
	"fmt"
	"net/http"

	"tryit.me/internal/handler"
)

func RoutingSetup(handler *handler.TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /task", handler.CreateTask)
	mux.HandleFunc("GET /task", handler.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handler.DeleteTask)
	mux.HandleFunc("PUT /task/{id}", handler.UpdateTask)
	mux.HandleFunc("GET /all", handler.GetAll)

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello there")
	})

	return mux
}
