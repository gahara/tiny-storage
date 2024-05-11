package customTypes

type EnvironmentalVariables struct {
	Env         string `mapstructure:"ENV"`
	StoragePath string `mapstructure:"STORAGE_PATH"`
}
