package domain

import (
	"Api1/db"
	domain_models "Api1/domain/models"
)

func DeletarTerefa(id int64) (err error) {

	result := db.DB.Delete(&domain_models.Tarefa{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
	
}