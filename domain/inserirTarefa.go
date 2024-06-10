package domain

import "Api1/db"
import "Api1/domain/models"

func InserirTarefa(tarefa domain_models.Tarefa) (int64, error) {
	result := db.DB.Create(&tarefa)
	
	if result.Error != nil {
		return 0, result.Error
	} 
	
	return tarefa.ID, nil
	
	// conn, err := db.OpenConnection()

	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()

	// sqlStatement := `INSERT INTO tarefas (titulo, descricao, data_criacao, data_conclusao, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// var id int64

	// err = conn.QueryRow(sqlStatement, tarefa.Titulo, tarefa.Descricao, tarefa.DataCriacao, tarefa.DataConclusao, tarefa.Status).Scan(&id)

	// if err != nil {
	// 	panic(err)
	// }

	// return id, err

}