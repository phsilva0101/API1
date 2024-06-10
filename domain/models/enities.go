package domain_models

import "time"

type Tarefa struct {
	ID            int64     `gorm:"primaryKey"`
	Titulo        string
	Descricao     string
	DataCriacao   time.Time
	DataConclusao time.Time
	Status        string
}

func (Tarefa) TableName() string {
	return "tarefas"
}
