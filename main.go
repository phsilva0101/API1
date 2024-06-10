package main

import (
	"Api1/configs"
	"Api1/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"Api1/db"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/tarefas", handlers.ListarTodasTarefasHandler).Methods(http.MethodGet)
	r.HandleFunc("/tarefas", handlers.InserirTarefaPostHandler).Methods(http.MethodPost)
	r.HandleFunc("/tarefas/{id}", handlers.ListarTarefaPorIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/tarefas/{id}", handlers.AtualizarTarefaHandler).Methods(http.MethodPut)
	r.HandleFunc("/tarefas/{id}", handlers.DeletarTarefaHandler).Methods(http.MethodDelete)

	apiConfig := configs.GetApiConfig()

	fmt.Println("Server running on port", apiConfig.Port)

	err = http.ListenAndServe(apiConfig.Host+":"+apiConfig.Port, r)
	if err != nil {
		panic(err)
	}
}
