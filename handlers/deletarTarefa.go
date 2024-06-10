package handlers

import (
	"Api1/domain"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func DeletarTarefaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	//Converter pra int64 nativo
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	 err = domain.DeletarTerefa(idInt64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Tarefa deletada com sucesso")
}