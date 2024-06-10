package domain

import (
	"Api1/db"
	"Api1/domain/models"
)

func AtualizarTarefa(id int64, tarefa domain_models.Tarefa ) ( error) {

	result := db.DB.Save(&tarefa)
	if result.Error != nil {
		return result.Error
	}
	return  nil

	// conn, err := db.OpenConnection()

	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()

	// //Update the task
	// sqlStatement := `UPDATE tarefas SET titulo = $1, descricao = $2, data_criacao = $3, data_conclusao = $4, status = $5 WHERE id = $6`

	// _, err = conn.Exec(sqlStatement, tarefa.Titulo, tarefa.Descricao, tarefa.DataCriacao, tarefa.DataConclusao, tarefa.Status, id)

	// if err != nil {

	// 	panic(err)
	// }


	// return nil
}