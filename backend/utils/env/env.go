package structureEnv

import (
	"log"

	"github.com/spf13/viper"
)

// if yaml value not found then use value from hardcoded
func GetConfigWithDefaultSetting(path string, hardCoded interface{}) interface{} {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("read config", hardCoded)
	}
	if viper.Get(path) == nil {
		return hardCoded
	}
	return viper.Get(path)
}

func GetConfig(path string) interface{} {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("read config", err)
	}

	return viper.Get(path)
}
