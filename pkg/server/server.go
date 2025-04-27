package server

import (
	"net/http"
	"os"

	"github.com/LiebeBleck/go_final_project/pkg/api"
)

func StartServer() error {
	port := "7540"
	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		port = envPort
	}

	api.Init()

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":"+port, nil)
}
