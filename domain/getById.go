package domain

import (
	"Api1/db"
	"Api1/domain/models"
)

func ObterPorId(id int64) (domain_models.Tarefa, error) {
	var tarefa domain_models.Tarefa

	result := db.DB.First(&tarefa, id)

	if result.Error != nil {
			return domain_models.Tarefa{}, result.Error
	} 

	return tarefa, nil
	
	// conn, err := db.OpenConnection()

	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()

	// sqlStatement := `SELECT * FROM tarefas WHERE id = $1`

	// var tarefa domain_models.Tarefa

	// err = conn.QueryRow(sqlStatement, id).Scan(&tarefa.ID, &tarefa.Titulo, &tarefa.Descricao, &tarefa.DataCriacao, &tarefa.DataConclusao, &tarefa.Status)

	// if err != nil {
	// 	panic(err)
	// }

	// return tarefa, err

}