package helpers

import (
	"github.com/spf13/viper"
	"log"
)

const pathToEnv = "./config/.env"

type EnvironmentalVariables struct {
	Env         string `mapstructure:"ENV"`
	StoragePath string `mapstructure:"STORAGE_PATH"`
}

func GetEnvironmentalVariables() EnvironmentalVariables {
	env := EnvironmentalVariables{}

	viper.SetConfigFile(pathToEnv)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
		return EnvironmentalVariables{}
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Println(err)
		return EnvironmentalVariables{}
	}

	return env
}
