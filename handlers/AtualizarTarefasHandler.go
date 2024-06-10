package handlers

import (
	"Api1/domain"
	domain_models "Api1/domain/models"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func AtualizarTarefaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	//Converter pra int64 nativo
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Parse request
	var tarefa domain_models.Tarefa
	err = json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = domain.AtualizarTarefa(idInt64, tarefa)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tarefa)
}