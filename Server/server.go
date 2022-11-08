package server

import (
	"net/http"

	handlers "github.com/PedroSMarcal/Naval_Battle/Server/Handlers"
)

func config(mux *http.ServeMux) {
	userHandler := handlers.NewUserHandler()

	mux.HandleFunc("/user", userHandler.UserGet)
}

func Start(port string) {
	mux := http.NewServeMux()

	config(mux)

	http.ListenAndServe(port, mux)
}
