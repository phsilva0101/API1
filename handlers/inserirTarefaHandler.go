package handlers

import (
	"Api1/domain/models"
	"Api1/domain"
	"encoding/json"
	"net/http"

)


func InserirTarefaPostHandler (w http.ResponseWriter, r *http.Request) {
	// Parse request
	var tarefa domain_models.Tarefa
	err := json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := domain.InserirTarefa(tarefa)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}