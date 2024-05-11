package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"s3/src/internal/constants"
	"s3/src/internal/customTypes"
)

//const pathToEnv = "../../internal/config/.env"
//const pathDoProdConfig = "../../internal/config/db-pg.json"

const pathToEnv = "config/.env"
const pathDoProdConfig = "config/db-pg.json"

func GetEnvironmentalVariables() customTypes.EnvironmentalVariables {
	env := customTypes.EnvironmentalVariables{}

	viper.SetConfigFile(pathToEnv)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
		return customTypes.EnvironmentalVariables{}
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Println(err)
		return customTypes.EnvironmentalVariables{}
	}
	println("_____")
	fmt.Printf("%+v\n", env)
	return env
}

func ResolveDbConf(env string) string {
	println("is this production?", env, env == "production")
	if env == "production" {
		return pathDoProdConfig
	}
	return constants.GormDB
}

func CreateDbConfig(pathToDbConfig string) string {
	if pathToDbConfig == constants.GormDB {
		return constants.GormDB
	}

	type Config struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
	}
	dbConfig := Config{}
	viper.SetConfigType("json")
	viper.SetConfigFile(pathToDbConfig)
	err := viper.ReadInConfig()
	if err != nil {
		return ""
	}

	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		return ""
	}

	stringConf := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	return stringConf
}
