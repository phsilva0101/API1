package handlers

import (
	"Api1/domain"
	"encoding/json"
	"net/http"
)

func ListarTodasTarefasHandler (w http.ResponseWriter, r *http.Request) {
	tarefas, err := domain.ListarTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefas)
}