package configs

import "github.com/spf13/viper"

var cfg *Configs
type Configs struct {
	ApiConfig      ApiConfig
	DataBaseConfig DataBaseConfig
}

type ApiConfig struct {
	Port string
	Host string
}

type DataBaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func Load() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg = new(Configs)
	cfg.ApiConfig = ApiConfig{
		Port: viper.GetString("api.port"),
		Host: viper.GetString("api.host"),
	}

	cfg.DataBaseConfig = DataBaseConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.dbName"),
	}

	return nil
	
}

func GetDbConfig() DataBaseConfig {
	return cfg.DataBaseConfig
}

func GetApiConfig() ApiConfig {
	return cfg.ApiConfig
}