package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"Api1/configs"
	"Api1/domain/models"
)

var DB *gorm.DB

func Init() {
	var err error
	conf := configs.GetDbConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DbName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Aplicar migrações
	err = DB.AutoMigrate(&domain_models.Tarefa{})
	if err != nil {
		log.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	fmt.Println("Conectado ao banco de dados e migrações aplicadas")
}
