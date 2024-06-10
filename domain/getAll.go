package domain

import (
	"Api1/db"
	"Api1/domain/models"
)

func ListarTodos() ([]domain_models.Tarefa, error) {
	var tarefas []domain_models.Tarefa
	result := db.DB.Find(&tarefas)
	if result.Error != nil {
		return nil, result.Error
	}
	return tarefas, nil

	// conn, err := db.OpenConnection()

	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()

	// sqlStatement := `SELECT * FROM tarefas`

	// rows, err := conn.Query(sqlStatement)

	// if err != nil {
	// 	panic(err)
	// }

	// var tarefas []domain_models.Tarefa

	// for rows.Next() {
	// 	var tarefa domain_models.Tarefa

	// 	err = rows.Scan(&tarefa.ID, &tarefa.Titulo, &tarefa.Descricao, &tarefa.DataCriacao, &tarefa.DataConclusao, &tarefa.Status)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	tarefas = append(tarefas, tarefa)
	// }

	// return tarefas, err

}